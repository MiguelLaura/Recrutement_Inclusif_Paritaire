package agt

import (
	"log"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// ---------------------
//     Constructeurs
// ---------------------

func NewRecrutement(ent *Entreprise, obj float64, sav StratParite, sap StratParite, trav TypeRecrutement, trap TypeRecrutement, ppav float64, ppap float64, logger *logger.Loggers) *Recrutement {
	return &Recrutement{
		entreprise:             ent,
		objectif:               obj,
		stratAvant:             sav,
		stratApres:             sap,
		typeRecrutementAvant:   trav,
		typeRecrutementApres:   trap,
		pourcentagePlacesAvant: ppav,
		pourcentagePlacesApres: ppap,
		chnl:                   ent.chnlRecrutement,
		fin:                    false,
		logger:                 logger,
	}
}

// ---------------------
//        Getters
// ---------------------

func (r *Recrutement) Entreprise() *Entreprise {
	return r.entreprise
}

func (r *Recrutement) Objectif() float64 {
	return r.objectif
}

func (r *Recrutement) StratAvant() StratParite {
	return r.stratAvant
}

func (r *Recrutement) StratApres() StratParite {
	return r.stratApres
}

func (r *Recrutement) TypeRecrutementAvant() TypeRecrutement {
	return r.typeRecrutementAvant
}

func (r *Recrutement) TypeRecrutementApres() TypeRecrutement {
	return r.typeRecrutementApres
}

func (r *Recrutement) PourcentagePlacesAvant() float64 {
	return r.pourcentagePlacesAvant
}

func (r *Recrutement) PourcentagePlacesApres() float64 {
	return r.pourcentagePlacesApres
}

func (r *Recrutement) Chnl() chan CommunicateurRecrutement {
	return r.chnl
}

func (r *Recrutement) Fin() bool {
	return r.fin
}

func (r *Recrutement) Logger() *logger.Loggers {
	return r.logger
}

// ---------------------
//        Setters
// ---------------------

func (r *Recrutement) SetEntreprise(entreprise *Entreprise) {
	r.entreprise = entreprise
}

func (r *Recrutement) SetObjectif(objectif float64) {
	r.objectif = objectif
}

func (r *Recrutement) SetStratAvant(stratParite StratParite) {
	r.stratAvant = stratParite
}

func (r *Recrutement) SetStratApres(stratParite StratParite) {
	r.stratApres = stratParite
}

func (r *Recrutement) SetTypeRecrutementAvant(typeRecrutement TypeRecrutement) {
	r.typeRecrutementAvant = typeRecrutement
}

func (r *Recrutement) SetTypeRecrutementApres(typeRecrutement TypeRecrutement) {
	r.typeRecrutementApres = typeRecrutement
}

func (r *Recrutement) SetPourcentagePlacesAvant(pourcentagePlacesAvant float64) {
	r.pourcentagePlacesAvant = pourcentagePlacesAvant
}

func (r *Recrutement) SetPourcentagePlacesApres(pourcentagePlacesApres float64) {
	r.pourcentagePlacesApres = pourcentagePlacesApres
}

func (r *Recrutement) SetChnl(chnl chan CommunicateurRecrutement) {
	r.chnl = chnl
}

func (r *Recrutement) SetFin(fin bool) {
	r.fin = fin
}

func (r *Recrutement) SetLogger(logger *logger.Loggers) {
	r.logger = logger
}

// ---------------------
//  Utilitaires spécifiques
// ---------------------

// Permet de générer n candidat.es aléatoirement pour le besoin du recrutement
func (r *Recrutement) GenererCandidats(nbCandidats int) (candidats []*Employe) {
	if nbCandidats < 0 {
		return nil
	}
	candidats = make([]*Employe, 0)
	for i := 0; i < nbCandidats; i++ {
		var genre Genre = genGenre()
		var anciennete int = 0 // anciennete = 0 car candidat
		var santeMentale int = 100
		var agresseur bool = genAgresseur(genre)
		var competence int = genCompetence()
		e := NewEmploye(genre, anciennete, santeMentale, agresseur, competence, r.entreprise, r.logger)
		candidats = append(candidats, e)
	}
	return candidats
}

// ---------------------
//  Fonctions de recrutement
// ---------------------

// Recrutement si TypeRecrutement = Competences
// Les candidat.es les plus compétent.es sont recrutés. En cas d'égalité, le choix diffère en fonction de StratParite.
func (r *Recrutement) RecrutementCompetencesEgales(nbARecruter int, strat StratParite, candidats []*Employe) (embauches []*Employe) {
	if nbARecruter < 0 {
		return nil
	}
	if strat != PrioFemme && strat != PrioHomme && strat != Hasard {
		return nil
	}
	embauches = make([]*Employe, 0)
	r.logger.LogfType(LOG_RECRUTEMENT, "Le service RH organise une campagne de recrutement pour %d postes", nbARecruter)
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
				r.logger.LogfType(LOG_RECRUTEMENT, "Plusieurs candidat.es avec des compétences équivalentes ont postulé. On les départage par le hasard.")
				idx = rand.Intn(len(maxCandidats))
				embauches = append(embauches, maxCandidats[idx])
				candidats = enleverEmploye(candidats, maxCandidats[idx])
			case PrioFemme:
				// Une femme au hasard parmi les candidat.es est recrutée
				lFemmes := FiltreFemme(maxCandidats) // permet d'isoler les femmes parmi les candidat.es
				if len(lFemmes) > 0 {
					r.logger.LogfType(LOG_RECRUTEMENT, "Plusieurs candidat.es avec des compétences équivalentes ont postulé. On privilégie les femmes.")
					idx = rand.Intn(len(lFemmes))
					embauches = append(embauches, lFemmes[idx])
					candidats = enleverEmploye(candidats, lFemmes[idx])
				} else {
					r.logger.LogfType(LOG_RECRUTEMENT, "Plusieurs candidat.es avec des compétences équivalentes ont postulé. On veut privilégier les femmes, mais il n'y en a pas dans le groupe. Un homme est recruté")
					// S'il n'y a pas de femmes parmi les candidats les plus compétents, on choisit au hasard
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmploye(candidats, maxCandidats[idx])
				}

			case PrioHomme:
				// Un homme au hasard parmi les candidat.es est recruté
				lHommes := FiltreHomme(maxCandidats) // permet d'isoler les hommes parmi les candidat.es
				if len(lHommes) > 0 {
					r.logger.LogfType(LOG_RECRUTEMENT, "Plusieurs candidat.es avec des compétences équivalentes ont postulé. On privilégie les hommes.")
					idx = rand.Intn(len(lHommes))
					embauches = append(embauches, lHommes[idx])
					candidats = enleverEmploye(candidats, lHommes[idx])
				} else {
					r.logger.LogfType(LOG_RECRUTEMENT, "Plusieurs candidat.es avec des compétences équivalentes ont postulé. On veut privilégier les hommes, mais il n'y en a pas dans le groupe. Une femme est recrutée")
					// S'il n'y a pas d'hommes parmi les candidats les plus compétents, on choisit au hasard
					idx = rand.Intn(len(maxCandidats))
					embauches = append(embauches, maxCandidats[idx])
					candidats = enleverEmploye(candidats, maxCandidats[idx])
				}

			default:
				return nil
			}

		} else {
			return nil
		}
	}
	return embauches
}

// Recrutement si TypeRecrutement = PlacesReservees
// Parmi les candidats à recruter, un certain pourcentage est réservé aux femmes, peu importe leurs compétences
func (r *Recrutement) RecrutementPlacesReservees(nbARecruter int, candidats []*Employe, pourcentagePlace float64) (embauches []*Employe) {
	if nbARecruter < 0 {
		return nil
	}
	if pourcentagePlace < 0 || pourcentagePlace > 1 {
		return nil
	}
	r.logger.LogfType(LOG_RECRUTEMENT, "Le service RH organise une campagne de recrutement pour %d postes", nbARecruter)
	// Pas d'erreur si len(candidats)=0 car dans ce cas, la fonction renvoie slice vide
	embauches = make([]*Employe, 0)
	// Hypothèse : si le résultat ne tombe pas juste, on arrondit le nombre de femmes au supérieur
	nbFemmesARecruter := int(math.Round(pourcentagePlace * float64(nbARecruter)))
	candidatsFemmes := FiltreFemme(candidats) // permet d'isoler les femmes parmi les candidat.es
	r.logger.LogfType(LOG_RECRUTEMENT, "Le service RH veut recruter %.2f pourcents de femmes soit %d femmes", pourcentagePlace, nbFemmesARecruter)
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
	r.logger.LogfType(LOG_RECRUTEMENT, "Le service RH a pu recruter %d femmes sur %d femmes souhaitées", len(embauches), nbFemmesARecruter)
	// S'il n'y a pas assez de femmes dans les candidats pour toutes les places réservées, on recrute des hommes

	// Le reste des candidats sont sélectionnés uniquement pour leurs compétences
	reste := nbARecruter - len(embauches)
	for i := 0; i < reste; i++ {
		maxCandidats := EmployeMaxCompetences(candidats)
		idx := rand.Intn(len(maxCandidats))
		embauches = append(embauches, maxCandidats[idx])
		candidats = enleverEmploye(candidats, maxCandidats[idx])
	}

	return embauches
}

// Fonction de recrutement générale que l'entreprise peut appeler à chaque pas de temps
// Réalise un recrutement à partir des choix renseignés par l'utilisation lors de l'initialisation
func (r *Recrutement) Recruter(nbARecruter int) (embauches []*Employe) {
	if nbARecruter < 0 {
		return nil
	}

	// Génération des candidats

	candidats := r.GenererCandidats(constantes.NB_CANDIDATS_PAR_OFFRE * nbARecruter)
	if candidats == nil {
		return nil
	}
	if r.objectif == -1 {
		// L'utilisateur n'a pas défini un objectif de parité à atteindre

		// Choix de la fonction de recrutement à appeler
		if r.typeRecrutementAvant == Competences {
			embauches := r.RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
			return embauches
		} else if r.typeRecrutementAvant == PlacesReservees {
			if r.pourcentagePlacesAvant < 0 || r.pourcentagePlacesAvant > 1 {
				return nil
			}
			embauches := r.RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesAvant)
			return embauches
		} else {
			return nil
		}

	} else {
		// L'utilisateur a défini un pourcentage de parité a atteindre

		// Choix de la fonction de recrutement à appeler en fonction de si l'objectif est atteint ou non
		if r.objectif > r.entreprise.PourcentageFemmes() {
			// L'objectif n'est pas atteint, on applique TypeRecrutementAvant
			if r.typeRecrutementAvant == Competences {
				embauches := r.RecrutementCompetencesEgales(nbARecruter, r.stratAvant, candidats)
				return embauches
			} else if r.typeRecrutementAvant == PlacesReservees {
				if r.pourcentagePlacesAvant < 0 || r.pourcentagePlacesAvant > 1 {
					return nil
				}
				embauches := r.RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesAvant)
				return embauches
			} else {
				return nil
			}
		} else {
			// L'objectif est atteint, on applique TypeRecrutementApres
			if r.typeRecrutementApres == Competences {
				embauches := r.RecrutementCompetencesEgales(nbARecruter, r.stratApres, candidats)
				return embauches
			} else if r.typeRecrutementApres == PlacesReservees {
				if r.pourcentagePlacesApres < 0 || r.pourcentagePlacesApres > 1 {
					return nil
				}
				embauches := r.RecrutementPlacesReservees(nbARecruter, candidats, r.pourcentagePlacesApres)
				return embauches
			} else {
				return nil
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
	for !r.fin {
		// Attend un message pour agir
		msg := <-r.chnl
		switch msg.Act {
		case RECRUTEMENT:
			embauches := r.Recruter(msg.Payload.(int))
			for _, emp := range embauches {
				go func(emp *Employe) {
					emp.Start()
				}(emp)
			}

			r.chnl <- CommunicateurRecrutement{FIN_RECRUTEMENT, embauches}

		case FIN_AGENT:
			r.fin = true
		}
	}
	log.Printf("Fin du recrutement")
}
