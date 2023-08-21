package chat

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"life_notepad_api/internal/consts"
	"life_notepad_api/internal/dao"
	"life_notepad_api/internal/model"
	"life_notepad_api/internal/model/entity"
	"strings"
)

func (c *Controller) WebSocket(r *ghttp.Request) {
	msg := &model.ChatMsg{}

	// 初始化WebSocket请求
	var (
		ws  *ghttp.WebSocket
		err error
	)
	ws, err = r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), err)
		return
	}

	userId := r.GetQuery("userId").String()
	if userId == "" {
		userId = r.Request.RemoteAddr
	}

	// 初始化时设置用户昵称为当前链接信息
	userIds.Add(userId)
	userWsMap.Set(userId, ws)

	// 初始化后向所有客户端发送上线消息
	//c.writeUserListToClient()

	for {
		// 阻塞读取WS数据
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			// 如果失败，那么表示断开，这里清除用户信息
			// 为简化演示，这里不实现失败重连机制
			userIds.Remove(userId)
			userWsMap.Remove(userId)
			// 通知所有客户端当前用户已下线
			//c.writeUserListToClient()
			break
		}
		// JSON参数解析
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			c.write(ws, model.ChatMsg{
				Type:       "error",
				Data:       "消息格式不正确: " + err.Error(),
				SenderId:   userId,
				ReceiverId: "",
			})
			continue
		}

		// WS操作类型
		switch msg.Type {
		//单发消息
		case consts.ChatTypeSingle:
			// 发送间隔检查
			intervalKey := fmt.Sprintf("%p", ws)
			if ok, _ := cache.SetIfNotExist(r.Context(), intervalKey, struct{}{}, consts.SendInterval); !ok {
				c.write(ws, model.ChatMsg{
					Type:       consts.ChatTypeError,
					Data:       "您的消息发送得过于频繁，请休息下再重试",
					SenderId:   userId,
					ReceiverId: "",
				})
				continue
			}
			// 有消息时，单发消息
			if msg.Data != nil {
				if err = c.writeToUser(r.Context(), msg.ReceiverId,
					model.ChatMsg{
						Type:       consts.ChatTypeSingle,
						Data:       gconv.String(msg.Data),
						SenderId:   userId,
						ReceiverId: msg.ReceiverId,
					}); err != nil {
					g.Log().Error(r.Context(), err)
				}
			}
		// 发送消息
		case consts.ChatTypeGroup:
			// 发送间隔检查
			intervalKey := fmt.Sprintf("%p", ws)
			if ok, _ := cache.SetIfNotExist(r.Context(), intervalKey, struct{}{}, consts.SendInterval); !ok {
				c.write(ws, model.ChatMsg{
					Type:       consts.ChatTypeError,
					Data:       "您的消息发送得过于频繁，请休息下再重试",
					SenderId:   userId,
					ReceiverId: "",
				})
				continue
			}
			// 有消息时，群发消息
			if msg.Data != nil {
				if err = c.writeGroup(
					model.ChatMsg{
						Type:       consts.ChatTypeGroup,
						Data:       ghtml.SpecialChars(gconv.String(msg.Data)),
						SenderId:   userId,
						ReceiverId: msg.ReceiverId,
					}); err != nil {
					g.Log().Error(r.Context(), err)
				}
			}
		}
	}
}

// PushChatToUsers http向指定用户发消息
func (c *Controller) PushChatToUsers(r *ghttp.Request) {
	msg := &model.ChatMsg{}
	msg.Data = r.Get("data").String()
	msg.SenderId = r.Get("SenderId").String()
	msg.ReceiverId = r.Get("ReceiverId").String()
	fmt.Print(msg)
	if msg.ReceiverId != "" && userWsMap.Get(msg.ReceiverId) != nil {
		// 有消息时，群发消息
		if msg.Data != nil {
			if err := c.writeToUser(r.Context(), msg.ReceiverId,
				model.ChatMsg{
					Type:       consts.ChatTypeSingle,
					Data:       gconv.String(msg.Data),
					SenderId:   msg.SenderId,
					ReceiverId: msg.ReceiverId,
				}); err != nil {
				g.Log().Error(r.Context(), err)
			}
		}
	} else {
		r.Response.Write("{abc:cdb}")
	}

}

// PushChat http群发消息
func (c *Controller) PushChat(r *ghttp.Request) {
	fmt.Print(c)
	fmt.Print(r)
	var (
		err error
	)
	msg := &model.ChatMsg{}
	msg.Data = r.Get("name").String()
	msg.SenderId = r.Get("from").String()
	fmt.Print(msg)
	// 有消息时，群发消息
	if msg.Data != nil {
		if err = c.writeGroup(
			model.ChatMsg{
				Type:     "send",
				Data:     gconv.String(msg.Data),
				SenderId: msg.SenderId,
			}); err != nil {
			g.Log().Error(r.Context(), err)
		}
	}
}

// 向客户端写入消息。
// 内部方法不会自动注册到路由中。
func (c *Controller) write(ws *ghttp.WebSocket, msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return ws.WriteMessage(websocket.BinaryMessage, msgBytes)
}

// 向指定用户名客户端群发消息。
// 内部方法不会自动注册到路由中。
func (c *Controller) writeToUser(ctx context.Context, user string, msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	// send
	userWs := userWsMap.Get(user)
	if userWs != nil {
		err = userWs.(*ghttp.WebSocket).WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			return err
		}
	}

	// db todo trans
	msgData := gconv.String(msg.Data)
	link := 2
	if strings.HasPrefix(msgData, "[link]") {
		msgData = strings.TrimLeft(msgData, "[link]")
		link = 1
	}
	_, err = dao.Message.Ctx(ctx).InsertIgnore(entity.Message{
		SenderId:   gconv.Int(msg.SenderId),
		ReceiverId: gconv.Int(msg.ReceiverId),
		Content:    msgData,
		Link:       link,
		Unread:     1,
	})
	if err != nil {
		return err
	}

	result, err := dao.Chat.Ctx(ctx).Save(entity.Chat{
		ReceiverId: gconv.Int(msg.ReceiverId),
		SenderId:   gconv.Int(msg.SenderId),
		Content:    msgData,
		Link:       link,
		Unread:     1,
	})
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 1 { // 插入,首次插入
		_, err := dao.Chat.Ctx(ctx).Save(entity.Chat{
			ReceiverId: gconv.Int(msg.SenderId),
			SenderId:   gconv.Int(msg.ReceiverId),
			Content:    "",
			Link:       2,
			Unread:     2,
		})
		if err != nil {
			return err
		}

	}

	return nil
}

// 向所有客户端群发消息。
// 内部方法不会自动注册到路由中。
func (c *Controller) writeGroup(msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	//users.RLockFunc(func(m map[interface{}]interface{}) {
	userWsMap.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			fmt.Print(user)
			err := m[user].(*ghttp.WebSocket).WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				return
			}
		}
	})
	return nil
}

// 向客户端返回用户列表。
// 内部方法不会自动注册到路由中。
func (c *Controller) writeUserListToClient() error {
	array := garray.NewSortedStrArray()
	userWsMap.Iterator(func(k interface{}, v interface{}) bool {
		array.Add(gconv.String(k))
		return true
	})
	if err := c.writeGroup(model.ChatMsg{
		Type:       "list",
		Data:       array.Slice(),
		SenderId:   "",
		ReceiverId: "",
	}); err != nil {
		return err
	}
	return nil
}
