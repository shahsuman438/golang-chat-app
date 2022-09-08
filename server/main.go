package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"server/controller"
	db "server/db"
)

func main() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		s.SetContext("")
		s.Join(url.Query().Get("room"))
		fmt.Println("connected:", url.Query().Get("email"))
		return nil
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		url := s.URL()
		log.Printf("message from %v >> %v : %v", url.Query().Get("email"), url.Query().Get("room"), msg)

		if server.BroadcastToRoom(s.Namespace(), url.Query().Get("room"), "chat message", msg) {
			chat := db.Chat{
				Source:      url.Query().Get("email"),
				Destination: url.Query().Get("room"),
				Message:     msg,
			}
			if msg != "" {
				controller.SaveChat(&chat)
			}
			return
		}
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()
	db.Connet()
	http.Handle("/socket.io/", server)
	http.HandleFunc("/api/login", controller.Login)
	http.HandleFunc("/api/register", controller.Register)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
