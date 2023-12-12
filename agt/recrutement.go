package agt

import (
	"errors"
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

// ---------------------
//     Constructeurs
// ---------------------

func NewRecrutement(ent *Entreprise, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64) *Recrutement {
	return &Recrutement{entreprise: ent, objectif: obj, stratAvant: sav, stratApres: sap, typeRecrutementAvant: trav, typeRecrutementApres: trap, pourcentagePlacesAvant: ppav, pourcentagePlacesApres: ppap, chnl: ent.chnlRecrutement}
}

// ---------------------
//        Getters
// ---------------------

func (r Recrutement) Objectif() float64 {
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

func (r Recrutement) PourcentagePlacesAvant() float64 {
	return r.pourcentagePlacesAvant
}

func (r Recrutement) PourcentagePlacesApres() float64 {
	return r.pourcentagePlacesApres
}

// ---------------------
//  Utilitaires spécifiques
// ---------------------

// Permet de générer n candidat.es aléatoirement pour le besoin du recrutement
func (r Recrutement) GenererCandidats(nbCandidats int) (candidats []Employe, err error) {
	if nbCandidats < 0 {
		err := errors.New("erreur : nombre de candidats à générer négatif")
		return nil, err
	}
	candidats = make([]Employe, 0)
	for i := 0; i < nbCandidats; i++ {
		var genre Genre = genGenre()
		var anciennete int = 0 // anciennete = 0 car candidat
		var santeMentale int = 100
		var agresseur bool = genAgresseur(genre)
		var comportement Comportement = genComportement()
		var competence int = genCompetence()
		e := NewEmploye(genre, anciennete, santeMentale, agresseur, comportement, competence, r.entreprise)
		candidats = append(candidats, *e)
	}
	return candidats, nil
}

// ---------------------
//  Fonctions de recrutement
// ---------------------

// Recrutement si TypeRecrutement = Competences
// Les candidat.es les plus compétent.es sont recrutés. En cas d'égalité, le choix diffère en fonction de StratParite.
func RecrutementCompetencesEgales(nbARecruter int, strat StratParite, candidats []Employe) (embauches []Employe, err error) {
	if nbARecruter < 0 {
		err := errors.New("erreur : nombre de candidats à recruter négatif")
		return nil, err
	}
	if strat != PrioFemme && strat != PrioHomme && strat != Hasard {
		err := errors.New("erreur : stratégie de recrutement inconnue")
		return nil, err
	}
	// Pas d'erreur si len(candidats)=0 car dans ce cas, la fonction renvoie slice vide
	embauches = make([]Employe, 0)

	for len(embauches) < nbARecruter {
		maxCandidats := EmployeMaxCompetences(candidats)
		if len(maxCandidats) == 1 {
			embauches = append(embauches, maxCandidats[0])
			candidats = enleverEmploye(candidats, maxCandidats[0])
		} else if len(maxCandidats) > 1 {
			// Cas d'une égalité de compétence entre les candidat.es
			var idx int
			switch strat {
			case Hasard:
				// Un.e candidat.e au hasard parmi les plus compétent.es est recruté
				idx = rand.Intn(len(maxCandidats))
				embauches = append(embauches, maxCandidats[idx])
				candidats = enleverEmploye(candidats, maxCandidats[idx])
			case PrioFemme:
				// Une femme au hasard parmi les candidat.es est recrutée
				lFemmes := FiltreFemme(maxCandidats) // permet d'isoler les femmes parmi les candidat.es
				if len(lFemmes) > 0 {
					idx = rand.Intn(len(lFemmes))
					embauches = append(embauches, lFemmes[idx])
					candidats = enleverEmploye(candidats, lFemmes[idx])
				} else {
					// S'il n'y a pas de femmes parmi les candidats les plus compétents, on choisit au hasard
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmploye(candidats, maxCandidats[idx])
				}

			case PrioHomme:
				// Un homme au hasard parmi les candidat.es est recruté
				lHommes := FiltreHomme(maxCandidats) // permet d'isoler les hommes parmi les candidat.es
				if len(lHommes) > 0 {
					idx = rand.Intn(len(lHommes))
					embauches = append(embauches, lHommes[idx])
					candidats = enleverEmploye(candidats, lHommes[idx])
				} else {
					// S'il n'y a pas d'hommes parmi les candidats les plus compétents, on choisit au hasard
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmploye(candidats, maxCandidats[idx])
				}

			default:
				err = errors.New("erreur : stratégie de traitement des égalités de compétences inconnue")
				return nil, err
			}

		} else {
			err = errors.New("erreur : employeMaxCompetences ne renvoie aucun candidat")
			return nil, err
		}
	}
	return embauches, nil
}

// Recrutement si TypeRecrutement = PlacesReservees
// Parmi les candidats à recruter, un certain pourcentage est réservé aux femmes, peu importe leurs compétences
func RecrutementPlacesReservees(nbARecruter int, candidats []Employe, pourcentagePlace float64) (embauches []Employe, err error) {
	if nbARecruter < 0 {
		err := errors.New("erreur : nombre de candidats à recruter négatif")
		return nil, err
	}
	if pourcentagePlace < 0 || pourcentagePlace > 1 {
		err := errors.New("erreur : pourcentagePlace doit être compris entre 0 et 1")
		return nil, err
	}
	// Pas d'erreur si len(candidats)=0 car dans ce cas, la fonction renvoie slice vide
	embauches = make([]Employe, 0)
	// Hypothèse : si le résultat ne tombe pas juste, on arrondit le nombre de femmes au supérieur
	nbFemmesARecruter := int(math.Round(pourcentagePlace * float64(nbARecruter)))
	candidatsFemmes := FiltreFemme(candidats) // permet d'isoler les femmes parmi les candidat.es

	// 1ere etape : recruter les femmes les plus compétentes pour les places réservées
	for i := 0; i < nbFemmesARecruter; i++ {
		if len(candidatsFemmes) == 0 {
			break
		}
		maxCandidates := EmployeMaxCompetences(candidatsFemmes)
		embauches = append(embauches, maxCandidates[0])
		candidatsFemmes = enleverEmploye(candidatsFemmes, maxCandidates[0])
		candidats = enleverEmploye(candidats, maxCandidates[0])
	}
	// S'il n'y a pas assez de femmes dans les candidats pour toutes les places réservées, on recrute des hommes

	// Le reste des candidats sont sélectionnés uniquement pour leurs compétences
	reste := nbARecruter - len(embauches)
	for i := 0; i < reste; i++ {
		maxCandidats := EmployeMaxCompetences(candidats)
		idx := rand.Intn(len(maxCandidats))
		embauches = append(embauches, maxCandidats[idx])
		candidats = enleverEmploye(candidats, maxCandidats[idx])
	}

	return embauches, nil
}

// Fonction de recrutement générale que l'entreprise peut appeler à chaque pas de temps
// Réalise un recrutement à partir des choix renseignés par l'utilisation lors de l'initialisation
func (r Recrutement) Recruter(nbARecruter int) (embauches []Employe, err error) {
	if nbARecruter < 0 {
		err := errors.New("erreur : nombre de candidats à recruter négatif")
		return nil, err
	}

	// Génération des candidats

	candidats, err := r.GenererCandidats(constantes.NB_CANDIDATS_PAR_OFFRE * nbARecruter)
	if err != nil {
		return nil, err
	}

	if r.objectif == -1 {
		// L'utilisateur n'a pas défini un objectif de parité à atteindre

		// Tests cohérence
		if r.stratApres != StratVide {
			err := errors.New("erreur: une stratégie de parité après atteinte d'objectif a été définie mais pas d'objectif")
			return nil, err
		} else if r.typeRecrutementApres != Vide {
			err := errors.New("erreur: un type de recrutement après atteinte d'objectif a été défini mais pas d'objectif")
			return nil, err
		} else if r.pourcentagePlacesApres != -1 {
			err := errors.New("erreur: un pourcentage de parité à appliquer au recrutement après atteinte d'objectif a été défini mais pas d'objectif")
			return nil, err
		} else if r.typeRecrutementAvant == Competences && r.stratAvant == StratVide {
			err := errors.New("erreur: typeRecrutement=Compétences mais aucune stratégie définie en cas d'égalité des compétences")
			return nil, err
		} else if r.typeRecrutementAvant == PlacesReservees && r.pourcentagePlacesAvant == -1 {
			err := errors.New("erreur: typeRecrutement=PlacesReservees mais pas de pourcentage de places renseigné")
			return nil, err
		} else if r.typeRecrutementAvant == Competences && r.pourcentagePlacesAvant != -1 {
			err := errors.New("erreur: typeRecrutement=Compétences mais pourcentage de places à réserver renseigné")
			return nil, err
		} else if r.typeRecrutementAvant == PlacesReservees && r.stratAvant != StratVide {
			err := errors.New("erreur: typeRecrutement=PlacesReservées mais stratégie en cas d'égalité de compétences renseignée")
			return nil, err
		}

		// Choix de la fonction de recrutement à appeler
		if r.typeRecrutementAvant == Competences {
			embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
			if err != nil {
				return nil, err
			}
			return embauches, nil
		} else if r.typeRecrutementAvant == PlacesReservees {
			if r.pourcentagePlacesAvant < 0 || r.pourcentagePlacesAvant > 1 {
				err := errors.New("erreur : pourcentagePlaces doit être entre 0 et 1")
				return nil, err
			}
			embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesAvant)
			if err != nil {
				return nil, err
			}
			return embauches, nil
		} else {
			err := errors.New("erreur : typeRecrutement inconnu")
			return nil, err
		}

	} else {
		// L'utilisateur a défini un pourcentage de parité a atteindre

		// Tests cohérences
		if r.objectif < 0 || r.objectif > 1 {
			err := errors.New("erreur : objectif doit être entre 0 et 1")
			return nil, err
		} else if r.typeRecrutementAvant == Vide || r.typeRecrutementApres == Vide {
			err := errors.New("erreur : le type de recrutement avant et après objectif doit être défini")
			return nil, err
		} else if (r.typeRecrutementAvant == Competences && r.stratAvant == StratVide) || (r.typeRecrutementApres == Competences && r.stratApres == StratVide) {
			err := errors.New("erreur : typeRecrutement=Compétences mais aucune stratégie définie en cas d'égalité des compétences")
			return nil, err
		} else if (r.typeRecrutementAvant == PlacesReservees && r.pourcentagePlacesAvant == -1) || (r.typeRecrutementApres == PlacesReservees && r.pourcentagePlacesApres == -1) {
			err := errors.New("erreur : typeRecrutement=PlacesReservees mais pas de pourcentage de places renseigné")
			return nil, err
		} else if (r.typeRecrutementAvant == Competences && r.pourcentagePlacesAvant != -1) || (r.typeRecrutementApres == Competences && r.pourcentagePlacesApres != -1) {
			err := errors.New("erreur : typeRecrutement=Compétences mais pourcentage de places à réserver renseigné")
			return nil, err
		} else if (r.typeRecrutementAvant == PlacesReservees && r.stratAvant != StratVide) || (r.typeRecrutementApres == PlacesReservees && r.stratApres != StratVide) {
			err := errors.New("erreur : typeRecrutement=PlacesReservées mais stratégie en cas d'égalité de compétences renseignée")
			return nil, err
		}

		// Choix de la fonction de recrutement à appeler en fonction de si l'objectif est atteint ou non
		if r.objectif > r.entreprise.PourcentageFemmes() {
			// L'objectif n'est pas atteint, on applique TypeRecrutementAvant
			if r.typeRecrutementAvant == Competences {
				embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else if r.typeRecrutementAvant == PlacesReservees {
				if r.pourcentagePlacesAvant < 0 || r.pourcentagePlacesAvant > 1 {
					err := errors.New("erreur : pourcentagePlaces doit être entre 0 et 1")
					return nil, err
				}
				embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesAvant)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else {
				err := errors.New("erreur : typeRecrutement inconnu")
				return nil, err
			}
		} else {
			// L'objectif est atteint, on applique TypeRecrutementApres
			if r.typeRecrutementApres == Competences {
				embauches, err := RecrutementCompetencesEgales(nbARecruter, r.stratApres, candidats)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else if r.typeRecrutementApres == PlacesReservees {
				if r.pourcentagePlacesApres < 0 || r.pourcentagePlacesApres > 1 {
					err := errors.New("erreur : pourcentagePlaces doit être entre 0 et 1")
					return nil, err
				}
				embauches, err := RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesApres)
				if err != nil {
					return nil, err
				}
				return embauches, nil
			} else {
				err := errors.New("erreur : typeRecrutement inconnu")
				return nil, err
			}
		}
	}
}

// ---------------------
//  Logique de simulation
// ---------------------

func (r *Recrutement) Start() {
	log.Printf("Le service de recrutement est opérationnel")

	// Boucle de vie
	for {
		// Attend un message pour agir
		msg := <-r.chnl
		if msg.Act == RECRUTEMENT {
			embauches, err := r.Recruter(msg.Payload.(int))
			for _, emp := range embauches {
				go func(emp Employe) {
					emp.Start()
				}(emp)
			}
			if err != nil {
				r.chnl <- Communicateur_recrutement{ERREUR_RECRUTEMENT, err}
			} else {
				r.chnl <- Communicateur_recrutement{FIN_RECRUTEMENT, embauches}
			}
		} else {
			err := errors.New("erreur : mauvaise action du channel")
			r.chnl <- Communicateur_recrutement{ERREUR_RECRUTEMENT, err}
		}
	}
}