package types

import (
	"message/types/message"
)

type MGetResult struct {
	Messages []*message.Message
	Errors   []MGetError
}

type MGetError struct {
	MessageID string
	Error     error
}
