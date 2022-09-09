package controller

import (
	"bufio"
	"client/models"
	"fmt"
	"log"
	"os"

	"github.com/zhouhui8915/go-socket.io-client"
)

func Chat(user models.Register, room string) {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["email"] = user.Email
	opts.Query["userName"] = user.UserName
	opts.Query["room"] = room
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
	fmt.Println("##### START YOUR CONVERSATION HERE ######")
	fmt.Println("           Room:- " + room)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("chat message", command)
	}
}
