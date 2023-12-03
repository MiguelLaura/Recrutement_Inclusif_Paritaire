package agt

import (
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

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
	ent := new(Entreprise)

	var nbFemmes int = int(math.Round(float64(nbEmployesInit) * float64(pariteInit)))
	var nbHommes int = nbEmployesInit - nbFemmes
	var employesInit []Employe

	for i := 0; i < nbFemmes; i++ {
		employesInit = append(employesInit, *GenererEmployeInit(&ent, Femme))
	}
	for i := 0; i < nbHommes; i++ {
		employesInit = append(employesInit, *GenererEmployeInit(&ent, Homme))
	}
	ent.employes = employesInit
	ent.departs = make([]Employe, 0)
	ent.plaintes = make([][]Employe, 0)
	ent.ca = 0.0
	return ent
}

// ---------------------
//        Getters
// ---------------------

func (ent Entreprise) Employes() []Employe {
	return ent.employes
}

func (ent Entreprise) Departs() []Employe {
	return ent.departs
}

func (ent Entreprise) Plaintes() [][]Employe {
	return ent.plaintes
}

func (ent Entreprise) Recrutement() Recrutement {
	return ent.recrutement
}

func (ent Entreprise) Ca() float64 {
	return ent.ca
}

// ---------------------
//     En cours d'année (appelées par les employés)
// ---------------------

func (ent *Entreprise) RecevoirDemission(emp *Employe) {
	ent.departs = append(ent.departs, *emp)
}

func (ent *Entreprise) RecevoirDepression(emp *Employe) {
	ent.depression += 1
	ent.departs = append(ent.departs, *emp)
}

func (ent *Entreprise) RecevoirRetraite(emp *Employe) {
	ent.departs = append(ent.departs, *emp)
}

func (ent *Entreprise) RecevoirPlainte(plaignant *Employe, accuse *Employe) {
	ent.plaintes = append(ent.plaintes, []Employe{*plaignant, *accuse})
}

// METTRE A JOUR LA FORMULE
func (ent *Entreprise) MettreAJourCA(santeMentale int, competence int) {
	ent.ca += float64(santeMentale) * float64(competence)
}

func (ent *Entreprise) NotifierAction() {
	ent.nbAction += 1
}

// ---------------------
//     Fin d'année
// ---------------------

func (ent *Entreprise) GestionRecrutements() (err error) {
	nbARecruter := float64(ent.nbEmployes()) * constantes.POURCENTAGE_RECRUTEMENT
	embauches, err := ent.recrutement.Recruter(int(math.Round(nbARecruter)))
	if err != nil {
		return err
	}

	ent.employes = append(ent.employes, embauches...)
	return nil
}

func (ent *Entreprise) GestionDeparts() {
	for _, emp := range ent.employes {
		ent.employes = enleverEmploye(ent.employes, emp)
	}
}

//     + GestionPlaintes()
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

func (ent *Entreprise) ajouterRecrutement(recrut Recrutement) {
	ent.recrutement = recrut
}

func (ent Entreprise) nbEmployes() int {
	return len(ent.employes)
}

func (ent Entreprise) PourcentageFemmes() float64 {
	femmes := FiltreFemme(ent.employes)
	return float64(len(femmes)) / float64(len(ent.employes))
}

func (ent Entreprise) EnvoyerEmploye() *Employe {
	idx := rand.Intn(len(ent.employes))
	emp := ent.employes[idx]
	return &emp
}

// func (ent Entreprise) ajusterImpactFemmes() {
// }
