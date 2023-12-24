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

// Permet de trouver un employé à partir d'une condition et d'un certain indice et renvoie l'employé et son indice
func TrouverEmploye(tab []*Employe, f func(*Employe) bool, from int) (index int, val *Employe) {
	if from < 0 {
		from = 0
	} else if from >= len(tab) {
		return -1, nil
	}

	for idx := from; idx < len(tab); idx++ {
		if f(tab[idx]) {
			return idx, tab[idx]
		}
	}

	return -1, nil
}

// Récupère index d'un employé au sein d'une slice d'employés
func obtenirIndexEmploye(emp []*Employe, e *Employe) int {
	for idx, val := range emp {
		if e.Id() == val.Id() {
			return idx
		}
	}
	return -1
}

// Renvoie la liste emp sans l'employé e
func enleverEmploye(emp []*Employe, e *Employe) []*Employe {
	if len(emp) <= 0 {
		vide := make([]*Employe, 0)
		return vide
	}
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

func genCompetence() int {
	// Loi normale définie pour modéliser les compétences
	// On veut que les compétences tournent autour de 5 sans trop s’éparpiller autour
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

func EstFemme(e *Employe) bool {
	return e.Genre() == Femme
}

func EstHomme(e *Employe) bool {
	return e.Genre() == Homme
}

// Renvoie l'employé ou les employés avec le maximum de compétence
func EmployeMaxCompetences(candidats []*Employe) (emp []*Employe) {
	emp = make([]*Employe, 0)
	var max int = 0
	for _, value := range candidats {
		if value.competence == max {
			emp = append(emp, value)
		} else if value.competence > max {
			max = value.competence
			emp = make([]*Employe, 0)
			emp = append(emp, value)
		}
	}
	return emp
}

// Renvoie le slice des employées femmes à partir d'une liste d'employé.es
func FiltreFemme(employes []*Employe) (f []*Employe) {
	f = make([]*Employe, 0)

	idx, _ := TrouverEmploye(employes, EstFemme, 0)
	for idx != -1 {
		f = append(f, employes[idx])
		idx, _ = TrouverEmploye(employes, EstFemme, idx+1)
	}
	return f
}

// Renvoie le slice des employés hommes à partir d'une liste d'employé.es
func FiltreHomme(employes []*Employe) (f []*Employe) {
	f = make([]*Employe, 0)

	idx, _ := TrouverEmploye(employes, EstHomme, 0)
	for idx != -1 {
		f = append(f, employes[idx])
		idx, _ = TrouverEmploye(employes, EstHomme, idx+1)
	}
	return f
}

// Renvoie string à partir de la valeur de type Genre
func StringGenre(g Genre) string {
	if g == Homme {
		return "H"
	} else if g == Femme {
		return "F"
	} else {
		return ""
	}
}
