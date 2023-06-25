package model

import "errors"

var (
	ErrMissingArgument    = errors.New("missing argument")
	ErrInvalidMessageType = errors.New("invalid message-type")
	ErrNotFound           = errors.New("not found")
)

// Error defines the ... metadata
type Error struct {
	ErrorMessage string `json:"errorMessage"`
}
