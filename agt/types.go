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
	PrioHomme StratParite = iota
	PrioFemme
	Hasard
)

type TypeRecrutement int

const (
	Competences TypeRecrutement = iota
	PlacesReservees
)

type Recrutement struct {
	objectif               float32
	stratAvant             StratParite
	stratApres             StratParite
	typeRecrutementAvant   TypeRecrutement
	typeRecrutementApres   TypeRecrutement
	pourcentagePlacesAvant float32
	pourcentagePlacesApres float32
}

func NewRecrutement(obj float32, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float32, ppap float32) *Recrutement {
	return &Recrutement{objectif: obj, stratAvant: sav, stratApres: sap, typeRecrutementAvant: trav, typeRecrutementApres: trap, pourcentagePlacesAvant: ppav, pourcentagePlacesApres: ppap}
}

func (r Recrutement) Objectif() float32 {
	return r.objectif
}

func (r Recrutement) StratAvant() StratParite {
	return r.stratAvant
}

func (r Recrutement) StratApres() StratParite {
	return r.stratApres
}

func (r Recrutement) TypeRecrutementAvant() TypeRecrutement {
	return r.typeRecrutementAvant
}

func (r Recrutement) TypeRecrutementApres() TypeRecrutement {
	return r.typeRecrutementApres
}

func (r Recrutement) PourcentagePlacesAvant() float32 {
	return r.pourcentagePlacesAvant
}

func (r Recrutement) PourcentagePlacesApres() float32 {
	return r.pourcentagePlacesApres
}
