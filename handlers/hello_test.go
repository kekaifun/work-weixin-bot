package handlers

//
//func TestHello(t *testing.T) {
//	body := `<xml><Encrypt><![CDATA[nwyeG0zMTIqJHG1asJmoHNUiGABBJcVeZ6Mo7wUp/8hIYSttwY6VNeY6JGQNKjl0ZcVLqYu26hta71qyGvuWquMX1Y10pAJeoVhIKFw4dqNCJKBewuGWmVODMBwimzCseUeCczNTTf3GLtSZHJyseFc2mgZmfFwwNFbJ+QtDxNZpjmEIamMTmSfedO6RzanziIydb+LJkHa04oi7/2GtzxrLLKOpfVUHqqQK9OSwrk/F5OjPvyUMxAb6eW9V8Pr/HDXfbanOXBlxpdSLnwGynwoBIAA9zU3pp3wGakDyVNpGGlc5ZRyP+XFcBN7ug7mIDl+Ci+oxBP4lMuRwftE1kvwvjTGHFycrW6xsi+Cyo/KOxn7sjZWhsnY++mIQk/dSSqtZhbNKIIokpnGVaT14GyMH52/VI4sWo1maGJzfJJEeYeIy/2NJAuHPNrV9iK9836H4ClK34dGOlvf2HUVz6gGFFPfhQoAbZ9qU09XY1qKEI8RFQsXQO6bS11YTPcR4ClyMzEgRo03pmYOWjRxE9hMscgJ2DpQtJhMxJ1rZxKHY4ZagTE8wDO6vbT+ZMvjO9xFaYo/Ivw4gNukuEB/mEegb7UfDGOCSo1jscp7e4WCiC49d6t+RUYGe5yIUC+2Eb0DPl48rRsX4WSKE6f22hep8yrPj9qXh5m6cabXu+iyhX4uYBitFMElmZQ7+KLNPlbyw5+uQsApBgEEFQUw5EZiWDwcTwmkw5pLMxiikYDVXkTRiDp1PcnPW3SFvK1q/SnhPxze86NK7YH8q+41WhQTDTFN0V3v9giSRaBG/j8WDjW5jh9JnjeQBEAK5gAHflxvHbHGwTnUmior6U0YdGZ/jxxRHJGi+WT7MGguZAR/OCD4zKEYn0z7Itik0PBQC14I0mIwA/LVDtb8yGH5Mbwv55IFGQWgyWvDXrmCgjio=]]></Encrypt></xml>`
//	wxcpt := wxcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxcrypt.XmlType)
//
//	msg, err := wxcpt.DecryptMsg("13bb530f7016ac49ca29ba800feb1723139a8a63", "1676961998", "1655500c4c7305bf", []byte(body))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(msg))
//
//	var msgContent model.MsgContent
//	nerr := xml.Unmarshal(msg, &msgContent)
//	if nerr != nil {
//		fmt.Println(nerr)
//		return
//	}
//	fmt.Println(msgContent)
//}
