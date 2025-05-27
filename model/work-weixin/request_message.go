package work_weixin

type MsgContent struct {
	WebhookUrl     string `json:"webhook_url" xml:"WebhookUrl"`
	From           From   `json:"from" xml:"From"`
	MsgType        string `json:"msgtype" xml:"MsgType"`
	MsgId          string `json:"msgid" xml:"MsgId"`
	ChatId         string `json:"chatid" xml:"ChatId"`
	PostId         string `json:"postid" xml:"PostId"`
	ChatType       string `json:"chattype" xml:"ChatType"`
	GetChatInfoUrl string `json:"get_chat_info_url" xml:"GetChatInfoUrl"`
	Text           Text   `json:"text" xml:"Text"`
}

type From struct {
	UserId string `json:"userid" xml:"UserId"`
	Name   string `json:"name" xml:"Name"`
	Alias  string `json:"alias" xml:"Alias"`
}

type Text struct {
	Content string `json:"content" xml:"Content"`
}
