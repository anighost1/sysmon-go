package sse

import (
	"fmt"
	"net/http"
)

type Client struct {
	ch chan string
}

var clients = make(map[*Client]bool)

func AddClient(c *Client) {
	clients[c] = true
}

func RemoveClient(c *Client) {
	delete(clients, c)
	close(c.ch)
}

func Broadcast(message string) {
	for client := range clients {
		select {
		case client.ch <- message:
		default:
			// skip if blocked
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	client := &Client{
		ch: make(chan string),
	}

	AddClient(client)
	defer RemoveClient(client)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	notify := r.Context().Done()

	for {
		select {
		case msg := <-client.ch:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()

		case <-notify:
			return
		}
	}
}
