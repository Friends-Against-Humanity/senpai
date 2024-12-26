package services

import "errors"

var (
	ErrInternal      = errors.New("an internal error occurred")
	ErrQuotaExceeded = errors.New("you have exceeded your requests quota")
)
