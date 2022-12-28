package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"shop-v2/internal/cmd"
	_ "shop-v2/internal/logic"
	_ "shop-v2/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
