package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var wsChannel = make(chan WsJsonRequest)
var clients = make(map[WebSocketConnection]string)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func renderPage(w http.ResponseWriter, template string, data jet.VarMap) error {
	view, err := views.GetTemplate(template)
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

type WebSocketConnection struct {
	*websocket.Conn
}
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsJsonRequest struct {
	Username string               `json:"username"`
	Message  string               `json:"message"`
	Action   string               `json:"action"`
	Conn     *WebSocketConnection `json:"-"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected")
	var response WsJsonResponse
	response.Message = "Welcome to the chat!"

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)
}

func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ListenForWs Error: ", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsJsonRequest

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = conn

			wsChannel <- payload
		}
	}
}

func ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-wsChannel
		response.Action = "Got Here!"
		response.Message = fmt.Sprintf("Message Here! , Action was: %v", e.Action)
		WsBroadcastToAll(response)
	}

}

func WsBroadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("WsBroadcastToAll Error: ", err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
