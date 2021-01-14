package main

import (
	_ "write_code_in_/boot"
	_ "write_code_in_/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
