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
	id               string
	addr             string
	simulations      map[string]*agt.Simulation
	websocketManager *Manager
}

// retourne un pointeur sur un nouveau ServeurAgent
func NewRestServerAgent(addr string) *RestServerAgent {
	return &RestServerAgent{id: addr, addr: addr}
}

// retourne un bool si la méthode est bien celle demandée (POST, GET, ..;)
func (rsa *RestServerAgent) verifierMethode(methode string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != methode {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "méthode %q pas autorisée", r.Method)
		return false
	}
	return true
}

// -----------------------------
//       Gestion requêtes
// -----------------------------

// retourne la structure en go depuis la requête JSON
func (*RestServerAgent) decoderRequeteNouvelleSimulation(r *http.Request) (req requeteNouvelleSimulation, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

// teste la requête de création, créé la simulation et la lance
func (rsa *RestServerAgent) creerNouvelleSimulation(w http.ResponseWriter, r *http.Request) {
	//set-up header
	ajouterCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.verifierMethode("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decoderRequeteNouvelleSimulation(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		log.Println(err.Error())
		log.Println("erreur : décodage requête /new_simulation")
		return
	}

	// traitement de la requête
	var resp reponseNouvelleSimulation

	if req.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque un identifiant"
		w.Write([]byte(msg))
		return
	} else if req.TypeRecrutementAvant == 0 { //impossible de ne pas avoir un type de recrutement avant vide
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.TypeRecrutementApres == 0 && req.Objectif != -1 { //cas où objectif mais pas de type de recrut
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.TypeRecrutementAvant == 1 && req.StratAvant == 0 { //cas où compétences avant choisi, mais pas de prio
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.TypeRecrutementApres == 1 && req.StratApres == 0 { //cas où compétences après choisi, mais pas de prio
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if (req.TypeRecrutementAvant == 2 || req.TypeRecrutementAvant == 3) && req.PourcentagePlacesAvant == -1 { //cas où places réservées avant choisi, mais pas de pourcentage
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if (req.TypeRecrutementApres == 2 || req.TypeRecrutementApres == 3) && req.PourcentagePlacesApres == -1 { //cas où places réservées après choisi, mais pas de pourcentage
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : il manque des informations sur le recrutement"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.NbEmployes <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : le nombre d'employés doit être > 0"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.NbAnnees <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : le nombre d'années doit être > 0"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.PourcentageFemmes < 0.0 || req.PourcentageFemmes > 1.0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : le pourcentage de femmes doit être entre 0 et 1"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.Objectif != -1 && !(req.Objectif >= 0.0 && req.Objectif <= 1.0) {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(req.Objectif)
		msg := "erreur : l'objectif doit être entre 0 et 1"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.PourcentagePlacesAvant != -1 && !(req.PourcentagePlacesAvant >= 0.0 && req.PourcentagePlacesAvant <= 1.0) {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : le pourcentage de places avant doit être entre 0 et 1"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else if req.PourcentagePlacesApres != -1 && !(req.PourcentagePlacesApres >= 0.0 && req.PourcentagePlacesApres <= 1.0) {
		w.WriteHeader(http.StatusBadRequest)
		msg := "erreur : le pourcentage de places après doit être entre 0 et 1"
		log.Println(msg)
		w.Write([]byte(msg))
		return
	} else {
		_, ok := rsa.simulations[req.ID]
		if ok {
			w.WriteHeader(http.StatusBadRequest)
			msg := "erreur : cette simulation existe déjà"
			log.Println(msg)
			w.Write([]byte(msg))
			return
		} else {
			resp.ID = req.ID
			w.WriteHeader(http.StatusCreated)
			serial, _ := json.Marshal(resp)
			w.Write(serial)
			log.Printf("\nOK création simulation %s\n", resp.ID)
			s := agt.NewSimulation(req.NbEmployes, req.PourcentageFemmes, req.Objectif, req.StratAvant, req.StratApres, req.TypeRecrutementAvant, req.TypeRecrutementApres, req.PourcentagePlacesAvant, req.PourcentagePlacesApres, req.NbAnnees)
			rsa.simulations[resp.ID] = s //pointeur sur la simulation
		}
	}
}

// ---------------------
//       Server
// ---------------------

// Ajoute les options CORS pour pouvoir envoyer les informations par formulaire sur page HTML
func ajouterCORS(w *http.ResponseWriter, req *http.Request) {
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

// TODO : check valeur id simulation quand demande 'visualisation.html'
// TODO : problème récupération page visualisationEntreprise quand l'url finit par '/'

// Lance le serveur
func (rsa *RestServerAgent) Start() {

	visualisationPath := "/viz_simulation/"

	// création du multiplexer
	mux := http.NewServeMux()

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./serveur/static/")))
	mux.Handle("/static/", fileHandler)

	mux.HandleFunc("/", home) //index

	mux.Handle(visualisationPath, http.StripPrefix(visualisationPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := path.Dir("./serveur/templates/visualisation.html")

		// Recupère l'ID de l'entreprise à visualiser
		// v := r.URL.Path

		// set header
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, p+"/visualisation.html")
	})))

	mux.HandleFunc("/new_simulation", rsa.creerNouvelleSimulation)

	rsa.setupWebsocket(mux)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	rsa.simulations = make(map[string]*agt.Simulation)

	// lancement du serveur
	fmt.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
