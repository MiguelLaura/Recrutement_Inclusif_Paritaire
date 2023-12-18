package serveur

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WSClientsList map[*WSClient]bool

type WSClient struct {
	connection *websocket.Conn
	manager    *Manager
	idSimu     string
}

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
		c.manager.removeClient(c)
	}()
	// Loop Forever
	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		req := ActionReq{}
		err = json.Unmarshal(payload, &req)
		if err != nil {
			log.Println("ERR:", err)
		}

		//actions : start, pause, continue, stop
		resp := ""
		if req.T == "action" {
			resp, err = c.handleMessageFromWebSocket(req.IdSimu, req.D) //si type == action, on envoie l'action
			if err != nil {
				log.Printf("ERREUR ACTION")
			}
		}

		json, err := json.Marshal(resp)
		if err != nil {
			log.Println("ERR:", err)
		}
		c.connection.WriteMessage(websocket.TextMessage, json)
	}
}

func (c *WSClient) handleMessageFromWebSocket(idSimulation string, message string) (resp string, err error) {
	simulation := c.manager.restServerAgent.simulations[idSimulation]
	if simulation == nil {
		fmt.Println("Simulation introuvable.")
		resp = "Simulation introuvable"
		err = errors.New("erreur : Simulation introuvable")
		return
	} else {
		switch message {
		case "start":
			resp = simulation.Start()
		case "pause":
			resp = simulation.Pause()
		case "continue":
			resp = simulation.Continue()
		case "stop":
			resp = simulation.End()
		default:
			fmt.Println("Err : Action non reconnue.")
			resp = ""
			err = errors.New("erreur : Action non reconnue")
			return
		}
	}
	return resp, nil
}
