package chatgpt

import (
	"fmt"
	urlpkg "net/url"
	"testing"
)

func TestURL(t *testing.T) {
	url := "https://api.openai.com/v1/chat/completions"
	u, err := urlpkg.Parse(url)
	if err != nil {
		t.Fatalf("invalid url: %s", url)
	}
	fmt.Println(u.Host)
	fmt.Println(u.Path)
}
