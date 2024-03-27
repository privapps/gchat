package main

import (
	// "internal/websocket"
	"log"
	"net/http"
	"os"

	"gopubsub/gwebsocket"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func handleDefault(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for:", r.URL.Path, r.Method)
	http.NotFound(w, r)
}
func main() {
	wsPath := getEnv("WS_PATH", "/socket")
	port := ":" + getEnv("PORT", "8000")

	http.HandleFunc(wsPath, gwebsocket.HandleWS)
	http.HandleFunc("/", handleDefault)
	log.Println("Server started on port: ", port, wsPath)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
