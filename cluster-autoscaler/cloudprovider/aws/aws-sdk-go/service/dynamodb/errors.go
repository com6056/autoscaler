// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBackupInUseException for service response error code
	// "BackupInUseException".
	//
	// There is another ongoing conflicting backup control plane operation on the
	// table. The backup is either being created, deleted or restored to a table.
	ErrCodeBackupInUseException = "BackupInUseException"

	// ErrCodeBackupNotFoundException for service response error code
	// "BackupNotFoundException".
	//
	// Backup not found for the given BackupARN.
	ErrCodeBackupNotFoundException = "BackupNotFoundException"

	// ErrCodeConditionalCheckFailedException for service response error code
	// "ConditionalCheckFailedException".
	//
	// A condition specified in the operation could not be evaluated.
	ErrCodeConditionalCheckFailedException = "ConditionalCheckFailedException"

	// ErrCodeContinuousBackupsUnavailableException for service response error code
	// "ContinuousBackupsUnavailableException".
	//
	// Backups have not yet been enabled for this table.
	ErrCodeContinuousBackupsUnavailableException = "ContinuousBackupsUnavailableException"

	// ErrCodeDuplicateItemException for service response error code
	// "DuplicateItemException".
	//
	// There was an attempt to insert an item with the same primary key as an item
	// that already exists in the DynamoDB table.
	ErrCodeDuplicateItemException = "DuplicateItemException"

	// ErrCodeExportConflictException for service response error code
	// "ExportConflictException".
	//
	// There was a conflict when writing to the specified S3 bucket.
	ErrCodeExportConflictException = "ExportConflictException"

	// ErrCodeExportNotFoundException for service response error code
	// "ExportNotFoundException".
	//
	// The specified export was not found.
	ErrCodeExportNotFoundException = "ExportNotFoundException"

	// ErrCodeGlobalTableAlreadyExistsException for service response error code
	// "GlobalTableAlreadyExistsException".
	//
	// The specified global table already exists.
	ErrCodeGlobalTableAlreadyExistsException = "GlobalTableAlreadyExistsException"

	// ErrCodeGlobalTableNotFoundException for service response error code
	// "GlobalTableNotFoundException".
	//
	// The specified global table does not exist.
	ErrCodeGlobalTableNotFoundException = "GlobalTableNotFoundException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// DynamoDB rejected the request because you retried a request with a different
	// payload but with an idempotent token that was already used.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeImportConflictException for service response error code
	// "ImportConflictException".
	//
	// There was a conflict when importing from the specified S3 source. This can
	// occur when the current import conflicts with a previous import request that
	// had the same client token.
	ErrCodeImportConflictException = "ImportConflictException"

	// ErrCodeImportNotFoundException for service response error code
	// "ImportNotFoundException".
	//
	// The specified import was not found.
	ErrCodeImportNotFoundException = "ImportNotFoundException"

	// ErrCodeIndexNotFoundException for service response error code
	// "IndexNotFoundException".
	//
	// The operation tried to access a nonexistent index.
	ErrCodeIndexNotFoundException = "IndexNotFoundException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// An error occurred on the server side.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidExportTimeException for service response error code
	// "InvalidExportTimeException".
	//
	// The specified ExportTime is outside of the point in time recovery window.
	ErrCodeInvalidExportTimeException = "InvalidExportTimeException"

	// ErrCodeInvalidRestoreTimeException for service response error code
	// "InvalidRestoreTimeException".
	//
	// An invalid restore time was specified. RestoreDateTime must be between EarliestRestorableDateTime
	// and LatestRestorableDateTime.
	ErrCodeInvalidRestoreTimeException = "InvalidRestoreTimeException"

	// ErrCodeItemCollectionSizeLimitExceededException for service response error code
	// "ItemCollectionSizeLimitExceededException".
	//
	// An item collection is too large. This exception is only returned for tables
	// that have one or more local secondary indexes.
	ErrCodeItemCollectionSizeLimitExceededException = "ItemCollectionSizeLimitExceededException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// There is no limit to the number of daily on-demand backups that can be taken.
	//
	// For most purposes, up to 500 simultaneous table operations are allowed per
	// account. These operations include CreateTable, UpdateTable, DeleteTable,UpdateTimeToLive,
	// RestoreTableFromBackup, and RestoreTableToPointInTime.
	//
	// When you are creating a table with one or more secondary indexes, you can
	// have up to 250 such requests running at a time. However, if the table or
	// index specifications are complex, then DynamoDB might temporarily reduce
	// the number of concurrent operations.
	//
	// When importing into DynamoDB, up to 50 simultaneous import table operations
	// are allowed per account.
	//
	// There is a soft account quota of 2,500 tables.
	//
	// GetRecords was called with a value of more than 1000 for the limit request
	// parameter.
	//
	// More than 2 processes are reading from the same streams shard at the same
	// time. Exceeding this limit may result in request throttling.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePointInTimeRecoveryUnavailableException for service response error code
	// "PointInTimeRecoveryUnavailableException".
	//
	// Point in time recovery has not yet been enabled for this source table.
	ErrCodePointInTimeRecoveryUnavailableException = "PointInTimeRecoveryUnavailableException"

	// ErrCodeProvisionedThroughputExceededException for service response error code
	// "ProvisionedThroughputExceededException".
	//
	// Your request rate is too high. The Amazon Web Services SDKs for DynamoDB
	// automatically retry requests that receive this exception. Your request is
	// eventually successful, unless your retry queue is too large to finish. Reduce
	// the frequency of requests and use exponential backoff. For more information,
	// go to Error Retries and Exponential Backoff (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Programming.Errors.html#Programming.Errors.RetryAndBackoff)
	// in the Amazon DynamoDB Developer Guide.
	ErrCodeProvisionedThroughputExceededException = "ProvisionedThroughputExceededException"

	// ErrCodeReplicaAlreadyExistsException for service response error code
	// "ReplicaAlreadyExistsException".
	//
	// The specified replica is already part of the global table.
	ErrCodeReplicaAlreadyExistsException = "ReplicaAlreadyExistsException"

	// ErrCodeReplicaNotFoundException for service response error code
	// "ReplicaNotFoundException".
	//
	// The specified replica is no longer part of the global table.
	ErrCodeReplicaNotFoundException = "ReplicaNotFoundException"

	// ErrCodeRequestLimitExceeded for service response error code
	// "RequestLimitExceeded".
	//
	// Throughput exceeds the current throughput quota for your account. Please
	// contact Amazon Web Services Support (https://aws.amazon.com/support) to request
	// a quota increase.
	ErrCodeRequestLimitExceeded = "RequestLimitExceeded"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The operation conflicts with the resource's availability. For example, you
	// attempted to recreate an existing table, or tried to delete a table currently
	// in the CREATING state.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The operation tried to access a nonexistent table or index. The resource
	// might not be specified correctly, or its status might not be ACTIVE.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTableAlreadyExistsException for service response error code
	// "TableAlreadyExistsException".
	//
	// A target table with the specified name already exists.
	ErrCodeTableAlreadyExistsException = "TableAlreadyExistsException"

	// ErrCodeTableInUseException for service response error code
	// "TableInUseException".
	//
	// A target table with the specified name is either being created or deleted.
	ErrCodeTableInUseException = "TableInUseException"

	// ErrCodeTableNotFoundException for service response error code
	// "TableNotFoundException".
	//
	// A source table with the name TableName does not currently exist within the
	// subscriber's account or the subscriber is operating in the wrong Amazon Web
	// Services Region.
	ErrCodeTableNotFoundException = "TableNotFoundException"

	// ErrCodeTransactionCanceledException for service response error code
	// "TransactionCanceledException".
	//
	// The entire transaction request was canceled.
	//
	// DynamoDB cancels a TransactWriteItems request under the following circumstances:
	//
	//    * A condition in one of the condition expressions is not met.
	//
	//    * A table in the TransactWriteItems request is in a different account
	//    or region.
	//
	//    * More than one action in the TransactWriteItems operation targets the
	//    same item.
	//
	//    * There is insufficient provisioned capacity for the transaction to be
	//    completed.
	//
	//    * An item size becomes too large (larger than 400 KB), or a local secondary
	//    index (LSI) becomes too large, or a similar validation error occurs because
	//    of changes made by the transaction.
	//
	//    * There is a user error, such as an invalid data format.
	//
	//    * There is an ongoing TransactWriteItems operation that conflicts with
	//    a concurrent TransactWriteItems request. In this case the TransactWriteItems
	//    operation fails with a TransactionCanceledException.
	//
	// DynamoDB cancels a TransactGetItems request under the following circumstances:
	//
	//    * There is an ongoing TransactGetItems operation that conflicts with a
	//    concurrent PutItem, UpdateItem, DeleteItem or TransactWriteItems request.
	//    In this case the TransactGetItems operation fails with a TransactionCanceledException.
	//
	//    * A table in the TransactGetItems request is in a different account or
	//    region.
	//
	//    * There is insufficient provisioned capacity for the transaction to be
	//    completed.
	//
	//    * There is a user error, such as an invalid data format.
	//
	// If using Java, DynamoDB lists the cancellation reasons on the CancellationReasons
	// property. This property is not set for other languages. Transaction cancellation
	// reasons are ordered in the order of requested items, if an item has no error
	// it will have None code and Null message.
	//
	// Cancellation reason codes and possible error messages:
	//
	//    * No Errors: Code: None Message: null
	//
	//    * Conditional Check Failed: Code: ConditionalCheckFailed Message: The
	//    conditional request failed.
	//
	//    * Item Collection Size Limit Exceeded: Code: ItemCollectionSizeLimitExceeded
	//    Message: Collection size exceeded.
	//
	//    * Transaction Conflict: Code: TransactionConflict Message: Transaction
	//    is ongoing for the item.
	//
	//    * Provisioned Throughput Exceeded: Code: ProvisionedThroughputExceeded
	//    Messages: The level of configured provisioned throughput for the table
	//    was exceeded. Consider increasing your provisioning level with the UpdateTable
	//    API. This Message is received when provisioned throughput is exceeded
	//    is on a provisioned DynamoDB table. The level of configured provisioned
	//    throughput for one or more global secondary indexes of the table was exceeded.
	//    Consider increasing your provisioning level for the under-provisioned
	//    global secondary indexes with the UpdateTable API. This message is returned
	//    when provisioned throughput is exceeded is on a provisioned GSI.
	//
	//    * Throttling Error: Code: ThrottlingError Messages: Throughput exceeds
	//    the current capacity of your table or index. DynamoDB is automatically
	//    scaling your table or index so please try again shortly. If exceptions
	//    persist, check if you have a hot key: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html.
	//    This message is returned when writes get throttled on an On-Demand table
	//    as DynamoDB is automatically scaling the table. Throughput exceeds the
	//    current capacity for one or more global secondary indexes. DynamoDB is
	//    automatically scaling your index so please try again shortly. This message
	//    is returned when writes get throttled on an On-Demand GSI as DynamoDB
	//    is automatically scaling the GSI.
	//
	//    * Validation Error: Code: ValidationError Messages: One or more parameter
	//    values were invalid. The update expression attempted to update the secondary
	//    index key beyond allowed size limits. The update expression attempted
	//    to update the secondary index key to unsupported type. An operand in the
	//    update expression has an incorrect data type. Item size to update has
	//    exceeded the maximum allowed size. Number overflow. Attempting to store
	//    a number with magnitude larger than supported range. Type mismatch for
	//    attribute to update. Nesting Levels have exceeded supported limits. The
	//    document path provided in the update expression is invalid for update.
	//    The provided expression refers to an attribute that does not exist in
	//    the item.
	ErrCodeTransactionCanceledException = "TransactionCanceledException"

	// ErrCodeTransactionConflictException for service response error code
	// "TransactionConflictException".
	//
	// Operation was rejected because there is an ongoing transaction for the item.
	ErrCodeTransactionConflictException = "TransactionConflictException"

	// ErrCodeTransactionInProgressException for service response error code
	// "TransactionInProgressException".
	//
	// The transaction with the given request token is already in progress.
	//
	// Recommended Settings
	//
	// This is a general recommendation for handling the TransactionInProgressException.
	// These settings help ensure that the client retries will trigger completion
	// of the ongoing TransactWriteItems request.
	//
	//    * Set clientExecutionTimeout to a value that allows at least one retry
	//    to be processed after 5 seconds have elapsed since the first attempt for
	//    the TransactWriteItems operation.
	//
	//    * Set socketTimeout to a value a little lower than the requestTimeout
	//    setting.
	//
	//    * requestTimeout should be set based on the time taken for the individual
	//    retries of a single HTTP request for your use case, but setting it to
	//    1 second or higher should work well to reduce chances of retries and TransactionInProgressException
	//    errors.
	//
	//    * Use exponential backoff when retrying and tune backoff if needed.
	//
	// Assuming default retry policy (https://github.com/aws/aws-sdk-java/blob/fd409dee8ae23fb8953e0bb4dbde65536a7e0514/aws-java-sdk-core/src/main/java/com/amazonaws/retry/PredefinedRetryPolicies.java#L97),
	// example timeout settings based on the guidelines above are as follows:
	//
	// Example timeline:
	//
	//    * 0-1000 first attempt
	//
	//    * 1000-1500 first sleep/delay (default retry policy uses 500 ms as base
	//    delay for 4xx errors)
	//
	//    * 1500-2500 second attempt
	//
	//    * 2500-3500 second sleep/delay (500 * 2, exponential backoff)
	//
	//    * 3500-4500 third attempt
	//
	//    * 4500-6500 third sleep/delay (500 * 2^2)
	//
	//    * 6500-7500 fourth attempt (this can trigger inline recovery since 5 seconds
	//    have elapsed since the first attempt reached TC)
	ErrCodeTransactionInProgressException = "TransactionInProgressException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BackupInUseException":                     newErrorBackupInUseException,
	"BackupNotFoundException":                  newErrorBackupNotFoundException,
	"ConditionalCheckFailedException":          newErrorConditionalCheckFailedException,
	"ContinuousBackupsUnavailableException":    newErrorContinuousBackupsUnavailableException,
	"DuplicateItemException":                   newErrorDuplicateItemException,
	"ExportConflictException":                  newErrorExportConflictException,
	"ExportNotFoundException":                  newErrorExportNotFoundException,
	"GlobalTableAlreadyExistsException":        newErrorGlobalTableAlreadyExistsException,
	"GlobalTableNotFoundException":             newErrorGlobalTableNotFoundException,
	"IdempotentParameterMismatchException":     newErrorIdempotentParameterMismatchException,
	"ImportConflictException":                  newErrorImportConflictException,
	"ImportNotFoundException":                  newErrorImportNotFoundException,
	"IndexNotFoundException":                   newErrorIndexNotFoundException,
	"InternalServerError":                      newErrorInternalServerError,
	"InvalidExportTimeException":               newErrorInvalidExportTimeException,
	"InvalidRestoreTimeException":              newErrorInvalidRestoreTimeException,
	"ItemCollectionSizeLimitExceededException": newErrorItemCollectionSizeLimitExceededException,
	"LimitExceededException":                   newErrorLimitExceededException,
	"PointInTimeRecoveryUnavailableException":  newErrorPointInTimeRecoveryUnavailableException,
	"ProvisionedThroughputExceededException":   newErrorProvisionedThroughputExceededException,
	"ReplicaAlreadyExistsException":            newErrorReplicaAlreadyExistsException,
	"ReplicaNotFoundException":                 newErrorReplicaNotFoundException,
	"RequestLimitExceeded":                     newErrorRequestLimitExceeded,
	"ResourceInUseException":                   newErrorResourceInUseException,
	"ResourceNotFoundException":                newErrorResourceNotFoundException,
	"TableAlreadyExistsException":              newErrorTableAlreadyExistsException,
	"TableInUseException":                      newErrorTableInUseException,
	"TableNotFoundException":                   newErrorTableNotFoundException,
	"TransactionCanceledException":             newErrorTransactionCanceledException,
	"TransactionConflictException":             newErrorTransactionConflictException,
	"TransactionInProgressException":           newErrorTransactionInProgressException,
}
