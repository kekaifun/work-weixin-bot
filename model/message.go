package model

type MsgContent struct {
	ToUsername   string `json:"ToUserName"`
	FromUsername string `json:"FromUserName"`
	CreateTime   uint32 `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Content      string `json:"Content"`
	MsgId        uint64 `json:"MsgId"`
	AgentId      uint32 `json:"AgentId"`
}
