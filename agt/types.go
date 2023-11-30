package agt

import (
	"math"
)

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
