package note

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "life_notepad_api/api/note/v1"
	"life_notepad_api/internal/common"
	"life_notepad_api/internal/dao"
	"life_notepad_api/internal/model/entity"
	"strings"
)

func (c *Controller) NoteList(ctx context.Context,
	req *v1.NoteListReq) (res *v1.NoteListRes, err error) {
	list := make([]v1.NoteItem, 0)

	// count
	total, err := dao.Note.Ctx(ctx).Count()
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	noteListRes := v1.NoteListRes{
		NoteList: list,
		Total:    total,
	}
	if total <= 0 {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code: 0,
			Data: noteListRes,
		})
	}

	// list
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	noteList := make([]entity.Note, 0, total)
	err = dao.Note.Ctx(ctx).Page(req.PageNum, req.PageSize).Order("id desc").Scan(&noteList)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	// get user map
	userIds := make([]int, 0)
	for _, noteItem := range noteList {
		userIds = append(userIds, noteItem.UserId)
	}
	userList := make([]entity.User, 0, total)
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Id: userIds,
	}).Scan(&userList)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: err.Error(),
		})
		return
	}
	userIdMap := map[int]entity.User{}
	for _, user := range userList {
		userIdMap[user.Id] = user
	}

	// gen res
	for _, noteItem := range noteList {
		images := strings.Split(noteItem.Images, ",")
		list = append(list, v1.NoteItem{
			Id:         noteItem.Id,
			UserId:     noteItem.UserId,
			Avatar:     userIdMap[noteItem.UserId].Avatar,
			NickName:   userIdMap[noteItem.UserId].NickName,
			Content:    noteItem.Content,
			Images:     images,
			Location:   noteItem.Location,
			CreateTime: noteItem.CreatedAt.Format("Y-m-d H:i:s"),
		})
	}
	noteListRes.NoteList = list
	g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
		Code: 0,
		Data: noteListRes,
	})
	return
}

func (c *Controller) AddNote(ctx context.Context, req *v1.AddNoteReq) (res *v1.AddNoteRes, err error) {
	if len(req.Content) <= 0 {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: "文本内容不能为空",
		})
		return
	}

	result, err := dao.Note.Ctx(ctx).InsertIgnore(entity.Note{
		UserId:   req.UserId,
		Content:  req.Content,
		Images:   strings.Join(req.Images, ","),
		Location: req.Location,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: err.Error(),
		})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	addNoteRes := v1.AddNoteRes{
		Id: int(id),
	}
	g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
		Code: 0,
		Data: addNoteRes,
	})
	return
}
