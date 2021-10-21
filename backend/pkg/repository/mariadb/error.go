package mariadb

import "errors"

var (
	ErrQueryNotFound       = errors.New("query not found")
	ErrInsertFailed        = errors.New("insert failed")
	ErrInternalServerError = errors.New("internal server error")
)
