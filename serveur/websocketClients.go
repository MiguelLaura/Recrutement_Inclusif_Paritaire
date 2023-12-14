package serveur

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// WSClientsList is a map used to help manage a map of WSClients
type WSClientsList map[*WSClient]bool

// WSClient is a websocket WSClient, basically a frontend visitor
type WSClient struct {
	// the websocket connection
	connection *websocket.Conn

	// manager is the manager used to manage the WSClient
	manager *Manager

	idSimu string
}

// NewWSClient is used to initialize a new WSClient with all required values initialized
func NewWSClient(conn *websocket.Conn, manager *Manager) *WSClient {
	return &WSClient{
		connection: conn,
		manager:    manager,
	}
}

func (cli *WSClient) SetIdSimu(idSimu string) {
	cli.idSimu = idSimu
}

func (c *WSClient) readMessages() {
	defer func() {
		// Graceful Close the Connection once this
		// function is done
		c.manager.removeClient(c)
	}()
	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but not simple Disconnection
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		log.Println("MessageType: ", messageType)
		log.Println("Payload: ", payload)

		req := DefaultReq{}
		err = json.Unmarshal(payload, &req)
		if err != nil {
			log.Println("ERR:", err)
		}

		log.Println(req.T, req.D)

		rep := Resp{"coucou du serveur"}

		json, err := json.Marshal(rep)
		if err != nil {
			log.Println("ERR:", err)
		}

		c.connection.WriteMessage(websocket.TextMessage, json)
	}
}
