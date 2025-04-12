package domain

import "errors"

var (
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)