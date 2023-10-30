package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// AllRooms is the global hashmap for the server
var AllRooms RoomMap

type Resp struct {
	RoomID string `json:"room_id"`
}

// CreateRoomRequestHandler Create a Room and return roomID
func CreateRoomRequestHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()
	response := Resp{RoomID: roomID}

	// Use Gin's JSON function to send the JSON response
	c.JSON(200, response)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}

// JoinRoomRequestHandler will join the client in a particular room
func JoinRoomRequestHandler(c *gin.Context) {
	roomID := c.Query("roomID")
	if roomID == "" {
		c.String(400, "roomID missing in URL Parameters")
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
		return
	}

	AllRooms.InsertIntoRoom(roomID, false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error: ", err)
			break
		}
		msg.Client = ws
		msg.RoomID = roomID
		log.Println(msg.Message)
		broadcast <- msg
	}
}
