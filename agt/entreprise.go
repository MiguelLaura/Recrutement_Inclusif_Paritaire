package agt

import "math"

type Entreprise struct {
	employes    []Employe
	departs     []Employe
	plaintes    [][]Employe
	depression  int
	recrutement Recrutement
	ca          float64
	nbAction    int
}

// ---------------------
//     Constructeur
// ---------------------

// La fonction NewEntreprise doit créer l'entreprise et générer les employés de façon à respecter le quota de parité initial
func NewEntreprise(nbEmployesInit int, pariteInit float32) *Entreprise {
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
	e.ca = 0.0
	return e
}

// ---------------------
//        Getters
// ---------------------

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

// ---------------------
//     En cours d'année (appelées par les employés)
// ---------------------

func (ent Entreprise) RecevoirDemission(emp *Employe) {
	ent.departs = append(ent.departs, *emp)
}

func (ent Entreprise) RecevoirDepression(emp *Employe) {
	ent.depression += 1
	ent.departs = append(ent.departs, *emp)
}

func (ent Entreprise) RecevoirRetraite(emp *Employe) {
	ent.departs = append(ent.departs, *emp)
}

func (e Entreprise) RecevoirPlainte(plaignant *Employe, accuse *Employe) {
	e.plaintes = append(e.plaintes, []Employe{*plaignant, *accuse})
}

// METTRE A JOUR LA FORMULE
func (e Entreprise) MettreAJourCA(santeMentale int, competence int) {
	e.ca += float64(santeMentale) * float64(competence)
}

func (e Entreprise) NotifierAction() {
	e.nbAction += 1
}

// ---------------------
//     Fin d'année
// ---------------------

//     + GestionRecrutements()
//     + GestionPlaintes()
//     + GestionDeparts()
//     + CalculerBenefice()
//     + ObtenirIndicateursSante() : map[string]float

// ---------------------
//  Logique de simulation
// ---------------------

//     + Start() -> start les employés
// 	   + FinirCycle()

// ---------------------
//     Autres
// ---------------------

func (e Entreprise) AjouterRecrutement(recrut Recrutement) {
	e.recrutement = recrut
}

func (ent Entreprise) nbEmployes() int {
	return len(ent.employes)
}

func (ent Entreprise) PourcentageFemmes() float64 {
	femmes := FiltreFemme(ent.employes)
	return float64(len(femmes)) / float64(len(ent.employes))
}

//     + NombreEmployes() : int -> DEJA FAIT PAR SOLENN SUR UNE AUTRE BRANCHE
//     + PourcentageFemmes() : float -> DEJA FAIT PAR SOLENN SUR UNE AUTRE BRANCHE
//     + SupprimerEmploye(employe)
//     + EnvoyerEmploye() : *Employe
//     + AjusterImpactFemmes()
