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
			var idx int
			switch strat {
			case Hasard:
				idx = rand.Intn(len(maxCandidats))
			case PrioFemme:
				l_femmes := FiltreFemme(candidats)
				idx = rand.Intn(len(l_femmes))
			case PrioHomme:
				l_hommes := FiltreHomme(candidats)
				idx = rand.Intn(len(l_hommes))
			default:
				err = errors.New("Stratégie de traitement des égalités de compétences inconnue")
			}
			embauches = append(embauches, maxCandidats[idx])
			candidats = enleverEmployer(candidats, maxCandidats[idx])
		} else {
			err = errors.New("EmployeMaxCompetences ne renvoie aucun candidat")
		}
	}
	return embauches, nil
}

func RecrutementPlacesReservees(nbARecruter int, candidats []Employe, pourcentagePlace float64) (embauches []Employe) {
	embauches = make([]Employe, 0)
	nbFemmesARecruter := int(pourcentagePlace * float64(nbARecruter))
	candidatsFemmes := FiltreFemme(candidats)

	for i := 0; i < nbFemmesARecruter; i++ {
		if len(candidatsFemmes) == 0 {
			break
		}
		// On récupère les femmes avec la compétence maximum
		maxCandidates := EmployeMaxCompetences(candidatsFemmes)
		embauches = append(embauches, maxCandidates[0])
		candidatsFemmes = enleverEmployer(candidatsFemmes, maxCandidates[0])
		candidats = enleverEmployer(candidats, maxCandidates[0])
	}
	// S'il n'y a pas assez de femmes dans les candidats que faire ?

	// Le reste des candidats sont sélectionnés uniquement pour leurs compétences
	reste := nbARecruter - len(embauches)
	for i := 0; i < reste; i++ {
		maxCandidats := EmployeMaxCompetences(candidats)
		idx := rand.Intn(len(maxCandidats))
		embauches = append(embauches, maxCandidats[idx])
		candidats = enleverEmployer(candidats, maxCandidats[idx])
	}

	return embauches
}

func (r Recrutement) Recruter(nbARecruter int, ent Entreprise) (embauches []Employe, err error) {
	// Cas où l'utilisateur n'a pas défini un objectif de parité à atteindre
	candidats := GenererCandidats(10, ent)
	if r.objectif == -1 {
		// Faire des test : si stratAvant et typeRecrutementAvant sont définies toutes les deux, erreur
		if r.stratAvant != -1 {
			// On recruter en fonction des competences et s'il y a égalité, on utilise la stratégie définie par l'utilisateur
			embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
			if err != nil {
				return nil, err
			}
			return embauches, nil
		} else if r.typeRecrutementAvant != -1 {
			// Verif coherence valeur pourcentage
			embauches = RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesAvant))
			return embauches, nil
		} else {
			err := errors.New("Les stratégies de recrutement n'ont pas été correctement définies")
			return nil, err
		}
	} else {
		if r.objectif < ent.PourcentageFemmes() {
			if r.stratAvant != -1 {
				// On recruter en fonction des competences et s'il y a égalité, on utilise la stratégie définie par l'utilisateur
				embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else if r.typeRecrutementAvant != -1 {
				// Verif coherence valeur pourcentage
				embauches = RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesAvant))
				return embauches, nil
			} else {
				err := errors.New("Les stratégies de recrutement n'ont pas été correctement définies")
				return nil, err
			}
		} else {
			if r.stratApres != -1 {
				// On recruter en fonction des competences et s'il y a égalité, on utilise la stratégie définie par l'utilisateur
				embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratApres, candidats)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else if r.typeRecrutementApres != -1 {
				// Verif coherence valeur pourcentage
				embauches = RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesApres))
				return embauches, nil
			} else {
				err := errors.New("Les stratégies de recrutement n'ont pas été correctement définies")
				return nil, err
			}
		}
	}
}
