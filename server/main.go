package main

import (
	"fmt"
	"log"
	"net/http"
	"server/controller"
	db "server/db"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		s.SetContext("")
		s.Join("room1")
		fmt.Println("connected:", url.Query().Get("email"))
		return nil
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		url := s.URL()
		log.Printf("message from %v >> room1 : %v", url.Query().Get("email"), msg)

		if server.BroadcastToRoom(s.Namespace(), "room1", "chat message", msg) {
			chat := db.Chat{
				Source:      url.Query().Get("email"),
				Destination: "room1",
				Message:     msg,
			}
			controller.SaveChat(&chat)
		}
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()
	// db.Connet()
	http.Handle("/socket.io/", server)
	http.HandleFunc("/api/login", controller.Login)
	http.HandleFunc("/api/register", controller.Register)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
