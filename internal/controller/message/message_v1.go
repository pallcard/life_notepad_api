package message

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "life_notepad_api/api/message/v1"
	"life_notepad_api/internal/common"
	"life_notepad_api/internal/common/cos"
	"life_notepad_api/internal/dao"
	"life_notepad_api/internal/model/entity"
	"strings"
	"time"
)

func (c *Controller) ChatList(ctx context.Context,
	req *v1.ChatListReq) (res *v1.ChatListRes, err error) {
	// count
	cond := g.Map{
		dao.Chat.Columns().ReceiverId: req.UserId,
	}
	total, err := dao.Chat.Ctx(ctx).Where(cond).Count()
	if err != nil {
		return nil, err
	}

	list := make([]v1.ChatItem, 0)
	res = &v1.ChatListRes{
		ChatList: list,
		Total:    total,
	}
	if total <= 0 {
		return
	}

	// list
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	chatList := make([]entity.Chat, 0, total)
	err = dao.Chat.Ctx(ctx).Page(req.PageNum, req.PageSize).
		Where(cond).Order("id desc").Scan(&chatList)
	if err != nil {
		return
	}

	// get user map
	sendIds := make([]int, 0)
	for _, noteItem := range chatList {
		sendIds = append(sendIds, noteItem.SenderId)
	}
	userList := make([]entity.User, 0, total)
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Id: sendIds,
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
	for _, chatItem := range chatList {
		avatar := userIdMap[chatItem.SenderId].Avatar
		if !strings.HasPrefix(avatar, "http") {
			avatar, err = cos.Cli.GetPresignedURL(ctx, avatar, 365*24*time.Hour)
			if err != nil {
				continue
			}
		}
		list = append(list, v1.ChatItem{
			Id:           chatItem.Id,
			SenderId:     chatItem.SenderId,
			SenderAvatar: avatar,
			NickName:     userIdMap[chatItem.SenderId].NickName,
			Content:      chatItem.Content,
			Unread:       chatItem.Unread,
			Link:         chatItem.Link,
			CreateTime:   chatItem.CreatedAt.Local().Format("Y-m-d H:i:s"),
		})
	}
	res.ChatList = list

	return
}

func (c *Controller) MessageList(ctx context.Context,
	req *v1.MessageListReq) (res *v1.MessageListRes, err error) {

	// count
	cond := g.Map{
		dao.Message.Columns().SenderId:   []int{req.SenderId, req.ReceiverId},
		dao.Message.Columns().ReceiverId: []int{req.SenderId, req.ReceiverId},
	}
	total, err := dao.Message.Ctx(ctx).Where(cond).Count()
	if err != nil {
		return nil, err
	}

	list := make([]v1.MessageItem, 0)
	res = &v1.MessageListRes{
		MessageList: list,
		Total:       total,
	}
	if total <= 0 {
		return
	}

	// list
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	messageList := make([]entity.Message, 0, total)
	err = dao.Message.Ctx(ctx).Page(req.PageNum, req.PageSize).
		Where(cond).Order("id desc").Scan(&messageList)
	if err != nil {
		return
	}

	// gen res
	for _, messageItem := range messageList {
		list = append(list, v1.MessageItem{
			Id:         messageItem.Id,
			SenderId:   messageItem.SenderId,
			ReceiverId: messageItem.ReceiverId,
			Content:    messageItem.Content,
			Unread:     messageItem.Unread,
			Link:       messageItem.Link,
			CreateTime: messageItem.CreatedAt.Local().Format("Y-m-d H:i:s"),
		})
	}
	res.MessageList = list

	dao.Chat.Ctx(ctx).Where(g.Map{
		dao.Message.Columns().ReceiverId: req.SenderId,
	}).Update(g.Map{
		dao.Message.Columns().Unread: 0,
	})
	return
}
