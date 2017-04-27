// weixin
package controllers

import (
	"fmt"
	"io/ioutil"
	"log"

	"logtest/model"

	"github.com/astaxie/beego"
	"github.com/mitchellh/mapstructure"
	"github.com/weixinhost/yar.go"
	"github.com/weixinhost/yar.go/server"
)

type WeixinController struct {
	beego.Controller
}

type YarClass struct{}

type Weixin struct{}

func (c *YarClass) Echo() string {
	log.Println("echo handler")
	return "echo teshdfa"
}

func (c *Weixin) Weixin(accountParam, messageParam, otherParam map[string]interface{}) map[string]interface{} {

	var reply = model.WeixinReply{}
	reply.Action = "reply"
	reply.Exit = false
	reply.Log = false
	var account model.Account
	err := mapstructure.Decode(accountParam, &account)
	if err != nil {
		fmt.Println(err)
		return model.Struct2Map(reply)
	}
	fmt.Printf("appid: %s, account name: %s", account.Appid, account.Name)

	var message model.WeiXinMessageInfo
	err = mapstructure.Decode(messageParam, &message)
	if err != nil {
		fmt.Println(err)
		return model.Struct2Map(reply)
	}
	fmt.Printf("message type: %s, message content:%s", message.MsgType, message.Content)

	var other model.HostWxOther
	err = mapstructure.Decode(otherParam, &other)
	if err != nil {
		fmt.Println(err)
		return model.Struct2Map(reply)
	}

	if message.MsgType == "text" {
		if message.Content == "a2" {
			reply.Exit = true
			reply.Log = true
			var nMap map[string]interface{}
			nMap = make(map[string]interface{})
			nMap["message"] = model.GetReplyWithSendMsg(message, "a2 + 付雨桐")
			nMap["type"] = "raw"
		}
	}

	return model.Struct2Map(reply)
}

func (this *WeixinController) Post() {
	c := this.Ctx
	r := c.Request
	w := c.ResponseWriter

	body, _ := ioutil.ReadAll(r.Body)

	//	s := server.NewServer(&YarClass{})

	//	s.Opt.LogLevel = yar.LogLevelDebug | yar.LoglevelNormal | yar.LogLevelError

	//	s.Register("echo", "Echo")

	//	_ = s.Handle(body, w)

	s1 := server.NewServer(&Weixin{})
	s1.Opt.LogLevel = yar.LogLevelDebug | yar.LoglevelNormal | yar.LogLevelError
	s1.Register("weixin", "Weixin")

	_ = s1.Handle(body, w)

	w.Write([]byte("11313"))
}
