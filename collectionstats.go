package mgo

type CollectionStats struct {
	AvgObjSize   int  `json:"avgObjSize"`
	Capped       bool `json:"capped"`
	Count        int  `json:"count"`
	IndexDetails struct {
		ID struct {
			Lsm struct {
				BloomFilterFalsePositives                                    int `json:"bloom filter false positives"`
				BloomFilterHits                                              int `json:"bloom filter hits"`
				BloomFilterMisses                                            int `json:"bloom filter misses"`
				BloomFilterPagesEvictedFromCache                             int `json:"bloom filter pages evicted from cache"`
				BloomFilterPagesReadIntoCache                                int `json:"bloom filter pages read into cache"`
				BloomFiltersInTheLSMTree                                     int `json:"bloom filters in the LSM tree"`
				ChunksInTheLSMTree                                           int `json:"chunks in the LSM tree"`
				HighestMergeGenerationInTheLSMTree                           int `json:"highest merge generation in the LSM tree"`
				QueriesThatCouldHaveBenefitedFromABloomFilterThatDidNotExist int `json:"queries that could have benefited from a Bloom filter that did not exist"`
				SleepForLSMCheckpointThrottle                                int `json:"sleep for LSM checkpoint throttle"`
				SleepForLSMMergeThrottle                                     int `json:"sleep for LSM merge throttle"`
				TotalSizeOfBloomFilters                                      int `json:"total size of bloom filters"`
			} `json:"LSM"`
			BlockManager struct {
				AllocationsRequiringFileExtension int `json:"allocations requiring file extension"`
				BlocksAllocated                   int `json:"blocks allocated"`
				BlocksFreed                       int `json:"blocks freed"`
				CheckpointSize                    int `json:"checkpoint size"`
				FileAllocationUnitSize            int `json:"file allocation unit size"`
				FileBytesAvailableForReuse        int `json:"file bytes available for reuse"`
				FileMagicNumber                   int `json:"file magic number"`
				FileMajorVersionNumber            int `json:"file major version number"`
				FileSizeInBytes                   int `json:"file size in bytes"`
				MinorVersionNumber                int `json:"minor version number"`
			} `json:"block-manager"`
			Btree struct {
				BtreeCheckpointGeneration               int `json:"btree checkpoint generation"`
				ColumnStoreFixedSizeLeafPages           int `json:"column-store fixed-size leaf pages"`
				ColumnStoreInternalPages                int `json:"column-store internal pages"`
				ColumnStoreVariableSizeRLEEncodedValues int `json:"column-store variable-size RLE encoded values"`
				ColumnStoreVariableSizeDeletedValues    int `json:"column-store variable-size deleted values"`
				ColumnStoreVariableSizeLeafPages        int `json:"column-store variable-size leaf pages"`
				FixedRecordSize                         int `json:"fixed-record size"`
				MaximumInternalPageKeySize              int `json:"maximum internal page key size"`
				MaximumInternalPageSize                 int `json:"maximum internal page size"`
				MaximumLeafPageKeySize                  int `json:"maximum leaf page key size"`
				MaximumLeafPageSize                     int `json:"maximum leaf page size"`
				MaximumLeafPageValueSize                int `json:"maximum leaf page value size"`
				MaximumTreeDepth                        int `json:"maximum tree depth"`
				NumberOfKeyValuePairs                   int `json:"number of key/value pairs"`
				OverflowPages                           int `json:"overflow pages"`
				PagesRewrittenByCompaction              int `json:"pages rewritten by compaction"`
				RowStoreInternalPages                   int `json:"row-store internal pages"`
				RowStoreLeafPages                       int `json:"row-store leaf pages"`
			} `json:"btree"`
			Cache struct {
				BytesCurrentlyInTheCache                                              int `json:"bytes currently in the cache"`
				BytesReadIntoCache                                                    int `json:"bytes read into cache"`
				BytesWrittenFromCache                                                 int `json:"bytes written from cache"`
				CheckpointBlockedPageEviction                                         int `json:"checkpoint blocked page eviction"`
				DataSourcePagesSelectedForEvictionUnableToBeEvicted                   int `json:"data source pages selected for eviction unable to be evicted"`
				EvictionWalkPassesOfAFile                                             int `json:"eviction walk passes of a file"`
				EvictionWalkTargetPagesHistogram09                                    int `json:"eviction walk target pages histogram - 0-9"`
				EvictionWalkTargetPagesHistogram1031                                  int `json:"eviction walk target pages histogram - 10-31"`
				EvictionWalkTargetPagesHistogram128AndHigher                          int `json:"eviction walk target pages histogram - 128 and higher"`
				EvictionWalkTargetPagesHistogram3263                                  int `json:"eviction walk target pages histogram - 32-63"`
				EvictionWalkTargetPagesHistogram64128                                 int `json:"eviction walk target pages histogram - 64-128"`
				EvictionWalksAbandoned                                                int `json:"eviction walks abandoned"`
				EvictionWalksGaveUpBecauseTheyRestartedTheirWalkTwice                 int `json:"eviction walks gave up because they restarted their walk twice"`
				EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundNoCandidates     int `json:"eviction walks gave up because they saw too many pages and found no candidates"`
				EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundTooFewCandidates int `json:"eviction walks gave up because they saw too many pages and found too few candidates"`
				EvictionWalksReachedEndOfTree                                         int `json:"eviction walks reached end of tree"`
				EvictionWalksStartedFromRootOfTree                                    int `json:"eviction walks started from root of tree"`
				EvictionWalksStartedFromSavedLocationInTree                           int `json:"eviction walks started from saved location in tree"`
				HazardPointerBlockedPageEviction                                      int `json:"hazard pointer blocked page eviction"`
				InMemoryPagePassedCriteriaToBeSplit                                   int `json:"in-memory page passed criteria to be split"`
				InMemoryPageSplits                                                    int `json:"in-memory page splits"`
				InternalPagesEvicted                                                  int `json:"internal pages evicted"`
				InternalPagesSplitDuringEviction                                      int `json:"internal pages split during eviction"`
				LeafPagesSplitDuringEviction                                          int `json:"leaf pages split during eviction"`
				ModifiedPagesEvicted                                                  int `json:"modified pages evicted"`
				OverflowPagesReadIntoCache                                            int `json:"overflow pages read into cache"`
				PageSplitDuringEvictionDeepenedTheTree                                int `json:"page split during eviction deepened the tree"`
				PageWrittenRequiringCacheOverflowRecords                              int `json:"page written requiring cache overflow records"`
				PagesReadIntoCache                                                    int `json:"pages read into cache"`
				PagesReadIntoCacheAfterTruncate                                       int `json:"pages read into cache after truncate"`
				PagesReadIntoCacheAfterTruncateInPrepareState                         int `json:"pages read into cache after truncate in prepare state"`
				PagesReadIntoCacheRequiringCacheOverflowEntries                       int `json:"pages read into cache requiring cache overflow entries"`
				PagesRequestedFromTheCache                                            int `json:"pages requested from the cache"`
				PagesSeenByEvictionWalk                                               int `json:"pages seen by eviction walk"`
				PagesWrittenFromCache                                                 int `json:"pages written from cache"`
				PagesWrittenRequiringInMemoryRestoration                              int `json:"pages written requiring in-memory restoration"`
				TrackedDirtyBytesInTheCache                                           int `json:"tracked dirty bytes in the cache"`
				UnmodifiedPagesEvicted                                                int `json:"unmodified pages evicted"`
			} `json:"cache"`
			CacheWalk struct {
				AverageDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Average difference between current eviction generation when the page was last considered"`
				AverageOnDiskPageImageSizeSeen                                                int `json:"Average on-disk page image size seen"`
				AverageTimeInCacheForPagesThatHaveBeenVisitedByTheEvictionServer              int `json:"Average time in cache for pages that have been visited by the eviction server"`
				AverageTimeInCacheForPagesThatHaveNotBeenVisitedByTheEvictionServer           int `json:"Average time in cache for pages that have not been visited by the eviction server"`
				CleanPagesCurrentlyInCache                                                    int `json:"Clean pages currently in cache"`
				CurrentEvictionGeneration                                                     int `json:"Current eviction generation"`
				DirtyPagesCurrentlyInCache                                                    int `json:"Dirty pages currently in cache"`
				EntriesInTheRootPage                                                          int `json:"Entries in the root page"`
				InternalPagesCurrentlyInCache                                                 int `json:"Internal pages currently in cache"`
				LeafPagesCurrentlyInCache                                                     int `json:"Leaf pages currently in cache"`
				MaximumDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Maximum difference between current eviction generation when the page was last considered"`
				MaximumPageSizeSeen                                                           int `json:"Maximum page size seen"`
				MinimumOnDiskPageImageSizeSeen                                                int `json:"Minimum on-disk page image size seen"`
				NumberOfPagesNeverVisitedByEvictionServer                                     int `json:"Number of pages never visited by eviction server"`
				OnDiskPageImageSizesSmallerThanASingleAllocationUnit                          int `json:"On-disk page image sizes smaller than a single allocation unit"`
				PagesCreatedInMemoryAndNeverWritten                                           int `json:"Pages created in memory and never written"`
				PagesCurrentlyQueuedForEviction                                               int `json:"Pages currently queued for eviction"`
				PagesThatCouldNotBeQueuedForEviction                                          int `json:"Pages that could not be queued for eviction"`
				RefsSkippedDuringCacheTraversal                                               int `json:"Refs skipped during cache traversal"`
				SizeOfTheRootPage                                                             int `json:"Size of the root page"`
				TotalNumberOfPagesCurrentlyInCache                                            int `json:"Total number of pages currently in cache"`
			} `json:"cache_walk"`
			Compression struct {
				CompressedPagesRead                               int `json:"compressed pages read"`
				CompressedPagesWritten                            int `json:"compressed pages written"`
				PageWrittenFailedToCompress                       int `json:"page written failed to compress"`
				PageWrittenWasTooSmallToCompress                  int `json:"page written was too small to compress"`
				RawCompressionCallFailedAdditionalDataAvailable   int `json:"raw compression call failed, additional data available"`
				RawCompressionCallFailedNoAdditionalDataAvailable int `json:"raw compression call failed, no additional data available"`
				RawCompressionCallSucceeded                       int `json:"raw compression call succeeded"`
			} `json:"compression"`
			CreationString string `json:"creationString"`
			Cursor         struct {
				BulkLoadedCursorInsertCalls          int `json:"bulk-loaded cursor-insert calls"`
				CloseCallsThatResultInCache          int `json:"close calls that result in cache"`
				CreateCalls                          int `json:"create calls"`
				CursorOperationRestarted             int `json:"cursor operation restarted"`
				CursorInsertKeyAndValueBytesInserted int `json:"cursor-insert key and value bytes inserted"`
				CursorRemoveKeyBytesRemoved          int `json:"cursor-remove key bytes removed"`
				CursorUpdateValueBytesUpdated        int `json:"cursor-update value bytes updated"`
				CursorsReusedFromCache               int `json:"cursors reused from cache"`
				InsertCalls                          int `json:"insert calls"`
				ModifyCalls                          int `json:"modify calls"`
				NextCalls                            int `json:"next calls"`
				OpenCursorCount                      int `json:"open cursor count"`
				PrevCalls                            int `json:"prev calls"`
				RemoveCalls                          int `json:"remove calls"`
				ReserveCalls                         int `json:"reserve calls"`
				ResetCalls                           int `json:"reset calls"`
				SearchCalls                          int `json:"search calls"`
				SearchNearCalls                      int `json:"search near calls"`
				TruncateCalls                        int `json:"truncate calls"`
				UpdateCalls                          int `json:"update calls"`
			} `json:"cursor"`
			Metadata struct {
				FormatVersion int    `json:"formatVersion"`
				InfoObj       string `json:"infoObj"`
			} `json:"metadata"`
			Reconciliation struct {
				DictionaryMatches                                   int `json:"dictionary matches"`
				FastPathPagesDeleted                                int `json:"fast-path pages deleted"`
				InternalPageKeyBytesDiscardedUsingSuffixCompression int `json:"internal page key bytes discarded using suffix compression"`
				InternalPageMultiBlockWrites                        int `json:"internal page multi-block writes"`
				InternalPageOverflowKeys                            int `json:"internal-page overflow keys"`
				LeafPageKeyBytesDiscardedUsingPrefixCompression     int `json:"leaf page key bytes discarded using prefix compression"`
				LeafPageMultiBlockWrites                            int `json:"leaf page multi-block writes"`
				LeafPageOverflowKeys                                int `json:"leaf-page overflow keys"`
				MaximumBlocksRequiredForAPage                       int `json:"maximum blocks required for a page"`
				OverflowValuesWritten                               int `json:"overflow values written"`
				PageChecksumMatches                                 int `json:"page checksum matches"`
				PageReconciliationCalls                             int `json:"page reconciliation calls"`
				PageReconciliationCallsForEviction                  int `json:"page reconciliation calls for eviction"`
				PagesDeleted                                        int `json:"pages deleted"`
			} `json:"reconciliation"`
			Session struct {
				ObjectCompaction int `json:"object compaction"`
			} `json:"session"`
			Transaction struct {
				UpdateConflicts int `json:"update conflicts"`
			} `json:"transaction"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"_id_"`
		FirstNameLastNameText struct {
			Lsm struct {
				BloomFilterFalsePositives                                    int `json:"bloom filter false positives"`
				BloomFilterHits                                              int `json:"bloom filter hits"`
				BloomFilterMisses                                            int `json:"bloom filter misses"`
				BloomFilterPagesEvictedFromCache                             int `json:"bloom filter pages evicted from cache"`
				BloomFilterPagesReadIntoCache                                int `json:"bloom filter pages read into cache"`
				BloomFiltersInTheLSMTree                                     int `json:"bloom filters in the LSM tree"`
				ChunksInTheLSMTree                                           int `json:"chunks in the LSM tree"`
				HighestMergeGenerationInTheLSMTree                           int `json:"highest merge generation in the LSM tree"`
				QueriesThatCouldHaveBenefitedFromABloomFilterThatDidNotExist int `json:"queries that could have benefited from a Bloom filter that did not exist"`
				SleepForLSMCheckpointThrottle                                int `json:"sleep for LSM checkpoint throttle"`
				SleepForLSMMergeThrottle                                     int `json:"sleep for LSM merge throttle"`
				TotalSizeOfBloomFilters                                      int `json:"total size of bloom filters"`
			} `json:"LSM"`
			BlockManager struct {
				AllocationsRequiringFileExtension int `json:"allocations requiring file extension"`
				BlocksAllocated                   int `json:"blocks allocated"`
				BlocksFreed                       int `json:"blocks freed"`
				CheckpointSize                    int `json:"checkpoint size"`
				FileAllocationUnitSize            int `json:"file allocation unit size"`
				FileBytesAvailableForReuse        int `json:"file bytes available for reuse"`
				FileMagicNumber                   int `json:"file magic number"`
				FileMajorVersionNumber            int `json:"file major version number"`
				FileSizeInBytes                   int `json:"file size in bytes"`
				MinorVersionNumber                int `json:"minor version number"`
			} `json:"block-manager"`
			Btree struct {
				BtreeCheckpointGeneration               int `json:"btree checkpoint generation"`
				ColumnStoreFixedSizeLeafPages           int `json:"column-store fixed-size leaf pages"`
				ColumnStoreInternalPages                int `json:"column-store internal pages"`
				ColumnStoreVariableSizeRLEEncodedValues int `json:"column-store variable-size RLE encoded values"`
				ColumnStoreVariableSizeDeletedValues    int `json:"column-store variable-size deleted values"`
				ColumnStoreVariableSizeLeafPages        int `json:"column-store variable-size leaf pages"`
				FixedRecordSize                         int `json:"fixed-record size"`
				MaximumInternalPageKeySize              int `json:"maximum internal page key size"`
				MaximumInternalPageSize                 int `json:"maximum internal page size"`
				MaximumLeafPageKeySize                  int `json:"maximum leaf page key size"`
				MaximumLeafPageSize                     int `json:"maximum leaf page size"`
				MaximumLeafPageValueSize                int `json:"maximum leaf page value size"`
				MaximumTreeDepth                        int `json:"maximum tree depth"`
				NumberOfKeyValuePairs                   int `json:"number of key/value pairs"`
				OverflowPages                           int `json:"overflow pages"`
				PagesRewrittenByCompaction              int `json:"pages rewritten by compaction"`
				RowStoreInternalPages                   int `json:"row-store internal pages"`
				RowStoreLeafPages                       int `json:"row-store leaf pages"`
			} `json:"btree"`
			Cache struct {
				BytesCurrentlyInTheCache                                              int `json:"bytes currently in the cache"`
				BytesReadIntoCache                                                    int `json:"bytes read into cache"`
				BytesWrittenFromCache                                                 int `json:"bytes written from cache"`
				CheckpointBlockedPageEviction                                         int `json:"checkpoint blocked page eviction"`
				DataSourcePagesSelectedForEvictionUnableToBeEvicted                   int `json:"data source pages selected for eviction unable to be evicted"`
				EvictionWalkPassesOfAFile                                             int `json:"eviction walk passes of a file"`
				EvictionWalkTargetPagesHistogram09                                    int `json:"eviction walk target pages histogram - 0-9"`
				EvictionWalkTargetPagesHistogram1031                                  int `json:"eviction walk target pages histogram - 10-31"`
				EvictionWalkTargetPagesHistogram128AndHigher                          int `json:"eviction walk target pages histogram - 128 and higher"`
				EvictionWalkTargetPagesHistogram3263                                  int `json:"eviction walk target pages histogram - 32-63"`
				EvictionWalkTargetPagesHistogram64128                                 int `json:"eviction walk target pages histogram - 64-128"`
				EvictionWalksAbandoned                                                int `json:"eviction walks abandoned"`
				EvictionWalksGaveUpBecauseTheyRestartedTheirWalkTwice                 int `json:"eviction walks gave up because they restarted their walk twice"`
				EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundNoCandidates     int `json:"eviction walks gave up because they saw too many pages and found no candidates"`
				EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundTooFewCandidates int `json:"eviction walks gave up because they saw too many pages and found too few candidates"`
				EvictionWalksReachedEndOfTree                                         int `json:"eviction walks reached end of tree"`
				EvictionWalksStartedFromRootOfTree                                    int `json:"eviction walks started from root of tree"`
				EvictionWalksStartedFromSavedLocationInTree                           int `json:"eviction walks started from saved location in tree"`
				HazardPointerBlockedPageEviction                                      int `json:"hazard pointer blocked page eviction"`
				InMemoryPagePassedCriteriaToBeSplit                                   int `json:"in-memory page passed criteria to be split"`
				InMemoryPageSplits                                                    int `json:"in-memory page splits"`
				InternalPagesEvicted                                                  int `json:"internal pages evicted"`
				InternalPagesSplitDuringEviction                                      int `json:"internal pages split during eviction"`
				LeafPagesSplitDuringEviction                                          int `json:"leaf pages split during eviction"`
				ModifiedPagesEvicted                                                  int `json:"modified pages evicted"`
				OverflowPagesReadIntoCache                                            int `json:"overflow pages read into cache"`
				PageSplitDuringEvictionDeepenedTheTree                                int `json:"page split during eviction deepened the tree"`
				PageWrittenRequiringCacheOverflowRecords                              int `json:"page written requiring cache overflow records"`
				PagesReadIntoCache                                                    int `json:"pages read into cache"`
				PagesReadIntoCacheAfterTruncate                                       int `json:"pages read into cache after truncate"`
				PagesReadIntoCacheAfterTruncateInPrepareState                         int `json:"pages read into cache after truncate in prepare state"`
				PagesReadIntoCacheRequiringCacheOverflowEntries                       int `json:"pages read into cache requiring cache overflow entries"`
				PagesRequestedFromTheCache                                            int `json:"pages requested from the cache"`
				PagesSeenByEvictionWalk                                               int `json:"pages seen by eviction walk"`
				PagesWrittenFromCache                                                 int `json:"pages written from cache"`
				PagesWrittenRequiringInMemoryRestoration                              int `json:"pages written requiring in-memory restoration"`
				TrackedDirtyBytesInTheCache                                           int `json:"tracked dirty bytes in the cache"`
				UnmodifiedPagesEvicted                                                int `json:"unmodified pages evicted"`
			} `json:"cache"`
			CacheWalk struct {
				AverageDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Average difference between current eviction generation when the page was last considered"`
				AverageOnDiskPageImageSizeSeen                                                int `json:"Average on-disk page image size seen"`
				AverageTimeInCacheForPagesThatHaveBeenVisitedByTheEvictionServer              int `json:"Average time in cache for pages that have been visited by the eviction server"`
				AverageTimeInCacheForPagesThatHaveNotBeenVisitedByTheEvictionServer           int `json:"Average time in cache for pages that have not been visited by the eviction server"`
				CleanPagesCurrentlyInCache                                                    int `json:"Clean pages currently in cache"`
				CurrentEvictionGeneration                                                     int `json:"Current eviction generation"`
				DirtyPagesCurrentlyInCache                                                    int `json:"Dirty pages currently in cache"`
				EntriesInTheRootPage                                                          int `json:"Entries in the root page"`
				InternalPagesCurrentlyInCache                                                 int `json:"Internal pages currently in cache"`
				LeafPagesCurrentlyInCache                                                     int `json:"Leaf pages currently in cache"`
				MaximumDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Maximum difference between current eviction generation when the page was last considered"`
				MaximumPageSizeSeen                                                           int `json:"Maximum page size seen"`
				MinimumOnDiskPageImageSizeSeen                                                int `json:"Minimum on-disk page image size seen"`
				NumberOfPagesNeverVisitedByEvictionServer                                     int `json:"Number of pages never visited by eviction server"`
				OnDiskPageImageSizesSmallerThanASingleAllocationUnit                          int `json:"On-disk page image sizes smaller than a single allocation unit"`
				PagesCreatedInMemoryAndNeverWritten                                           int `json:"Pages created in memory and never written"`
				PagesCurrentlyQueuedForEviction                                               int `json:"Pages currently queued for eviction"`
				PagesThatCouldNotBeQueuedForEviction                                          int `json:"Pages that could not be queued for eviction"`
				RefsSkippedDuringCacheTraversal                                               int `json:"Refs skipped during cache traversal"`
				SizeOfTheRootPage                                                             int `json:"Size of the root page"`
				TotalNumberOfPagesCurrentlyInCache                                            int `json:"Total number of pages currently in cache"`
			} `json:"cache_walk"`
			Compression struct {
				CompressedPagesRead                               int `json:"compressed pages read"`
				CompressedPagesWritten                            int `json:"compressed pages written"`
				PageWrittenFailedToCompress                       int `json:"page written failed to compress"`
				PageWrittenWasTooSmallToCompress                  int `json:"page written was too small to compress"`
				RawCompressionCallFailedAdditionalDataAvailable   int `json:"raw compression call failed, additional data available"`
				RawCompressionCallFailedNoAdditionalDataAvailable int `json:"raw compression call failed, no additional data available"`
				RawCompressionCallSucceeded                       int `json:"raw compression call succeeded"`
			} `json:"compression"`
			CreationString string `json:"creationString"`
			Cursor         struct {
				BulkLoadedCursorInsertCalls          int `json:"bulk-loaded cursor-insert calls"`
				CloseCallsThatResultInCache          int `json:"close calls that result in cache"`
				CreateCalls                          int `json:"create calls"`
				CursorOperationRestarted             int `json:"cursor operation restarted"`
				CursorInsertKeyAndValueBytesInserted int `json:"cursor-insert key and value bytes inserted"`
				CursorRemoveKeyBytesRemoved          int `json:"cursor-remove key bytes removed"`
				CursorUpdateValueBytesUpdated        int `json:"cursor-update value bytes updated"`
				CursorsReusedFromCache               int `json:"cursors reused from cache"`
				InsertCalls                          int `json:"insert calls"`
				ModifyCalls                          int `json:"modify calls"`
				NextCalls                            int `json:"next calls"`
				OpenCursorCount                      int `json:"open cursor count"`
				PrevCalls                            int `json:"prev calls"`
				RemoveCalls                          int `json:"remove calls"`
				ReserveCalls                         int `json:"reserve calls"`
				ResetCalls                           int `json:"reset calls"`
				SearchCalls                          int `json:"search calls"`
				SearchNearCalls                      int `json:"search near calls"`
				TruncateCalls                        int `json:"truncate calls"`
				UpdateCalls                          int `json:"update calls"`
			} `json:"cursor"`
			Metadata struct {
				FormatVersion int    `json:"formatVersion"`
				InfoObj       string `json:"infoObj"`
			} `json:"metadata"`
			Reconciliation struct {
				DictionaryMatches                                   int `json:"dictionary matches"`
				FastPathPagesDeleted                                int `json:"fast-path pages deleted"`
				InternalPageKeyBytesDiscardedUsingSuffixCompression int `json:"internal page key bytes discarded using suffix compression"`
				InternalPageMultiBlockWrites                        int `json:"internal page multi-block writes"`
				InternalPageOverflowKeys                            int `json:"internal-page overflow keys"`
				LeafPageKeyBytesDiscardedUsingPrefixCompression     int `json:"leaf page key bytes discarded using prefix compression"`
				LeafPageMultiBlockWrites                            int `json:"leaf page multi-block writes"`
				LeafPageOverflowKeys                                int `json:"leaf-page overflow keys"`
				MaximumBlocksRequiredForAPage                       int `json:"maximum blocks required for a page"`
				OverflowValuesWritten                               int `json:"overflow values written"`
				PageChecksumMatches                                 int `json:"page checksum matches"`
				PageReconciliationCalls                             int `json:"page reconciliation calls"`
				PageReconciliationCallsForEviction                  int `json:"page reconciliation calls for eviction"`
				PagesDeleted                                        int `json:"pages deleted"`
			} `json:"reconciliation"`
			Session struct {
				ObjectCompaction int `json:"object compaction"`
			} `json:"session"`
			Transaction struct {
				UpdateConflicts int `json:"update conflicts"`
			} `json:"transaction"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"first_name_last_name_text"`
	} `json:"indexDetails"`
	IndexSizes struct {
		ID                    int `json:"_id_"`
		FirstNameLastNameText int `json:"first_name_last_name_text"`
	} `json:"indexSizes"`
	Nindexes       int    `json:"nindexes"`
	Ns             string `json:"ns"`
	Ok             int    `json:"ok"`
	Size           int    `json:"size"`
	StorageSize    int    `json:"storageSize"`
	TotalIndexSize int    `json:"totalIndexSize"`
	WiredTiger     struct {
		Lsm struct {
			BloomFilterFalsePositives                                    int `json:"bloom filter false positives"`
			BloomFilterHits                                              int `json:"bloom filter hits"`
			BloomFilterMisses                                            int `json:"bloom filter misses"`
			BloomFilterPagesEvictedFromCache                             int `json:"bloom filter pages evicted from cache"`
			BloomFilterPagesReadIntoCache                                int `json:"bloom filter pages read into cache"`
			BloomFiltersInTheLSMTree                                     int `json:"bloom filters in the LSM tree"`
			ChunksInTheLSMTree                                           int `json:"chunks in the LSM tree"`
			HighestMergeGenerationInTheLSMTree                           int `json:"highest merge generation in the LSM tree"`
			QueriesThatCouldHaveBenefitedFromABloomFilterThatDidNotExist int `json:"queries that could have benefited from a Bloom filter that did not exist"`
			SleepForLSMCheckpointThrottle                                int `json:"sleep for LSM checkpoint throttle"`
			SleepForLSMMergeThrottle                                     int `json:"sleep for LSM merge throttle"`
			TotalSizeOfBloomFilters                                      int `json:"total size of bloom filters"`
		} `json:"LSM"`
		BlockManager struct {
			AllocationsRequiringFileExtension int `json:"allocations requiring file extension"`
			BlocksAllocated                   int `json:"blocks allocated"`
			BlocksFreed                       int `json:"blocks freed"`
			CheckpointSize                    int `json:"checkpoint size"`
			FileAllocationUnitSize            int `json:"file allocation unit size"`
			FileBytesAvailableForReuse        int `json:"file bytes available for reuse"`
			FileMagicNumber                   int `json:"file magic number"`
			FileMajorVersionNumber            int `json:"file major version number"`
			FileSizeInBytes                   int `json:"file size in bytes"`
			MinorVersionNumber                int `json:"minor version number"`
		} `json:"block-manager"`
		Btree struct {
			BtreeCheckpointGeneration               int `json:"btree checkpoint generation"`
			ColumnStoreFixedSizeLeafPages           int `json:"column-store fixed-size leaf pages"`
			ColumnStoreInternalPages                int `json:"column-store internal pages"`
			ColumnStoreVariableSizeRLEEncodedValues int `json:"column-store variable-size RLE encoded values"`
			ColumnStoreVariableSizeDeletedValues    int `json:"column-store variable-size deleted values"`
			ColumnStoreVariableSizeLeafPages        int `json:"column-store variable-size leaf pages"`
			FixedRecordSize                         int `json:"fixed-record size"`
			MaximumInternalPageKeySize              int `json:"maximum internal page key size"`
			MaximumInternalPageSize                 int `json:"maximum internal page size"`
			MaximumLeafPageKeySize                  int `json:"maximum leaf page key size"`
			MaximumLeafPageSize                     int `json:"maximum leaf page size"`
			MaximumLeafPageValueSize                int `json:"maximum leaf page value size"`
			MaximumTreeDepth                        int `json:"maximum tree depth"`
			NumberOfKeyValuePairs                   int `json:"number of key/value pairs"`
			OverflowPages                           int `json:"overflow pages"`
			PagesRewrittenByCompaction              int `json:"pages rewritten by compaction"`
			RowStoreInternalPages                   int `json:"row-store internal pages"`
			RowStoreLeafPages                       int `json:"row-store leaf pages"`
		} `json:"btree"`
		Cache struct {
			BytesCurrentlyInTheCache                                              int `json:"bytes currently in the cache"`
			BytesReadIntoCache                                                    int `json:"bytes read into cache"`
			BytesWrittenFromCache                                                 int `json:"bytes written from cache"`
			CheckpointBlockedPageEviction                                         int `json:"checkpoint blocked page eviction"`
			DataSourcePagesSelectedForEvictionUnableToBeEvicted                   int `json:"data source pages selected for eviction unable to be evicted"`
			EvictionWalkPassesOfAFile                                             int `json:"eviction walk passes of a file"`
			EvictionWalkTargetPagesHistogram09                                    int `json:"eviction walk target pages histogram - 0-9"`
			EvictionWalkTargetPagesHistogram1031                                  int `json:"eviction walk target pages histogram - 10-31"`
			EvictionWalkTargetPagesHistogram128AndHigher                          int `json:"eviction walk target pages histogram - 128 and higher"`
			EvictionWalkTargetPagesHistogram3263                                  int `json:"eviction walk target pages histogram - 32-63"`
			EvictionWalkTargetPagesHistogram64128                                 int `json:"eviction walk target pages histogram - 64-128"`
			EvictionWalksAbandoned                                                int `json:"eviction walks abandoned"`
			EvictionWalksGaveUpBecauseTheyRestartedTheirWalkTwice                 int `json:"eviction walks gave up because they restarted their walk twice"`
			EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundNoCandidates     int `json:"eviction walks gave up because they saw too many pages and found no candidates"`
			EvictionWalksGaveUpBecauseTheySawTooManyPagesAndFoundTooFewCandidates int `json:"eviction walks gave up because they saw too many pages and found too few candidates"`
			EvictionWalksReachedEndOfTree                                         int `json:"eviction walks reached end of tree"`
			EvictionWalksStartedFromRootOfTree                                    int `json:"eviction walks started from root of tree"`
			EvictionWalksStartedFromSavedLocationInTree                           int `json:"eviction walks started from saved location in tree"`
			HazardPointerBlockedPageEviction                                      int `json:"hazard pointer blocked page eviction"`
			InMemoryPagePassedCriteriaToBeSplit                                   int `json:"in-memory page passed criteria to be split"`
			InMemoryPageSplits                                                    int `json:"in-memory page splits"`
			InternalPagesEvicted                                                  int `json:"internal pages evicted"`
			InternalPagesSplitDuringEviction                                      int `json:"internal pages split during eviction"`
			LeafPagesSplitDuringEviction                                          int `json:"leaf pages split during eviction"`
			ModifiedPagesEvicted                                                  int `json:"modified pages evicted"`
			OverflowPagesReadIntoCache                                            int `json:"overflow pages read into cache"`
			PageSplitDuringEvictionDeepenedTheTree                                int `json:"page split during eviction deepened the tree"`
			PageWrittenRequiringCacheOverflowRecords                              int `json:"page written requiring cache overflow records"`
			PagesReadIntoCache                                                    int `json:"pages read into cache"`
			PagesReadIntoCacheAfterTruncate                                       int `json:"pages read into cache after truncate"`
			PagesReadIntoCacheAfterTruncateInPrepareState                         int `json:"pages read into cache after truncate in prepare state"`
			PagesReadIntoCacheRequiringCacheOverflowEntries                       int `json:"pages read into cache requiring cache overflow entries"`
			PagesRequestedFromTheCache                                            int `json:"pages requested from the cache"`
			PagesSeenByEvictionWalk                                               int `json:"pages seen by eviction walk"`
			PagesWrittenFromCache                                                 int `json:"pages written from cache"`
			PagesWrittenRequiringInMemoryRestoration                              int `json:"pages written requiring in-memory restoration"`
			TrackedDirtyBytesInTheCache                                           int `json:"tracked dirty bytes in the cache"`
			UnmodifiedPagesEvicted                                                int `json:"unmodified pages evicted"`
		} `json:"cache"`
		CacheWalk struct {
			AverageDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Average difference between current eviction generation when the page was last considered"`
			AverageOnDiskPageImageSizeSeen                                                int `json:"Average on-disk page image size seen"`
			AverageTimeInCacheForPagesThatHaveBeenVisitedByTheEvictionServer              int `json:"Average time in cache for pages that have been visited by the eviction server"`
			AverageTimeInCacheForPagesThatHaveNotBeenVisitedByTheEvictionServer           int `json:"Average time in cache for pages that have not been visited by the eviction server"`
			CleanPagesCurrentlyInCache                                                    int `json:"Clean pages currently in cache"`
			CurrentEvictionGeneration                                                     int `json:"Current eviction generation"`
			DirtyPagesCurrentlyInCache                                                    int `json:"Dirty pages currently in cache"`
			EntriesInTheRootPage                                                          int `json:"Entries in the root page"`
			InternalPagesCurrentlyInCache                                                 int `json:"Internal pages currently in cache"`
			LeafPagesCurrentlyInCache                                                     int `json:"Leaf pages currently in cache"`
			MaximumDifferenceBetweenCurrentEvictionGenerationWhenThePageWasLastConsidered int `json:"Maximum difference between current eviction generation when the page was last considered"`
			MaximumPageSizeSeen                                                           int `json:"Maximum page size seen"`
			MinimumOnDiskPageImageSizeSeen                                                int `json:"Minimum on-disk page image size seen"`
			NumberOfPagesNeverVisitedByEvictionServer                                     int `json:"Number of pages never visited by eviction server"`
			OnDiskPageImageSizesSmallerThanASingleAllocationUnit                          int `json:"On-disk page image sizes smaller than a single allocation unit"`
			PagesCreatedInMemoryAndNeverWritten                                           int `json:"Pages created in memory and never written"`
			PagesCurrentlyQueuedForEviction                                               int `json:"Pages currently queued for eviction"`
			PagesThatCouldNotBeQueuedForEviction                                          int `json:"Pages that could not be queued for eviction"`
			RefsSkippedDuringCacheTraversal                                               int `json:"Refs skipped during cache traversal"`
			SizeOfTheRootPage                                                             int `json:"Size of the root page"`
			TotalNumberOfPagesCurrentlyInCache                                            int `json:"Total number of pages currently in cache"`
		} `json:"cache_walk"`
		Compression struct {
			CompressedPagesRead                               int `json:"compressed pages read"`
			CompressedPagesWritten                            int `json:"compressed pages written"`
			PageWrittenFailedToCompress                       int `json:"page written failed to compress"`
			PageWrittenWasTooSmallToCompress                  int `json:"page written was too small to compress"`
			RawCompressionCallFailedAdditionalDataAvailable   int `json:"raw compression call failed, additional data available"`
			RawCompressionCallFailedNoAdditionalDataAvailable int `json:"raw compression call failed, no additional data available"`
			RawCompressionCallSucceeded                       int `json:"raw compression call succeeded"`
		} `json:"compression"`
		CreationString string `json:"creationString"`
		Cursor         struct {
			BulkLoadedCursorInsertCalls          int `json:"bulk-loaded cursor-insert calls"`
			CloseCallsThatResultInCache          int `json:"close calls that result in cache"`
			CreateCalls                          int `json:"create calls"`
			CursorOperationRestarted             int `json:"cursor operation restarted"`
			CursorInsertKeyAndValueBytesInserted int `json:"cursor-insert key and value bytes inserted"`
			CursorRemoveKeyBytesRemoved          int `json:"cursor-remove key bytes removed"`
			CursorUpdateValueBytesUpdated        int `json:"cursor-update value bytes updated"`
			CursorsReusedFromCache               int `json:"cursors reused from cache"`
			InsertCalls                          int `json:"insert calls"`
			ModifyCalls                          int `json:"modify calls"`
			NextCalls                            int `json:"next calls"`
			OpenCursorCount                      int `json:"open cursor count"`
			PrevCalls                            int `json:"prev calls"`
			RemoveCalls                          int `json:"remove calls"`
			ReserveCalls                         int `json:"reserve calls"`
			ResetCalls                           int `json:"reset calls"`
			SearchCalls                          int `json:"search calls"`
			SearchNearCalls                      int `json:"search near calls"`
			TruncateCalls                        int `json:"truncate calls"`
			UpdateCalls                          int `json:"update calls"`
		} `json:"cursor"`
		Metadata struct {
			FormatVersion int `json:"formatVersion"`
		} `json:"metadata"`
		Reconciliation struct {
			DictionaryMatches                                   int `json:"dictionary matches"`
			FastPathPagesDeleted                                int `json:"fast-path pages deleted"`
			InternalPageKeyBytesDiscardedUsingSuffixCompression int `json:"internal page key bytes discarded using suffix compression"`
			InternalPageMultiBlockWrites                        int `json:"internal page multi-block writes"`
			InternalPageOverflowKeys                            int `json:"internal-page overflow keys"`
			LeafPageKeyBytesDiscardedUsingPrefixCompression     int `json:"leaf page key bytes discarded using prefix compression"`
			LeafPageMultiBlockWrites                            int `json:"leaf page multi-block writes"`
			LeafPageOverflowKeys                                int `json:"leaf-page overflow keys"`
			MaximumBlocksRequiredForAPage                       int `json:"maximum blocks required for a page"`
			OverflowValuesWritten                               int `json:"overflow values written"`
			PageChecksumMatches                                 int `json:"page checksum matches"`
			PageReconciliationCalls                             int `json:"page reconciliation calls"`
			PageReconciliationCallsForEviction                  int `json:"page reconciliation calls for eviction"`
			PagesDeleted                                        int `json:"pages deleted"`
		} `json:"reconciliation"`
		Session struct {
			ObjectCompaction int `json:"object compaction"`
		} `json:"session"`
		Transaction struct {
			UpdateConflicts int `json:"update conflicts"`
		} `json:"transaction"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"wiredTiger"`
}
