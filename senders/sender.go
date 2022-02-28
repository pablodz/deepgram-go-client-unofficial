package senders

import (
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

func Send2Deepgram(data []byte, ws *websocket.Conn) (err error) {

	var i, chunks int
	t := time.NewTicker(20 * time.Millisecond)
	defer t.Stop()

	// send message each 20 miliseconds
	for range t.C {

		if i >= len(data) {
			return nil
		}

		var DeepGramChunkSize = 1000
		if i+DeepGramChunkSize > len(data) {
			DeepGramChunkSize = len(data) - i
		}
		// Send chunk
		err := ws.WriteMessage(websocket.BinaryMessage, data[i:i+DeepGramChunkSize])
		if err != nil {
			panic(err)
		}

		chunks++
		i += DeepGramChunkSize
	}

	return errors.New("ticker unexpectedly stopped")

}

func Send2DeepgramNoTicker(data []byte, ws *websocket.Conn) (err error) {

	var i, chunks int

	// send message each 20 miliseconds
	for {

		if i >= len(data) {
			return nil
		}

		var DeepGramChunkSize = 1000
		if i+DeepGramChunkSize > len(data) {
			DeepGramChunkSize = len(data) - i
		}
		// Send chunk
		err := ws.WriteMessage(websocket.BinaryMessage, data[i:i+DeepGramChunkSize])
		if err != nil {
			return err
		}

		chunks++
		i += DeepGramChunkSize
	}

}
