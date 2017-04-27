// weixin
package controllers

import (
	//	"fmt"
	"io/ioutil"
	"log"

	"github.com/astaxie/beego"
	"github.com/weixinhost/yar.go"
	"github.com/weixinhost/yar.go/server"
)

type WeixinController struct {
	beego.Controller
}

type YarClass struct{}

func (c *YarClass) Echo() string {
	log.Println("echo handler")
	return "string"
}

func (this *WeixinController) Post() {
	c := this.Ctx
	r := c.Request
	w := c.ResponseWriter

	body, _ := ioutil.ReadAll(r.Body)

	s := server.NewServer(&YarClass{})

	s.Opt.LogLevel = yar.LogLevelDebug | yar.LoglevelNormal | yar.LogLevelError

	s.Register("echo", "Echo")

	_ = s.Handle(body, w)

	w.Write([]byte("11313"))
}
