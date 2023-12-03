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
		employesInit = append(employesInit, *GenererEmployeInit(e, Femme))
	}
	for i := 0; i < nbHommes; i++ {
		employesInit = append(employesInit, *GenererEmployeInit(e, Homme))
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
//        Setter
// ---------------------

//     + AjouterRecrutement(recrut Recrutement)

// ---------------------
//     En cours d'année (appelés par les employés)
// ---------------------

//     + RecevoirDemission(Employe)
//	   + RecevoirDepression(Employe)
// 	   + RecevoirRetraite(Employe)
//     + RecevoirPlainte(plaignant Employe, accuse Employe)
//     + MettreAJourCA(santeMentale : int, competence : int)
// 	   + NotifierAction()

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

//     + NombreEmployes() : int
//     + PourcentageFemmes() : float
//     + SupprimerEmploye(employe)
//     + EnvoyerEmploye() : *Employe
//     + AjusterImpactFemmes()
