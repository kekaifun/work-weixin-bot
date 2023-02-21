package work_weixin

import (
	"encoding/xml"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

type MentionedList struct {
	xml.Name `xml:"MentionedList"`
	Item     []wxbizmsgcrypt.CDATA `xml:"Item"`
}

type MentionedMobileList struct {
	xml.Name `xml:"MentionedMobileList"`
	Item     []wxbizmsgcrypt.CDATA `xml:"Item"`
}

type ResponseText struct {
	Content             wxbizmsgcrypt.CDATA `json:"Content" xml:"Content"`
	MentionedList       MentionedList       `json:"MentionedList,omitempty" xml:"MentionedList"`
	MentionedMobileList MentionedMobileList `json:"MentionedMobileList,omitempty" xml:"MentionedMobileList"`
}

type ReplyMsgContent struct {
	XMLName       xml.Name     `xml:"xml"`
	MsgType       string       `xml:"MsgType"`
	VisibleToUser string       `xml:"VisibleToUser,omitempty"`
	Text          ResponseText `json:"Text" xml:"Text"`
}
