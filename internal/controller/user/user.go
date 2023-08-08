package user

import (
	"life_notepad_api/api/user"
)

type Controller struct{}

func NewV1() user.IUserV1 {
	return &Controller{}
}
