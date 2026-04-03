package entity

import "errors"

var (
	ENOTFOUND      = errors.New("not found")
	EUNAUTHORIZED  = errors.New("unauthorized")
	ENOTAVAILABLE  = errors.New("not available")
	EINTERNAL      = errors.New("internal error")
	EINVALID       = errors.New("invalid")
	ECONFLICT      = errors.New("conflict")
)
