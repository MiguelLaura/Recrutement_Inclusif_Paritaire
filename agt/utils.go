package agt

import (
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
	"gonum.org/v1/gonum/stat/distuv"
)

// ---------------------
//        Général
// ---------------------

// Vérifie la présence d'un employé dans un slice employé
func EstDansSliceEmploye(employes []Employe, e Employe) bool {
	for _, val := range employes {
		if e.Id() == val.Id() {
			return true
		}
	}
	return false
}

// Permet de trouver un employé à partir d'une condition et renvoie l'employé et son indice
func TrouverEmploye(tab []Employe, f func(Employe) bool) (index int, val Employe) {
	var e Employe
	for idx, value := range tab {
		if f(value) {
			return idx, value
		}
	}
	return -1, e
}

// Récupère index d'un employé au sein d'une slice d'employés
func obtenirIndexEmploye(emp []Employe, e Employe) int {
	for idx, val := range emp {
		if e.Id() == val.Id() {
			return idx
		}
	}
	return -1
}

// Renvoie la liste emp sans l'employé e
func enleverEmploye(emp []Employe, e Employe) []Employe {
	i := obtenirIndexEmploye(emp, e)
	emp[i] = emp[len(emp)-1]
	return emp[:len(emp)-1]
}

// ---------------------
//       Génération
// ---------------------

func genAgresseur(genre Genre) (agg bool) {

	if genre == Homme {
		if rand.Float64() <= constantes.POURCENTAGE_AGRESSEUR_H {
			agg = true
		}
	} else {
		if rand.Float64() <= constantes.POURCENTAGE_AGRESSEUR_H {
			agg = true
		}
	}
	return agg
}

func genGenre() (genre Genre) {
	if rand.Float64() <= 0.5 {
		genre = Femme
	} else {
		genre = Homme
	}
	return genre
}

func genComportement() Comportement {
	r := rand.Float64()
	var compor Comportement
	if r >= 0 && r < 0.2 {
		compor = Plainte0
	} else if r >= 0.2 && r < 0.4 {
		compor = Plainte25
	} else if r >= 0.4 && r < 0.6 {
		compor = Plainte50
	} else if r >= 0.6 && r < 0.8 {
		compor = Plainte75
	} else {
		compor = Plainte100
	}
	return compor
}

func genCompetence() int {
	// Loi normale définie pour modéliser les compétences
	var loiNormale = distuv.Normal{
		Mu:    5,
		Sigma: 3,
	}
	// Permet de ne pas avoir de compétence négative et de ne pas aller au dessus du seuil max de compétence
	return int(math.Abs(loiNormale.Rand())) % (constantes.COMPETENCE_MAX + 1)
}

// ---------------------
//     Recrutement
// ---------------------

func EstFemme(e Employe) bool {
	return e.Genre() == Femme
}

func EstHomme(e Employe) bool {
	return e.Genre() == Homme
}

// Renvoie l'employé ou les employés avec le maximum de compétence
func EmployeMaxCompetences(candidats []Employe) (emp []Employe) {
	emp = make([]Employe, 0)
	var max int = 0
	for _, value := range candidats {
		if value.competence == max {
			emp = append(emp, value)
		} else if value.competence > max {
			max = value.competence
			emp = make([]Employe, 0)
			emp = append(emp, value)
		}
	}
	return emp
}

// Renvoie le slice des employées femmes à partir d'une liste d'employé.es
func FiltreFemme(employes []Employe) (f []Employe) {
	emp := make([]Employe, len(employes))
	copy(emp, employes)
	f = make([]Employe, 0)

	idx, e := TrouverEmploye(emp, EstFemme)
	for idx != -1 {
		f = append(f, emp[idx])
		emp = enleverEmploye(emp, e)
		idx, e = TrouverEmploye(emp, EstFemme)
	}
	return f
}

// Renvoie le slice des employés hommes à partir d'une liste d'employé.es
func FiltreHomme(employes []Employe) (f []Employe) {
	emp := make([]Employe, len(employes))
	copy(emp, employes)
	f = make([]Employe, 0)

	idx, e := TrouverEmploye(emp, EstHomme)
	for idx != -1 {
		f = append(f, emp[idx])
		emp = enleverEmploye(emp, e)
		idx, e = TrouverEmploye(emp, EstHomme)
	}
	return f
}
