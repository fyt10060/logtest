// router
package controllers

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"logtest/model"
	"logtest/service"

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
		var content string
		content = doOperation(msg.Content)
		rMsg := model.TextMessage{
			Message: model.Message{
				ToUser:     msg.FromUser,
				FromUser:   msg.ToUser,
				MsgType:    model.MTText,
				CreateTime: msg.CreateTime + 1,
			},
			Content: content,
		}
		response := model.GetResponseMessage(rMsg)
		w.Write(response)
	case []byte:
		response, _ := result.([]byte)
		w.Write(response)
	}

}

func doOperation(content string) (respContent string) {
	switch content {
	case "算术", "算数", "做算术", "做算数":
		service.SetDoingOperation(true)
		return "开始计算，请输入第一个数字"
	case "结束", "=":
		service.SetDoingOperation(false)
		service.SetShouldBeNumber(true)
		return "已结束，你的结果是：呵呵哒"
	default:
		if service.CheckDoingOperation() {
			if service.CheckShouldBeNum() {
				num, err := strconv.ParseFloat(content, 10)
				if err != nil {
					fmt.Println(num)
					return "请输入数字"
				}
				service.SetShouldBeNumber(false)
				return "请输入符号"
			} else {
				if content == "+" || content == "-" || content == "*" || content == "/" {
					service.SetShouldBeNumber(true)
					return "请输入数字"
				} else {
					return "请输入一个符号，目前仅支持'+''-''*''/'"
				}
			}
		} else {
			return content
		}

	}
}
