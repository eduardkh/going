package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

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
		return
	}
	log.Println("Client connected")
	var response WsJsonResponse
	response.Message = "Welcome to the chat!"
	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
		return
	}
}
