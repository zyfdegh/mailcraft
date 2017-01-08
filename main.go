package main

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/mailcraft/api"
)

func main() {
	ir := iris.New()
	v1 := ir.Party("/v1")
	v1.Post("/sender", api.HandlePostSender)

	ir.Listen(":8080")
}
