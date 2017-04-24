// router
package controllers

import (
	"fmt"
	"io/ioutil"

	"logtest/model"

	"github.com/astaxie/beego"
)

type RouterController struct {
	beego.Controller
}

//func init() {
//	RouterController.EnableRender = false
//}

func (this *RouterController) Post() {
	c := this.Ctx
	w := c.ResponseWriter
	r := c.Request
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		result := model.ParseResult(model.ErrNoMsgFound, nil)
		w.Write(result)
		return
	}
	result := model.GetMessageDetail(data)
	switch result.(type) {
	case nil:
		result := model.ParseResult(model.ErrNoMsgFound, nil)
		w.Write(result)
		return
	case model.TextMessage:
		msg, _ := result.(model.TextMessage)
		fmt.Println(msg.MsgId)
		rMsg := model.TextMessage{
			Message: model.Message{
				ToUser:     msg.FromUser,
				FromUser:   msg.ToUser,
				MsgType:    model.MTText,
				CreateTime: msg.CreateTime + 10,
			},
			Content: msg.Content + "heheda",
		}
		response := model.GetResponseMessage(rMsg)
		w.Write(response)
	case []byte:
		response, _ := result.([]byte)
		w.Write(response)
	}

}
