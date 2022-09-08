package controller

import (
	"bufio"
	"client/interfaces"
	"log"
	"os"
	"github.com/zhouhui8915/go-socket.io-client"
)

func Chat(user interfaces.Register) {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["email"] = user.Email
	opts.Query["userName"] = user.UserName
	uri := "http://localhost:8000/"
	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
		return
	})
	client.On("chat message", func(msg string) {
		log.Printf("message:>>  %v\n", msg)
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("chat message", command)
	}
}
