package model

import "encoding/xml"

type MsgContent struct {
	WebhookUrl     string `json:"WebhookUrl" xml:"WebhookUrl,cdata"`
	From           From   `json:"From"`
	MsgType        string `json:"MsgType" xml:"MsgType,cdata"`
	MsgId          string `json:"MsgId" xml:"MsgId,cdata"`
	ChatId         string `json:"ChatId" xml:"ChatId,cdata"`
	PostId         string `json:"PostId" xml:"PostId,cdata"`
	ChatType       string `json:"ChatType" xml:"ChatType,cdata"`
	GetChatInfoUrl string `json:"GetChatInfoUrl" xml:"GetChatInfoUrl,cdata"`
	Text           Text   `json:"Text"`
}

type From struct {
	UserId string `json:"UserId" xml:"UserId,cdata"`
	Name   string `json:"Name" xml:"Name,cdata"`
	Alias  string `json:"Alias" xml:"Alias,cdata"`
}

type Text struct {
	Content             string              `json:"Content" xml:"Content,cdata"`
	MentionedList       MentionedList       `json:"MentionedList,omitempty"`
	MentionedMobileList MentionedMobileList `json:"MentionedMobileList,omitempty"`
}

type Item struct {
	XMLName xml.Name `xml:"Item"`
	Value   string   `xml:",cdata"`
}

type MentionedList []Item
type MentionedMobileList []Item

type ReplyMsgContent struct {
	XMLName       xml.Name `xml:"xml"`
	MsgType       string   `json:"MsgType"`
	VisibleToUser string   `json:"VisibleToUser,omitempty" xml:"VisibleToUser,omitempty"`
	Text          Text     `json:"Text"`
}

type MessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}
