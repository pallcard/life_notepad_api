package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"life_notepad_api/internal/controller/file"
	"life_notepad_api/internal/controller/note"
	"life_notepad_api/internal/controller/user"

	"life_notepad_api/internal/controller/hello"
)

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"*"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
					note.NewV1(),
					file.NewV1(),
				)
			})
			s.BindHandler("/ws", func(r *ghttp.Request) {
				ws, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}
				for {
					msgType, msg, err := ws.ReadMessage()
					if err != nil {
						return
					}
					fmt.Println(msg)
					if err = ws.WriteMessage(msgType, msg); err != nil {
						return
					}
				}
			})
			s.Run()
			return nil
		},
	}
)
