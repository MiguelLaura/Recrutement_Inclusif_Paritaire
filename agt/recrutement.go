package agt

import (
	"errors"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

// Pour le recrutement, des candidats sont générés aléatoirement
func (r Recrutement) GenererCandidats(nb_candidats int) (candidats []Employe, err error) {
	if nb_candidats < 0 {
		err := errors.New("Nombre de candidats à générer négatif")
		return nil, err
	}
	candidats = make([]Employe, 0)
	for i := 0; i < nb_candidats; i++ {
		var genre Genre = genGenre()
		// anciennete = 0 car candidat
		var anciennete int = 0
		var santeMentale int = 100
		var agresseur bool = genAgresseur(genre)
		var comportement Comportement = genComportement()
		var competence int = genCompetence()
		e := NewEmploye(genre, anciennete, santeMentale, agresseur, comportement, competence, r.entreprise)
		candidats = append(candidats, *e)
	}
	return candidats, nil
}

// nbARecruter ne doit jamais depasser 10 actuellement -> a ameliorer car nb_candidats=10
func RecrutementCompetencesEgales(nbARecruter int, strat StratParite, candidats []Employe) (embauches []Employe, err error) {
	if nbARecruter < 0 {
		err := errors.New("Nombre de candidats à recruter négatif")
		return nil, err
	}
	if strat != PrioFemme && strat != PrioHomme && strat != Hasard {
		err := errors.New("Stratégie de recrutement inconnue")
		return nil, err
	}
	// On ne verifie pas si candidats est vide car dans ce cas, liste vide renvoyée ce qui est cohérent
	embauches = make([]Employe, 0)

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
				embauches = append(embauches, maxCandidats[idx])
				candidats = enleverEmployer(candidats, maxCandidats[idx])
			case PrioFemme:
				l_femmes := FiltreFemme(maxCandidats)
				// si parmi les plus compétents, il y a des femmes, on choisit parmi elles sinon au hasard
				if len(l_femmes) > 0 {
					idx = rand.Intn(len(l_femmes))
					embauches = append(embauches, l_femmes[idx])
					candidats = enleverEmployer(candidats, l_femmes[idx])
				} else {
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmployer(candidats, maxCandidats[idx])
				}

			case PrioHomme:
				l_hommes := FiltreHomme(maxCandidats)
				// si parmi les plus compétents, il y a des hommes, on choisit parmi eux sinon au hasard
				if len(l_hommes) > 0 {
					idx = rand.Intn(len(l_hommes))
					embauches = append(embauches, l_hommes[idx])
					candidats = enleverEmployer(candidats, l_hommes[idx])
				} else {
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmployer(candidats, maxCandidats[idx])
				}

			default:
				err = errors.New("Stratégie de traitement des égalités de compétences inconnue")
				return nil, err
			}

		} else {
			err = errors.New("EmployeMaxCompetences ne renvoie aucun candidat")
		}
	}
	return embauches, nil
}

func RecrutementPlacesReservees(nbARecruter int, candidats []Employe, pourcentagePlace float64) (embauches []Employe, err error) {
	if nbARecruter < 0 {
		err := errors.New("Nombre de candidats à recruter négatif")
		return nil, err
	}
	if pourcentagePlace < 0 || pourcentagePlace > 1 {
		err := errors.New("pourcentagePlace doit être compris entre 0 et 1")
		return nil, err
	}
	// On ne verifie pas si candidats est vide car dans ce cas, liste vide renvoyée ce qui est cohérent
	embauches = make([]Employe, 0)
	// Hypothèse : si le résultat ne tombe pas juste, on arrondit le nombre de femmes au supérieur
	nbFemmesARecruter := int(math.Round(pourcentagePlace * float64(nbARecruter)))
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
	// S'il n'y a pas assez de femmes dans les candidats, on recrute des hommes

	// Le reste des candidats sont sélectionnés uniquement pour leurs compétences
	reste := nbARecruter - len(embauches)
	for i := 0; i < reste; i++ {
		maxCandidats := EmployeMaxCompetences(candidats)
		idx := rand.Intn(len(maxCandidats))
		embauches = append(embauches, maxCandidats[idx])
		candidats = enleverEmployer(candidats, maxCandidats[idx])
	}

	return embauches, nil
}

func (r Recrutement) Recruter(nbARecruter int) (embauches []Employe, err error) {
	// Cas où l'utilisateur n'a pas défini un objectif de parité à atteindre
	candidats, err := r.GenererCandidats(constantes.NBCANDIDATS)
	if err != nil {
		return nil, err
	}
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
			embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesAvant))
			if err != nil {
				return nil, err
			}
			return embauches, nil
		} else {
			err := errors.New("Les stratégies de recrutement n'ont pas été correctement définies")
			return nil, err
		}
	} else {
		if r.objectif < r.entreprise.PourcentageFemmes() {
			if r.stratAvant != -1 {
				// On recruter en fonction des competences et s'il y a égalité, on utilise la stratégie définie par l'utilisateur
				embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else if r.typeRecrutementAvant != -1 {
				embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesAvant))
				if err != nil {
					return nil, err
				}
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
				embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, float64(r.pourcentagePlacesApres))
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else {
				err := errors.New("Les stratégies de recrutement n'ont pas été correctement définies")
				return nil, err
			}
		}
	}
}
