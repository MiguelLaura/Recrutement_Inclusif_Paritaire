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

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	sync.Mutex
	clients WSClientsList
}

// NewManager is used to initalize all the values inside the manager
func NewManager() *Manager {
	return &Manager{sync.Mutex{}, make(WSClientsList, 0)}
}

func setupWebsocket(mux *http.ServeMux) {
	manager := NewManager()
	mux.HandleFunc("/ws", manager.serveWS)
}

// serveWS is a HTTP Handler that the has the Manager that allows connections
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

// addClient will add clients to our clientList
func (m *Manager) addClient(client *WSClient) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.clients[client] = true
}

// removeClient will remove the client and clean up
func (m *Manager) removeClient(client *WSClient) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		// close connection
		client.connection.Close()
		// remove
		delete(m.clients, client)

		// TODO : terminer la simulation associée
	}
}

func (m *Manager) initClient() {}
