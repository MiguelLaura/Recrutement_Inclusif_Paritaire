package agt

// ------------ EMPLOYE ------------
type Genre int

const (
	Homme Genre = iota
	Femme
)

type Comportement float32

// A MODIFIER APRES DISCUSSION
const (
	Plainte100 Comportement = 1.0
	Plainte75  Comportement = 0.75
	Plainte50  Comportement = 0.5
	Plainte25  Comportement = 0.25
	Plainte0   Comportement = 0.0
)

type Employe struct {
	genre        Genre
	anciennete   int //entre 0 et 40
	santeMentale int //entre 0 et 100
	agresseur    bool
	comportement Comportement
	competence   int //entre 0 et 10
	entreprise   Entreprise
}

func NewEmploye(gen Genre, anc int, san int, ag bool, compor Comportement, compe int, ent Entreprise) *Employe {
	return &Employe{genre: gen, anciennete: anc, santeMentale: san, agresseur: ag, comportement: compor, competence: compe, entreprise: ent}
}

func (e Employe) Genre() Genre {
	return e.genre
}

func (e Employe) Anciennete() int {
	return e.anciennete
}

func (e Employe) GagnerAnciennete() {
	e.anciennete += 1
}

func (e Employe) SanteMentale() int {
	return e.santeMentale
}

func (e Employe) Agresseur() bool {
	return e.agresseur
}

func (e Employe) Comportement() Comportement {
	return e.comportement
}

func (e Employe) Competence() int {
	return e.competence
}

func (e Employe) Entreprise() Entreprise {
	return e.entreprise
}

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
