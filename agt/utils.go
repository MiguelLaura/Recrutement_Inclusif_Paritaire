package agt

import (
	"math/rand"
	"sync"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

// GENERAL

func Trouver_Employe(tab []Employe, f func(Employe) bool) (index int, val Employe) {
	var e Employe
	for idx, value := range tab {
		if f(value) {
			return idx, value
		}
	}
	return -1, e
}

func Trouver_Employe_conc(tab []Employe, f func(Employe) bool, n int) (index int, value Employe) {

	c := make(chan int)

	go func() {
		var wg sync.WaitGroup
		defer close(c)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				idx, _ := Trouver_Employe(tab[i*len(tab)/n:(len(tab)/n+i*len(tab)/n)], f)
				if idx != -1 {
					c <- (idx + i*len(tab)/n)
				}
			}(i)
		}
		wg.Wait()
	}()
	idx_min := len(tab) - 1
	for i := range c {
		if i != -1 {
			idx_min = min(idx_min, i)
		}
	}
	return idx_min, tab[idx_min]
}

// Problème : pour vérifier que les employés sont égaux, il faut vérifier égalité entreprise
// pour vérifier que les entreprises sont égales, il faut vérifier que les employés sont égaux
func listesEmployesEgales(l1 []Employe, l2 []Employe)

func entreprisesEgales(ent1 Entreprise, ent2 Entreprise) bool

func employesEgaux(e1 Employe, e2 Employe) bool {
	// ajouter verif entreprise si on regle le pb au dessus
	return (e1.Agresseur() == e2.Agresseur() && e1.Anciennete() == e2.Anciennete() && e1.Competence() == e2.Competence() && e1.Comportement() == e2.Comportement() && e1.Genre() == e2.Genre() && e1.SanteMentale() == e2.SanteMentale())
}

func obtenirIndexEmploye(emp []Employe, e Employe) int {
	for idx, val := range emp {
		if employesEgaux(e, val) {
			return idx
		}
	}
	return -1
}

// renvoie la liste emp sans l'employe e
func enleverEmployer(emp []Employe, e Employe) []Employe {
	i := obtenirIndexEmploye(emp, e)
	emp[i] = emp[len(emp)-1]
	return emp[:len(emp)-1]
}

// GENERATION
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
	return rand.Intn(10)
}

// RECRUTEMENT
func EstFemme(e Employe) bool {
	return e.Genre() == Femme
}

func EstHomme(e Employe) bool {
	return e.Genre() == Homme
}

// Recherche avec des goroutines ? 4 par exemple ?
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

func FiltreFemme(employes []Employe) (f []Employe) {
	emp := make([]Employe, 0)
	copy(emp, employes)
	f = make([]Employe, 0)

	// on utilise 4 goroutines -> voir si y a plus opti
	idx, e := Trouver_Employe_conc(emp, EstFemme, 4)
	for idx != -1 {
		f = append(f, emp[idx])
		emp = enleverEmployer(emp, e)
		idx, e = Trouver_Employe_conc(emp, EstFemme, 4)
	}
	return f
}
