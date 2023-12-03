package agt

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
