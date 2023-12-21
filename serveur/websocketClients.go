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
			err = c.handleMessageFromWebSocket(req.IdSimu, req.D) //si type == action, on envoie l'action
			if err != nil {
				log.Println("erreur :", err)
			}
		}
	}
}

func (c *WSClient) handleMessageFromWebSocket(idSimulation string, message string) (err error) {
	simulation := c.manager.restServerAgent.simulations[idSimulation]
	myLogger := logger.NewSocketLogger(c.connection, 10)
	if c.idSimu != idSimulation {
		c.idSimu = idSimulation
		simulation.AjouteWebSockerLogger(myLogger)
	}

	if simulation == nil {
		log.Println("erreur : Simulation introuvable")
		myLogger.Err("Simulation introuvable")
		err = errors.New("erreur : Simulation introuvable")
		return
	} else {
		switch message {
		case "start":
			simulation.Start()
		case "pause":
			simulation.Pause()
		case "continue":
			simulation.Continue()
		case "stop":
			simulation.End()
		case "relancer":
			simulation.Relancer()
		default:
			log.Println("erreur : Action non reconnue")
			err = errors.New("erreur : Action non reconnue")
			return
		}
	}

	return
}
