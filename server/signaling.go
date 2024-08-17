package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

type resp struct {
	RoomId string `json: "room_id"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId := AllRooms.CreateRoom()

	json.NewEncoder(w).Encode(resp{RoomId: roomId})

}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {

	roomId, ok := r.URL.Query()["roomId"]
	if !ok {
		log.Fatal("room id is missing")
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("web socket error")
	}

	AllRooms.InsertIntoRoom(roomId, false, ws)

}
