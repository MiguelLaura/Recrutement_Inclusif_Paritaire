package serveur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"
	"time"

	"gitlab.utc.fr/mennynat/ia04-project/agt"
)

// -----------------------------------
//       Fonctions utils serveur
// -----------------------------------

type RestServerAgent struct {
	sync.Mutex
	id         string
	addr       string
	simulation agt.Simulation
}

// retourne un pointeur sur un nouveau ServeurAgent
func NewRestServerAgent(addr string) *RestServerAgent {
	return &RestServerAgent{id: addr, addr: addr}
}

// retourne un bool si la méthode est bien celle demandée (POST, GET, ..;)
func (rsa *RestServerAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

// -----------------------------
//       Gestion requêtes
// -----------------------------

// retourne la structure en go depuis la requête JSON
func (*RestServerAgent) decodeRequestNewSimulation(r *http.Request) (req requestNewSimulation, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

// teste la requête de création, créé la simulation et la lance
func (rsa *RestServerAgent) createNewSimulation(w http.ResponseWriter, r *http.Request) {
	//set-up header
	setupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestNewSimulation(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		fmt.Println(err.Error())
		fmt.Println("erreur décodage requête /new_simulation")
		return
	}

	// traitement de la requête
	var resp responseNewSimulation

	//TO DO : test sur l'id - simulation déjà créée ??

	if req.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := "il manque un identifiant"
		w.Write([]byte(msg))
		return
	} else if req.NbEmployes <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "le nombre d'employés doit être > 0"
		w.Write([]byte(msg))
		return
	} else if req.PourcentageFemmes < 0.0 || req.PourcentageFemmes > 1.0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "le pourcentage de femmes doit être entre 0 et 1"
		w.Write([]byte(msg))
		return
	} else if req.Objectif < 0.0 || req.Objectif > 1.0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "l'objectif doit être entre 0 et 1"
		w.Write([]byte(msg))
		return
	} else if req.PourcentagePlacesAvant < 0.0 || req.PourcentagePlacesAvant > 1.0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "le pourcentage de places avant doit être entre 0 et 1"
		w.Write([]byte(msg))
		return
	} else if req.PourcentagePlacesApres < 0.0 || req.PourcentagePlacesApres > 1.0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "le pourcentage de places après doit être entre 0 et 1"
		w.Write([]byte(msg))
		return
	} else {
		resp.Id_simulation = req.ID
		w.WriteHeader(http.StatusCreated)
		serial, _ := json.Marshal(resp)
		w.Write(serial)
		fmt.Println("OK création et lancement simulation")

		//*********** CREATION & LANCEMENT SIMULATION ****************************
		s := agt.NewSimulation(req.NbEmployes, req.PourcentageFemmes, -1, 600*time.Second)
		rsa.simulation = *s //la simulation (non un pointeur)
		s.Run()
	}
}

// ---------------------
//       Server
// ---------------------

// Ajoute les options CORS pour pouvoir envoyer les informations par formulaire sur page HTML
func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Retourne la page index
func home(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("./serveur/templates/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

// Lance le serveur
func (rsa *RestServerAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", home) //index
	mux.HandleFunc("/new_simulation", rsa.createNewSimulation)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
