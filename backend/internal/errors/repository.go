package errors

import "errors"

// CustomerRepo
var (
	ErrQueryNotFound = errors.New("query not found")
	ErrInsertFailed  = errors.New("insert failed")
)
