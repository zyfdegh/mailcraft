package api

import (
	"log"

	"github.com/kataras/iris"
	"github.com/zyfdegh/mailcraft/entity"
	"github.com/zyfdegh/mailcraft/service"
)

// HandlePostSender handle POST /sender
func HandlePostSender(ctx *iris.Context) {
	resp := entity.Resp{}

	sender := &entity.Sender{}
	err := ctx.ReadJSON(sender)
	if err != nil {
		log.Printf("read json error: %v\n", err)
		resp.ErrNo = iris.StatusBadRequest
		resp.Errmsg = "parse body error"
		resp.Success = false
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	path, err := service.SaveSender(sender)
	if err != nil {
		log.Printf("save sender to file error: %v\n", err)
		resp.ErrNo = iris.StatusInternalServerError
		resp.Errmsg = "save sender to file error, " + err.Error()
		resp.Success = false
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	log.Printf("sender saved to %s\n", path)
	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}
