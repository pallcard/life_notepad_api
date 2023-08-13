package note

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	v1 "life_notepad_api/api/note/v1"
	"life_notepad_api/internal/common"
)

func (c *Controller) NoteList(ctx context.Context,
	req *v1.NoteListReq) (res *v1.NoteListRes, err error) {
	fmt.Println(req)
	list := make([]v1.NoteItem, 0)

	total := 14
	// 1 5
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	for i := 0; i < req.PageSize; i++ {
		index := req.PageSize*(req.PageNum-1) + i
		if index >= total {
			break
		}
		list = append(list, v1.NoteItem{
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9",
			NickName:   "用爱发电的小程序开发者",
			CreateTime: "2022-11-11 11:11:11",
			Content:    fmt.Sprintf("内容%d", index),
			Images: []string{"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9",
				"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9"},
			Location: "湖北武汉",
		})
	}

	noteListRes := v1.NoteListRes{
		NoteList: list,
		Total:    total,
	}
	g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
		Code: 0,
		Data: noteListRes,
	})
	return
}

func (c *Controller) AddNote(ctx context.Context, req *v1.AddNoteReq) (res *v1.AddNoteRes, err error) {
	fmt.Println(req)
	if len(req.Content) <= 0 {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: "文本内容不能为空",
		})
		return
	}

	addNoteRes := v1.AddNoteRes{
		Id: 1,
	}
	g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
		Code: 0,
		Data: addNoteRes,
	})
	return
}
