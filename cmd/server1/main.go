package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	currentTime := time.Now().String()
	mux := http.NewServeMux()
	mux.Handle("/health", NewHealthHandler(currentTime))
	mux.Handle("/messages", NewMessageHandler())

	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              ":8080",
		Handler:           mux,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

type HealthHandler struct {
	serveTime string
}

type MessageHandler struct {
	Messages []Message
	Counter  int
}

type Message struct {
	Content string `json:"content"`
}

func NewHealthHandler(serveTime string) *HealthHandler {
	return &HealthHandler{
		serveTime: serveTime,
	}
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{Counter: 1}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type in the HTTP header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Return a JSON response `{"status": "ok"}`
	w.Write([]byte(`{"status": "ok", "startedAt": "` + h.serveTime + `"}`))
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type in the HTTP header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Return a JSON response `{"status": "ok"}`
	if r.Method == "GET" {
		msg, _ := json.Marshal(m.Messages)
		w.Write(msg)
	} else {
		currMsg := Message{Content: fmt.Sprintf("message %d", m.Counter)}
		m.Messages = append(m.Messages, currMsg)
		currMsgResp, _ := json.Marshal(currMsg)
		w.Write(currMsgResp)
		m.Counter++
	}
}
