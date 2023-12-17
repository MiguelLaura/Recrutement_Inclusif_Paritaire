package serveur

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	/**
	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
	*/
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	sync.Mutex
	clients         WSClientsList
	restServerAgent *RestServerAgent
}

func NewManager(rsa *RestServerAgent) *Manager {
	return &Manager{sync.Mutex{}, make(WSClientsList, 0), rsa}
}

func (rs *RestServerAgent) setupWebsocket(mux *http.ServeMux) {
	manager := NewManager(rs)
	rs.websocketManager = manager
	mux.HandleFunc("/ws", manager.serveWS)
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	log.Println("New connection")
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	client := NewWSClient(conn, m)

	m.addClient(client)
	go client.readMessages()
}

func (m *Manager) addClient(client *WSClient) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) removeClient(client *WSClient) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)

		// on ne veut pas que la simulation se termine le client est arrêté
		// exemple : reload de la page = la simulation peut continuer
	}
}
