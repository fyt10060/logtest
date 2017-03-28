// solveWechat
package solvewechatmsg

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type WechatMsg struct {
	typestring string
	msgid      string
	content    string
}

func (m *WechatMsg) PraseJson() string {
	return fmt.Sprintf("{\"type\": %s, \"msgid\": %s, \"content\": %s}", m.typestring, m.msgid, m.content)
}

func SolveMsg(w http.ResponseWriter, r *http.Request) {

}
