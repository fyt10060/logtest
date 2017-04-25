package model

//	"fmt"
import (
	"encoding/xml"
)

type MessageType string

const (
	MTText  MessageType = "text"
	MTImg               = "image"
	MTVoice             = "voice"
	MTVideo             = "video"
	MTLoc               = "location"
	MTLink              = "link"
	MTNull              = "null"
	MTNews              = "news"
	MTEvent             = "event"
)

type Message struct {
	ToUser     string      `xml:"ToUserName"`
	FromUser   string      `xml:"FromUserName"`
	CreateTime int64       `xml:"CreateTime"`
	MsgType    MessageType `xml:"MsgType"`
	MsgId      string      `xml:"MsgId"`
}

func (msg *Message) ReverseUserInMessage(newType MessageType) {
	tmp := msg.ToUser
	msg.ToUser = msg.FromUser
	msg.FromUser = tmp
	msg.MsgId = ""
	msg.MsgType = newType
}

type TextMessage struct {
	Message
	XMLName xml.Name `xml:"xml"`
	Content string   `xml:""`
}

type NewsMessage struct {
	Message
	XMLName      xml.Name `xml:"xml"`
	ArticleCount int
	Articles     ArticleList
}

type ArticleList struct {
	Articles []NewsDetail
	XMLName  xml.Name `xml:"Articles"`
}

type NewsDetail struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type result struct {
	Data interface{}
}

func GetMessageDetail(data []byte) interface{} {
	var msg Message
	err := xml.Unmarshal(data, &msg)
	if err != nil {
		result := ParseResult(ErrNoMsgFound, nil)
		return result
	}
	var content string
	switch msg.MsgType {
	case MTText:
		var txtMsg TextMessage
		err = xml.Unmarshal(data, &txtMsg)
		if err != nil {
			result := ParseResult(ErrNoMsgFound, nil)
			return result
		}
		return txtMsg
	case MTImg:
		content = "这是一条图片消息"
	case MTLink:
		content = "这是一条链接消息"
	case MTLoc:
		content = "这是一条位置消息"
	case MTVoice:
		content = "这是一条语音消息"
	case MTVideo:
		content = "这是一条视频消息"
	case MTEvent:
		content = "这是一条事件消息"
	default:
		content = "这是一条未知消息"
	}
	txtMsg := TextMessage{
		Message: msg,
		Content: content,
	}
	return txtMsg
}

func GetResponseMessage(response interface{}) []byte {
	//	r := result{
	//		Data: response,
	//	}
	rMsg, err := xml.Marshal(response)
	if err != nil {
		return nil
	}
	return rMsg
}
