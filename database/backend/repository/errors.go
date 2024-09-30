package repository

import "errors"

var (
	ErrBlockNotFound                = errors.New("block is not found")
	ErrEventNotFound                = errors.New("event is not found")
	ErrEventHasAlreadyBeenProcessed = errors.New("event has already been processed")
)
