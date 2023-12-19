package serveur

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
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
				log.Printf("erreur : lecture du message : %v", err)
			}
			break
		}

		req := ActionReq{}
		err = json.Unmarshal(payload, &req)
		if err != nil {
			log.Println("erreur :", err)
		}

		//actions : start, pause, continue, stop
		if req.T == "action" {
			_, err = c.handleMessageFromWebSocket(req.IdSimu, req.D) //si type == action, on envoie l'action
			if err != nil {
				log.Println("erreur :", err)
			}
		}
	}
}

func (c *WSClient) handleMessageFromWebSocket(idSimulation string, message string) (resp string, err error) {
	simulation := c.manager.restServerAgent.simulations[idSimulation]
	if c.idSimu != idSimulation {
		c.idSimu = idSimulation
		simulation.AjouteWebSockerLogger(logger.NewSocketLogger(c.connection, 10))
	}

	if simulation == nil {
		log.Println("erreur : Simulation introuvable")
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
			log.Println("erreur : Action non reconnue")
			resp = ""
			err = errors.New("erreur : Action non reconnue")
			return
		}
	}

	return resp, nil
}
