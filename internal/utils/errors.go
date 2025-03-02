package utils

import "errors"

var (
	ErrInvalidWord      = errors.New("word is invalid")
	ErrInvalidEmptyWord = errors.New("word is empty")
	ErrInvalidWordLen   = errors.New("word len is invalid")
)
