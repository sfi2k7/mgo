package mgo

import (
	"github.com/sfi2k7/mgo/bson"
)

type utilities struct {
}

var Utilities utilities

type MoveCollection2Options struct {
	SourceCollection         string
	TargetCollection         string
	Query                    bson.M
	RowLimit                 int
	DeleteCollectionIfExists bool
	BackupExistingCollection bool
}

type MoveDatabaseOptions struct {
	DeleteDatabaseIfExists   bool
	DeleteCollectionIfExists bool
	RowsLimit                int
	BackupExistingCollection bool
	CollectionsToCopy        []string
	CollectionsToSkip        []string
	DatabaseName             string
}

type MoveDatabaseResponse struct {
	Error                     string
	TotalCollections          int
	TotalCollectionCopied     int
	TotalCollectionsWithError []string
	TotalCollectionSkipped    int
}

func (u utilities) CopyDatabaseWithURI(sourceURI, targetURI string, options *MoveDatabaseOptions) *MoveDatabaseResponse {
	s, err := Dial(sourceURI)
	if err != nil {
		return &MoveDatabaseResponse{Error: err.Error()}
	}
	defer s.Close()

	t, err := Dial(targetURI)
	if err != nil {
		return &MoveDatabaseResponse{Error: err.Error()}
	}
	defer t.Close()

	return u.CopyDatabase(s.DB(options.DatabaseName), t.DB(options.DatabaseName), options)
}

func arrayContains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func (u utilities) CopyDatabase(sourceDb *Database, targetDb *Database, options *MoveDatabaseOptions) *MoveDatabaseResponse {
	var response MoveDatabaseResponse

	//get all collection names
	collectionNames, err := sourceDb.CollectionNames()
	if err != nil {
		response.Error = err.Error()
		return &response
	}

	response.TotalCollections = len(collectionNames)

	//Check if database exists
	existingCollectionNames, err := targetDb.CollectionNames()
	if err != nil {
		response.Error = err.Error()
		return &response
	}

	//Drop database if exists
	if len(existingCollectionNames) > 0 && options.DeleteDatabaseIfExists {
		err = targetDb.DropDatabase()
		if err != nil {
			response.Error = err.Error()
			return &response
		}
	}

	//Copy all collections
	for _, collectioName := range collectionNames {
		//Check if collection exists
		n, err := targetDb.C(collectioName).EstimatedCount()
		if err != nil {
			response.Error = err.Error()
			return &response
		}

		if options.CollectionsToSkip != nil {
			if arrayContains(options.CollectionsToSkip, collectioName) {
				response.TotalCollectionSkipped++
				continue
			}
		}

		if options.CollectionsToCopy != nil {
			if !arrayContains(options.CollectionsToCopy, collectioName) {
				continue
			}
		}

		if n > 0 && !options.DeleteCollectionIfExists {
			response.TotalCollectionSkipped++
			continue
		}

		err = u.copyCollection(sourceDb, targetDb, &MoveCollection2Options{
			RowLimit:                 options.RowsLimit,
			SourceCollection:         collectioName,
			Query:                    bson.M{},
			DeleteCollectionIfExists: options.DeleteCollectionIfExists,
			BackupExistingCollection: options.BackupExistingCollection,
		})

		if err != nil {
			response.Error = err.Error()
			return &response
		}

		response.TotalCollectionCopied++
	}

	return nil
}

func (u utilities) copyCollection(sourceDb *Database, targetDb *Database, op *MoveCollection2Options) error {
	if len(op.TargetCollection) == 0 {
		op.TargetCollection = op.SourceCollection
	}

	count, err := sourceDb.C(op.SourceCollection).Find(op.Query).Count()
	if err != nil {
		return err
	}

	if count > 0 && op.DeleteCollectionIfExists {
		err = targetDb.C(op.TargetCollection).DropCollection()
		if err != nil {
			return err
		}
	}

	var limit = 1000
	if op.RowLimit > 0 {
		limit = op.RowLimit
	}

	var skip = 0

	if op.BackupExistingCollection {
		err = targetDb.RenameCollection(op.TargetCollection, op.TargetCollection+"_backup")
		if err != nil {
			return err
		}
	}

	for {
		var all []interface{}
		err = sourceDb.C(op.SourceCollection).Find(op.Query).Skip(skip).Limit(limit).All(&all)
		if err != nil {
			return err
		}

		if len(all) == 0 {
			return nil
		}

		skip += len(all)
		err = targetDb.C(op.TargetCollection).Insert(all...)
		if err != nil {
			return err
		}

		if skip == count {
			return nil
		}
	}
}
