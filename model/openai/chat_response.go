package openai

type ChatResponse struct {
	Choices []ResponseChoice `json:"choices"`
}

type ResponseChoice struct {
	FinishReason string          `json:"finish_reason"`
	Index        int             `json:"index"`
	Message      ResponseMessage `json:"message"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
