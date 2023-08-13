package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NoteListReq note列表
type NoteListReq struct {
	g.Meta   `path:"/noteList" tags:"Note" method:"post" summary:"获取笔记列表"`
	PageNum  int
	PageSize int
}

type NoteListRes struct {
	g.Meta   `mime:"application/json"`
	NoteList []NoteItem
	Total    int
}

type NoteItem struct {
	Avatar     string
	NickName   string
	CreateTime string
	Content    string
	Images     []string
	Location   string
}

type AddNoteReq struct {
	g.Meta   `path:"/addNote" tags:"Note" method:"post" summary:"增加笔记"`
	UserId   int
	Content  string
	Images   []string
	Location string
}

type AddNoteRes struct {
	g.Meta `mime:"application/json"`
	Id     int
}
