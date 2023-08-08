package main

import (
	_ "life_notepad_api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"life_notepad_api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
