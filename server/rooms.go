package server

import (
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.Mutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant, 0)
}

func (r *RoomMap) Get(roomId string) []Participant {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Map[roomId]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	roomId := uuid.NewString()
	r.Map[roomId] = make([]Participant, 0)
	return roomId
}

func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}

func (r *RoomMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn) {

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{Host: host, Conn: conn}
	r.Map[roomId] = append(r.Map[roomId], p)

}
