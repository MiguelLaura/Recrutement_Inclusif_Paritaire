package agt

import "sync"

// ------------ SIMULATION ------------

type Action int

// Action est une enumeration
const (
	NOOP Action = iota
	LIBRE
	AGRESSION
	FIN
)

// Permet la communication entre agents
type Communicateur struct {
	Act     Action
	Payload any
}

// ------------ EMPLOYE ------------

type EmployeID string

var agtCnt int = 0

type Employe struct {
	id           EmployeID
	genre        Genre
	anciennete   int //entre 0 et 40
	santeMentale int //entre 0 et 100
	agresseur    bool
	comportement Comportement
	competence   int //entre 0 et 10
	entreprise   *Entreprise
	chnl         chan Communicateur
}

type Genre int

const (
	Homme Genre = iota
	Femme
)

type Comportement float64

// Probabilité de porter plainte pour les Employés
const (
	Plainte100 Comportement = 1.0
	Plainte75  Comportement = 0.75
	Plainte50  Comportement = 0.5
	Plainte25  Comportement = 0.25
	Plainte0   Comportement = 0.0
)

type Entreprise struct {
	sync.Mutex
	employes      []Employe
	departs       []Employe
	plaintes      [][]Employe
	nbDepressions int
	nbRenvois     int
	recrutement   Recrutement
	ca            float64
	nbActions     int
	nbAgresseurs  int
	chnl          chan Communicateur
	chnlActions   chan Communicateur
}

// ------------ RECRUTEMENT ------------

type StratParite int

const (
	StratVide StratParite = iota // décrit l'absence de StartParite
	PrioHomme
	PrioFemme
	Hasard
)

type TypeRecrutement int

const (
	Vide TypeRecrutement = iota // décrit l'absence de TypeRecrutement
	Competences
	PlacesReservees
)

type Recrutement struct {
	entreprise             *Entreprise
	objectif               float64     // -1 si non renseigné, entre 0 et 1 sinon
	stratAvant             StratParite // stratVide si non renseigné
	stratApres             StratParite
	typeRecrutementAvant   TypeRecrutement // Vide si non renseigné
	typeRecrutementApres   TypeRecrutement
	pourcentagePlacesAvant float64 // -1 si non renseigné, entre 0 et 1 sinon
	pourcentagePlacesApres float64
}
