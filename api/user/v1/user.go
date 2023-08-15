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
}

type LoginReq struct {
	g.Meta   `path:"/login" tags:"User" method:"post" summary:"You first hello api"`
	Email    string `v:"required|length:6,30#请输入邮箱|邮箱错误"`
	Password string `v:"required|length:6,30#请输入密码|密码长度不够"`
}

type LoginRes struct {
	g.Meta      `mime:"application/json"`
	UserId      int
	Avatar      string
	NickName    string
	Description string
	CreateTime  string
}

type UpdateUserReq struct {
	g.Meta      `path:"/updateUser" tags:"User" method:"post" summary:"You first hello api"`
	UserId      int `v:"required#UserId不能为空"`
	Avatar      string
	NickName    string
	Description string
}

type UpdateUserRes struct {
	g.Meta      `mime:"application/json"`
	UserId      int
	Avatar      string
	NickName    string
	Description string
	CreateTime  string
}
