package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "life_notepad_api/api/user/v1"
	"life_notepad_api/internal/common"
)

func (c *Controller) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	userRes := v1.UserRes{
		Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9",
		NickName:   "用爱发电的小程序开发者",
		CreateTime: "2022-11-11 11:11:11",
		Content:    "对研究经济学的学者来讲，2003年以来中国经济出现了许多不容易理解的现象。",
		Images: []string{"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9",
			"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9"},
	}
	g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
		Code: 0,
		Data: userRes,
	})
	return
}

func (c *Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	loginRes := v1.LoginRes{}
	if req.Email == "xx@qq.com" && req.Password == "123456" {
		loginRes.UserId = 1
		loginRes.Avatar = "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fsafe-img.xhscdn.com%2Fbw1%2F1c5a5c88-3063-4615-905a-a9b9e4c2acb5%3FimageView2%2F2%2Fw%2F1080%2Fformat%2Fjpg&refer=http%3A%2F%2Fsafe-img.xhscdn.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1694020103&t=15637d7ccac5a81aa1e0fa4a558efed9"
		loginRes.NickName = "开发者"
		loginRes.Description = "用爱发电的小程序开发者"
		loginRes.CreateTime = "2022-11-11 11:11:11"
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code: 0,
			Data: loginRes,
		})
	} else {
		g.RequestFromCtx(ctx).Response.WriteJson(common.Res{
			Code:    1,
			Message: "用户名或密码错误",
		})
	}
	return
}
