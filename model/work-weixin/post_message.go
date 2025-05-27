package work_weixin

type PostMessage struct {
	ChatId  string   `json:"chatid"`
	MsgType string   `json:"msgtype"`
	Text    TextJson `json:"text"`
}

type TextJson struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list,omitempty"`
}
