package mgo

import (
	"fmt"
	"log"
	"time"

	"github.com/sfi2k7/mgo/bson"
)

// CopyOptions defines all options for the database copy operation
type CopyOptions struct {
	// Source database configuration
	SourceURI      string
	SourceDatabase string

	// Target database configuration
	TargetURI      string
	TargetDatabase string

	// Collections to copy (empty means all collections)
	CollectionsToCopy []string

	// Whether to delete the target database before copying
	DeleteTargetDatabase bool

	// Whether to delete collections in the target before copying
	DeleteTargetCollections bool

	// Number of documents to process in each batch
	BatchSize int

	// Optional logger (nil means use standard log)
	Logger *log.Logger
}

// CopyStats provides statistics about the copy operation
type CopyStats struct {
	CollectionsCopied int
	DocumentsCopied   int
	BytesCopied       int64
	StartTime         time.Time
	EndTime           time.Time
}

// CopyDatabase copies data from one MongoDB database to another
func CopyDatabase(options CopyOptions) (CopyStats, error) {
	stats := CopyStats{
		StartTime: time.Now(),
	}

	// Use default batch size if not specified
	if options.BatchSize <= 0 {
		options.BatchSize = 1000
	}

	// Setup logging
	logger := options.Logger
	if logger == nil {
		logger = log.New(log.Writer(), "[MongoDB Copy] ", log.LstdFlags)
	}

	// Connect to source database
	sourceSession, err := Dial(options.SourceURI)
	if err != nil {
		return stats, fmt.Errorf("failed to connect to source MongoDB: %w", err)
	}
	defer sourceSession.Close()
	// sourceSession.SetMode(mgo.Strong, true)
	sourceDB := sourceSession.DB(options.SourceDatabase)

	// Connect to target database
	targetSession, err := Dial(options.TargetURI)
	if err != nil {
		return stats, fmt.Errorf("failed to connect to target MongoDB: %w", err)
	}
	defer targetSession.Close()

	targetDB := targetSession.DB(options.TargetDatabase)

	// Handle database deletion if requested
	if options.DeleteTargetDatabase {
		logger.Printf("Dropping target database: %s", options.TargetDatabase)
		if err := targetDB.DropDatabase(); err != nil {
			return stats, fmt.Errorf("failed to drop target database: %w", err)
		}
	}

	// Get collections to copy
	collections := options.CollectionsToCopy
	if len(collections) == 0 {
		// If no collections specified, get all collections from source
		collections, err = sourceDB.CollectionNames()
		if err != nil {
			return stats, fmt.Errorf("failed to get collection names: %w", err)
		}
	}

	// Copy each collection
	for _, collName := range collections {
		// Skip system collections
		if len(collName) > 0 && collName[0] == '_' {
			logger.Printf("Skipping system collection: %s", collName)
			continue
		}

		logger.Printf("Processing collection: %s", collName)

		// Handle collection deletion if requested
		if options.DeleteTargetCollections {
			if err := targetDB.C(collName).DropCollection(); err != nil {
				logger.Printf("Warning: Failed to drop target collection %s: %v", collName, err)
				// Continue anyway
			}
		}

		// Copy indexes before data
		if err := copyIndexes(sourceDB, targetDB, collName, logger); err != nil {
			return stats, err
		}

		// Copy the collection data
		collStats, err := copyCollection(sourceDB, targetDB, collName, options.BatchSize, logger)
		if err != nil {
			return stats, err
		}

		// Update stats
		stats.CollectionsCopied++
		stats.DocumentsCopied += collStats.DocumentsCopied
		stats.BytesCopied += collStats.BytesCopied
	}

	stats.EndTime = time.Now()
	duration := stats.EndTime.Sub(stats.StartTime).Seconds()
	logger.Printf("Copy completed in %.2f seconds. Copied %d collections, %d documents, %.2f MB",
		duration, stats.CollectionsCopied, stats.DocumentsCopied, float64(stats.BytesCopied)/(1024*1024))

	return stats, nil
}

// collectionStats holds statistics for a collection copy operation
type collectionStats struct {
	DocumentsCopied int
	BytesCopied     int64
}

// copyIndexes copies all indexes from source collection to target collection
func copyIndexes(sourceDB, targetDB *Database, collName string, logger *log.Logger) error {
	// Get indexes from source collection
	indexes := sourceDB.C(collName).ListIndexes()

	// Create each index on target collection
	for _, idx := range indexes {
		// Skip the default _id index as it's created automatically
		if idx.Key[0].Key == "_id" && len(idx.Key) == 1 {
			continue
		}

		logger.Printf("Creating index %v on collection %s", idx.Key, collName)
		if err := targetDB.C(collName).CreateIndex(idx); err != nil {
			return fmt.Errorf("failed to create index %v on collection %s: %w", idx.Key, collName, err)
		}
	}

	return nil
}

// copyCollection copies all documents from source collection to target collection
func copyCollection(sourceDB, targetDB *Database, collName string, batchSize int, logger *log.Logger) (collectionStats, error) {
	stats := collectionStats{}
	sourceCollection := sourceDB.C(collName)
	targetCollection := targetDB.C(collName)

	// Get total count for progress reporting
	total, err := sourceCollection.EstimatedCount()
	if err != nil {
		return stats, fmt.Errorf("failed to get document count for collection %s: %w", collName, err)
	}
	logger.Printf("Collection %s has %d documents to copy", collName, total)

	// Initialize variables for pagination
	var lastID interface{} = ""
	batchNum := 0

	for {
		batchNum++
		query := bson.M{}
		if lastID != "" {
			query["_id"] = bson.M{"$gt": lastID}
		}

		// Find documents in the current batch
		var docs []bson.M
		err := sourceCollection.Find(query).Sort("_id").Limit(batchSize).All(&docs)
		if err != nil {
			return stats, fmt.Errorf("failed to read documents from collection %s: %w", collName, err)
		}

		if len(docs) == 0 {
			break // No more documents to process
		}

		// Update lastID for the next iteration
		lastID = docs[len(docs)-1]["_id"]

		// Insert batch into target collection
		batchSize := 0
		for _, doc := range docs {
			// Estimate document size (rough approximation)
			docBytes, err := bson.Marshal(doc)
			if err == nil {
				batchSize += len(docBytes)
				stats.BytesCopied += int64(len(docBytes))
			}

			err = targetCollection.Insert(doc)
			if err != nil {
				return stats, fmt.Errorf("failed to insert document into target collection %s: %w", collName, err)
			}
			stats.DocumentsCopied++
		}

		logger.Printf("Collection %s: Copied batch %d with %d documents (%.2f%% complete)",
			collName, batchNum, len(docs), float64(stats.DocumentsCopied)*100/float64(total))

		// Small pause to avoid overwhelming the servers
		time.Sleep(10 * time.Millisecond)
	}

	logger.Printf("Completed copying collection %s: %d documents copied", collName, stats.DocumentsCopied)
	return stats, nil
}
