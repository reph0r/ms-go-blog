package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/context"
	"net/http"
)

func (*HTMLApi) LoginNew(ctx *context.MsContext)  {
	login := common.Template.Login

	login.WriteData(ctx.W,config.Cfg.Viewer)
}
func (*HTMLApi) Login(w http.ResponseWriter,r *http.Request)  {
	login := common.Template.Login

	login.WriteData(w,config.Cfg.Viewer)
}
