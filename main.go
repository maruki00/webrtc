package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maruki00/Streaming_app/server"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	fmt.Println("Server Started On localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// router := gin.Default()
	// http.Rou
	// router.POST("/rooms", server.CreateRoomRequestHandler)
	// router.POST("/rooms", server.JoinRoomRequestHandler)
	// fmt.Println(router)
}
