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

type BroadcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	client  *websocket.Conn
}

var broadcast = make(chan BroadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomId] {
			if client.Conn != msg.client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
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
		log.Fatal("web socket error ", err.Error())
	}

	AllRooms.InsertIntoRoom(string(roomId[0]), false, ws)

	go broadcaster()
	for {
		var msg BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("read error : ", err.Error())
		}
		msg.client = ws
		msg.RoomId = roomId[0]
		broadcast <- msg
	}

}
