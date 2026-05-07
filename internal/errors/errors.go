package errors

import "errors"

var (
	ErrMissingDataBaseUrl = errors.New("require database_url missing")
)
