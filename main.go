package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/websocket"
	"github.com/pablodz/deepgram-go-client-unofficial/config"
	"github.com/pablodz/deepgram-go-client-unofficial/models"
	"github.com/pablodz/deepgram-go-client-unofficial/senders"
)

func SendInBackground(audioRaw []byte, clientWss *websocket.Conn) {
	err := senders.Send2Deepgram(audioRaw, clientWss)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create websocket client
	clientWss := config.ConfigSTTDeepgram("YOUR_APIKEY_HERE", "8000", "es", "1")
	// Close websocket when finish
	defer clientWss.Close()

	// Get audio data from file PCM 16 bits, 8000 Hz, mono
	audioRaw, err := ioutil.ReadFile("audio.raw") // the file is inside the local directory
	if err != nil {
		panic(err)
	}

	// Start sending audio data in background
	go SendInBackground(audioRaw, clientWss)

	// Receive messages from Deepgram
	for {
		_, message, err := clientWss.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		var resp models.Response
		if err := json.Unmarshal(message, &resp); err != nil {
			panic(err)
		}

		if resp.IsFinal {
			fmt.Println("Message: ", resp.Channel.Alternatives[0].Transcript, "Confidence", resp.Channel.Alternatives[0].Confidence)
		}
	}
}
