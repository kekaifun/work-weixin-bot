package work_weixin

type MsgContent struct {
	WebhookUrl     string `json:"WebhookUrl" xml:"WebhookUrl"`
	From           From   `json:"From"`
	MsgType        string `json:"MsgType" xml:"MsgType"`
	MsgId          string `json:"MsgId" xml:"MsgId"`
	ChatId         string `json:"ChatId" xml:"ChatId"`
	PostId         string `json:"PostId" xml:"PostId"`
	ChatType       string `json:"ChatType" xml:"ChatType"`
	GetChatInfoUrl string `json:"GetChatInfoUrl" xml:"GetChatInfoUrl"`
	Text           Text   `json:"Text"`
}

type From struct {
	UserId string `json:"UserId" xml:"UserId"`
	Name   string `json:"Name" xml:"Name"`
	Alias  string `json:"Alias" xml:"Alias"`
}

type Text struct {
	Content string `json:"Content" xml:"Content"`
}
