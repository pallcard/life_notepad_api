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

type LoginReq struct {
	g.Meta   `path:"/login" tags:"User" method:"post" summary:"You first hello api"`
	Email    string `v:"required|length:6,30#请输入邮箱|邮箱错误"`
	PassWord string `v:"required|length:6,30#请输入密码|密码长度不够"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	UserId int
}
