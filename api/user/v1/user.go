package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"get" summary:"You first hello api"`
}

type UserRes struct {
	g.Meta     `mime:"application/json"`
	Avatar     string
	NickName   string
	CreateTime string
	Content    string
	Images     []string
	Location   string
}
