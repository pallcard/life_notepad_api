package note

import (
	"life_notepad_api/api/note"
)

type Controller struct{}

func NewV1() note.INoteV1 {
	return &Controller{}
}
