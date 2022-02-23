# deepgram-go-client-unofficial

Unofficial client for the [Deepgram](https://deepgram.com) API

## Steps to use

Record an audio PCM 16 bit, rate 8k, mono

```bash
pw-cat --record audio.raw --channels 1 --rate 8000 
```

Download the Golang package

```bash
go get -v github.com/pablodz/deepgram-go-client-unofficial
```

Create a client

```golang
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
```

Run the program

```bash
go run main.go
```

Get responses
```bash
Client created at  wss://api.deepgram.com/v1/listen?sample_rate=8000&language=es&channels=1&
Message:  hola Confidence 0.8359375
Message:  esto es una prueba Confidence 0.99902344
Message:  uno Confidence 0.9980469
Message:  dos Confidence 0.9970703
Message:  tres Confidence 0.9921875
Message:  cuatro Confidence 0.9951172
```