package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maruki00/Streaming_app/server"
)

func main() {

	router := gin.Default()
	router.POST("/rooms", server.CreateRoomRequestHandler)
	router.POST("/rooms", server.JoinRoomRequestHandler)
	fmt.Println(router)
}
