package controllers

import (
	"condo-control/models"
	"container/list"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// Subscription ...
type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

// Subscriber ...
type Subscriber struct {
	Name    string
	CondoID int
	Conn    *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 100)
	// Channel for exit users.
	unsubscribe = make(chan string, 100)
	// Send events here to publish them.
	publish     = make(chan models.Event, 100)
	waitingList = list.New()
	subscribers = list.New()
)

func newEvent(ep models.EventType, user string, condoID int, msg string) models.Event {
	v := models.Event{Type: ep, User: user, Timestamp: int(time.Now().Unix()), Content: msg, CondoID: condoID}
	return v
}

// Join ...
func Join(user string, condoID int, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws, CondoID: condoID}
}

// Leave ...
func Leave(user string) {
	unsubscribe <- user
}

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	BaseController
}

// Test ...
// @Title Test
// @Description test
// @router /test [get]
func (c *WebSocketController) Test() {

	publish <- newEvent(models.EventMessage, "cis", 2, "hola")

}

// Join ...
// @Title Join
// @Description Join
// @router /join [get]
func (c *WebSocketController) Join() {

	token := c.Ctx.Input.Query("token")
	token = "Bearer " + token //c.Ctx.Input.Header("Authorization")

	decodedToken, err := VerifyToken(token, "Supervisor")

	if err != nil {
		c.BadRequest(err)
		return
	}

	supervisorID, err := strconv.Atoi(decodedToken.UserID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	condoID, err := strconv.Atoi(decodedToken.CondoID)
	if err != nil {
		c.BadRequest(err)
		return
	}

	supervisor, err := models.GetSupervisorsByID(supervisorID)
	if err != nil {
		c.ServeErrorJSON(err)
		return
	}

	username := supervisor.Username

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	Join(username, condoID, ws)
	defer Leave(username)

	// Message receive loop.
	for {

		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EventMessage, username, condoID, string(p))
	}
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws == nil {
			continue
		}

		subCondoID := sub.Value.(Subscriber).CondoID

		if subCondoID != event.CondoID {
			continue
		}

		if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
			// User disconnected.
			unsubscribe <- sub.Value.(Subscriber).Name
		}
	}
}

// This function handles all incoming chan messages.
func channelsControl() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub) // Add user to the end of list.
			}
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			models.NewArchive(event)

		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					break
				}
			}
		}
	}
}

func init() {
	go channelsControl()
}
