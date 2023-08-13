package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"life_notepad_api/internal/cmd"
	_ "life_notepad_api/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
