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
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Closed from server, dont print close 1011 (internal server error): NET-0001
				break
			}
			fmt.Println("Error reading websocket message response:", err)
			break
		}

		var resp models.Response
		if err := json.Unmarshal(message, &resp); err != nil {
			fmt.Println("Error parsing json response:", err)
		}

		if resp.IsFinal {
			// Only get final transcript
			fmt.Println("Confidence", resp.Channel.Alternatives[0].Confidence, "\tMessage: ", resp.Channel.Alternatives[0].Transcript)
		}
	}
}
