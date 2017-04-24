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
)

type Message struct {
	ToUser     string      `xml:"ToUserName"`
	FromUser   string      `xml:"FromUserName"`
	CreateTime int64       `xml:"CreateTime"`
	MsgType    MessageType `xml:"MsgType"`
	MsgId      string      `xml:"MsgId"`
}

type TextMessage struct {
	Message
	Content string   `xml:""`
	XMLName xml.Name `xml:"xml"`
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
	switch msg.MsgType {
	case MTText:
		var txtMsg TextMessage
		err = xml.Unmarshal(data, &txtMsg)
		if err != nil {
			result := ParseResult(ErrNoMsgFound, nil)
			return result
		}
		return txtMsg
	default:
		break
	}
	return nil
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
