package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMessageHandler(t *testing.T) {

	bot := &Bot{}
	router := gin.Default()
	router.POST("", bot.MessageHandler)

	body := strings.NewReader(`<xml><Encrypt><![CDATA[rCKLpm+EoceXu6zb9gBblYx3puZOmybDRhzisunsVjZzts66SA5baQIGcsMKZW4+od6r8rUMzKAwx8v30Dez34SqUAGzRwJLk/PK+OwDjY3DUypTfPQWuKOPt40FtTgbBSAxu1ip+7NDtRDE+JzcaOwIJosjDkZhpR20yGbL5rhNB7k0WuWRrc2S8HuctAidv580UDWjom6G8HnCC4s8mcQqOg+wyXEXMz1CMcmonwQXVnX07o3s4EhHzZlyjL9oCiO+vW9XRlkzj0IoXSOWBNBsN6vxhWQXD/wYMA8FrVMCk+PU/I+YKRNNpee28w5IVg7/W6cvNV3FtrhcNotXkmAiNHq1+qVQNLvJ8MbdE3n4C8N+mPsAOxtu0zVX9qV7R+GjaAfNXEQ2ylsBnITUIRaYqn68XtOtG9ZY5fxEXohnu5j/OBpKHCW7MkkOnn+aXWT8ndcY+arpIXUdgdMx4foIcUS+nO2GLGkIwUqf4F55nTry2wo6GhInDihnd/2LBpNRYSbPnY2gcE6qk2eBDS5cOszUNxttECCp6mwAMfVwU/sJ/2iSKdnKnXXOVMkJwE+yrbRESqJ8YIrd247VxqHb3XdK9DwslPYpvHg3WogFdJkxh5iAEV+G+zo1N9I3bEcUt/yRzRnMJxvP2+fldrGHChZlStMcycEUq0km/iJxW7cUbEjw+VH7ghLbFqciDdUrMuFqXK2ImpkVh5JTpK8OGn3QicmFFKkiuiL4lYzLsaoL9dubSpTadZJcaA5YbMUozKmNJDuS8ODCRoSktT9BTSC9lBUwh8x0zR/qF8LTqjCRFBNFRmR8F2s3fgVtSr093wbef9L0I8G4XO4H9Z0eI60ezf5X6xpGLkEdXQNq5VnApqrLGE0dvXgXvSpyulEswOSzbq9kuCX7kN0tIS7UdA5TIpW4zPBI5z8nwU0=]]></Encrypt></xml>`)

	req, _ := http.NewRequest(http.MethodPost, "/?msg_signature=c3a02a3479ae5d297ddc16f16747bb11d544d49b&timestamp=1676968052&nonce=f32cf1946af7b1a3", body)

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, req)
	fmt.Println(recoder.Body.String())
}
