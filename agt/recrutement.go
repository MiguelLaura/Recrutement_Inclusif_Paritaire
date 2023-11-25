package agt

import (
	"errors"
	"fmt"
	"math/rand"
)

// Comment le recrutement récupère les informations de l'entreprise ?

// a mettre en minuscule apres tests
func GenererCandidats(nb_candidats int, ent Entreprise) []Employe {
	// anciennete = 0 car candidat
	var emp []Employe
	emp = make([]Employe, 0)
	for i := 0; i < nb_candidats; i++ {
		var e Employe
		e.anciennete = 0
		e.santeMentale = 100
		e.genre = genGenre()
		e.agresseur = genAgresseur(e.genre)
		e.comportement = genComportement()
		e.entreprise = ent
		e.competence = genCompetence()
		emp = append(emp, e)
		fmt.Println("employe ", i, ": ", e)
	}
	fmt.Println("emp : ", emp)
	return emp
}

func RecrutementCompetencesEgales(nbARecruter int, strat StratParite, candidats []Employe) (embauches []Employe, err error) {
	embauches = make([]Employe, 0)

	// nbARecruter ne doit jamais depasser 10 actuellement -> a ameliorer
	for len(embauches) < nbARecruter {
		maxCandidats := EmployeMaxCompetences(candidats)
		if len(maxCandidats) == 1 {
			embauches = append(embauches, maxCandidats[0])
			candidats = enleverEmployer(candidats, maxCandidats[0])
		} else if len(maxCandidats) > 1 {
			// appliquer differentes strategies
			switch strat {
			case Hasard:
				h := rand.Intn(len(maxCandidats))
				embauches = append(embauches, maxCandidats[h])
				candidats = enleverEmployer(candidats, maxCandidats[h])
			case PrioFemme:
			case PrioHomme:
			default:
				err = errors.New("Stratégie de traitement des égalités de compétences inconnue")
			}
		} else {
			err = errors.New("EmployeMaxCompetences ne renvoie aucun candidat")
		}
	}
	return embauches, nil
}
