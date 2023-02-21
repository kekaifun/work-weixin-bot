package openai

type ChatRequest struct {
	Messages []RequestMessage `json:"messages"`
	Model    string           `json:"model"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
