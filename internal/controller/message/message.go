package message

import (
	"life_notepad_api/api/message"
)

type Controller struct{}

func NewV1() message.IMessageV1 {
	return &Controller{}
}
