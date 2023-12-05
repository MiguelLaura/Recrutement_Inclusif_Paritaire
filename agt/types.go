package agt

import (
	"math"
)

// ------------ SIMULATION ------------

type Action int

// Action est une enumeration
const (
	NOOP Action = iota
	LIBRE
	AGRESSION
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

type Comportement float32

// Probabilité de porter plainte pour les Employés
const (
	Plainte100 Comportement = 1.0
	Plainte75  Comportement = 0.75
	Plainte50  Comportement = 0.5
	Plainte25  Comportement = 0.25
	Plainte0   Comportement = 0.0
)

// ------------ ENTREPRISE ------------

type Entreprise struct {
	employes    []Employe
	departs     []Employe
	plaintes    [][]Employe
	recrutement Recrutement
	ca          float64
}

// La fonction NewEntreprise doit créer l'entreprise et générer les employés de façon à respecter le quota de parité initial

func NewEntreprise(nbEmployesInit int, pariteInit float32, recrut Recrutement) *Entreprise {
	e := new(Entreprise)

	var nbFemmes int = int(math.Round(float64(nbEmployesInit) * float64(pariteInit)))
	var nbHommes int = nbEmployesInit - nbFemmes
	var employesInit []Employe

	for i := 0; i < nbFemmes; i++ {
		employesInit = append(employesInit, *GenererEmployeInit(&e, Femme))
	}
	for i := 0; i < nbHommes; i++ {
		employesInit = append(employesInit, *GenererEmployeInit(&e, Homme))
	}
	e.employes = employesInit
	e.departs = make([]Employe, 0)
	e.plaintes = make([][]Employe, 0)
	e.recrutement = recrut
	e.ca = 0.0
	return e
}

func (e Entreprise) Employes() []Employe {
	return e.employes
}

func (e Entreprise) Departs() []Employe {
	return e.departs
}

func (e Entreprise) Plaintes() [][]Employe {
	return e.plaintes
}

func (e Entreprise) Recrutement() Recrutement {
	return e.recrutement
}

func (e Entreprise) Ca() float64 {
	return e.ca
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
