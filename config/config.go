package config

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func ConfigSTTDeepgram(apikey string, sampleRate string, language string, channels string) (ws *websocket.Conn) {

	u := url.URL{Scheme: "wss", Host: "api.deepgram.com", Path: "/v1/listen"}
	wsUrl := u.String() + "?"
	if sampleRate != "" {
		wsUrl += "sample_rate=" + sampleRate + "&"
	}
	if language != "" {
		wsUrl += "language=" + language + "&"
	}
	if channels != "" {
		wsUrl += "channels=" + channels + "&"
	}
	// dial with header
	header := http.Header{}
	header.Add("Authorization", "Token "+apikey)

	ws, _, err := websocket.DefaultDialer.Dial(wsUrl, header)
	if err != nil {
		panic(err)
	}

	fmt.Println("Client created at ", wsUrl)
	return ws
}
