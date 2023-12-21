package agt

import (
	"fmt"
	"log"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
	"gitlab.utc.fr/mennynat/ia04-project/utils/logger"
)

// ---------------------
//        Général
// ---------------------

// Permet d'envoyer un certain message à un Employe. Ce message contient une action qu'il va effectuer
// ainsi qu'un payload optionnel permettant de transmettre des informations en plus à l'agent.
func EnvoyerMessage(dest *Employe, act Action, payload any) {
	dest.chnl <- Communicateur{act, payload}
}

// ---------------------
//     Constructeurs
// ---------------------

func genererIDEmploye() EmployeID {
	res := fmt.Sprintf("Employe%d", agtCnt)
	agtCnt++

	return EmployeID(res)
}

func GenererEmployeInit(ent **Entreprise, genre Genre, logger *logger.Loggers) *Employe {
	// Génération aléatoire de l'attribut agresseur
	agg := genAgresseur(genre)

	// Génération aléatoire de l'ancienneté de l'employé entre 0 et ANCIENNETE_MAX
	anc := rand.Intn(constantes.ANCIENNETE_MAX)

	// Génération aléatoire de la compétence de l'employé
	competence := genCompetence()

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, competence, *ent, logger)
}

func NewEmploye(gen Genre, anc int, san int, ag bool, compet int, ent *Entreprise, logger *logger.Loggers) *Employe {
	return &Employe{
		id:              genererIDEmploye(),
		genre:           gen,
		anciennete:      anc,
		santeMentale:    san,
		agresseur:       ag,
		competence:      compet,
		cmpt_competence: 0,
		entreprise:      ent,
		fin:             false,
		chnl:            make(chan Communicateur),
		logger:          logger,
	}
}

// ---------------------
//        Getters
// ---------------------

func (e *Employe) Id() EmployeID {
	return e.id
}

func (e *Employe) Genre() Genre {
	return e.genre
}

func (e *Employe) Anciennete() int {
	return e.anciennete
}

func (e *Employe) SanteMentale() int {
	return e.santeMentale
}

func (e *Employe) Agresseur() bool {
	return e.agresseur
}

func (e *Employe) Competence() int {
	return e.competence
}

func (e *Employe) Cmpt_competence() int {
	return e.cmpt_competence
}

func (e *Employe) Entreprise() *Entreprise {
	return e.entreprise
}

func (e *Employe) String() string {
	return fmt.Sprintf("%s (%s)", e.id, StringGenre(e.genre))
}

// ---------------------
//     Comportement
// ---------------------

// L'Employe a passé une nouvelle année dans l'entreprise
func (e *Employe) gagnerAnciennete() {
	e.anciennete += 1
}

// L'Employé porte plainte à son entreprise au sujet d'un autre employé.
func (plaignant *Employe) porterPlainte(accuse *Employe) {
	plaignant.entreprise.RecevoirPlainte(plaignant, accuse)
}

// L'Employé pose sa démission auprès de son entreprise
func (e *Employe) poserDemission() {
	e.entreprise.RecevoirDemission(e)
}

// L'Employé pose sa démission auprès de son entreprise après son congé maternité
func (e *Employe) poserDemissionMaternite() {
	e.entreprise.RecevoirDemissionMaternite(e)
}

// L'Employé pose sa démission auprès de son entreprise pour cause de dépression
func (e *Employe) partirDepression() {
	e.entreprise.RecevoirDepression(e)
}

// L'Employé arrive à la fin de sa carrière et pose donc sa retraite
func (e *Employe) partirRetraite() {
	e.entreprise.RecevoirRetraite(e)
}

// Peut-être à nuancer si trop de gains de compétences
func (e *Employe) seFormer() {
	e.logger.LogfType(LOG_EVENEMENT, "%s a participé à une formation", e.String())
	e.cmpt_competence += 1
	if e.competence < 10 && e.cmpt_competence == 5 {
		e.logger.LogfType(LOG_EMPLOYE, "%s a amélioré ses compétences", e.String())
		e.competence += 1
		e.cmpt_competence = 0
	}
}

func (e *Employe) avoirEnfant() {
	e.logger.LogfType(LOG_EMPLOYE, "%s a un enfant", e.String())
	if e.Genre() == Femme {
		if rand.Float64() < constantes.PROBA_CONGE_F {
			e.entreprise.CongeParental(e)
		}
	} else {
		if rand.Float64() < constantes.PROBA_CONGE_H {
			e.entreprise.CongeParental(e)
		}
	}
}

// ---------------------
//      Evenements
// ---------------------

// L'Employé est agressé par quelqu'un
func (agresse *Employe) etreAgresse(agresseur *Employe) {

	agresse.logger.LogfType(LOG_AGRESSION, "%s agresse %s", agresseur.String(), agresse.String())

	// Selon son comportement, il va porter plainte ou non
	if rand.Float64() < constantes.PROBA_PLAINTE {
		agresse.porterPlainte(agresseur)
	}

	agresse.santeMentale -= constantes.DEGATS_TRAUMATISME
}

// ---------------------
//  Actions sur autres
// ---------------------

// L'agent agresse quelqu'un pris au hasard dans son entreprise
func (agresseur *Employe) agresser() {
	var genreAgresse Genre
	if agresseur.genre == Femme {
		genreAgresse = Homme
	} else {
		genreAgresse = Femme
	}

	timeout := 0

	cible := agresseur.entreprise.EnvoyerEmploye(genreAgresse)

	// S'assure de ne pas s'agresser lui-même
	for (cible == nil || cible.id == agresseur.id) && timeout < constantes.TIMEOUT_AGRESSION {
		cible = agresseur.entreprise.EnvoyerEmploye(genreAgresse)
		timeout++
	}
	// agresseur.logger.LogfType(LOG_AGRESSION, "Employé %s agresse %s", agresseur.id, cible.id)

	if timeout < constantes.TIMEOUT_AGRESSION {
		go EnvoyerMessage(cible, AGRESSION, agresseur)
	} else {
		// Il a trouvé personne à agresser
		go EnvoyerMessage(cible, AGRESSION, nil)
	}
}

// ---------------------
//  Logique de simulation
// ---------------------

// Lance la vie de l'agent
func (e *Employe) Start() {
	log.Printf("Démarrage de l'employé %s", e.id)

	// Initialisation

	// Boucle de vie
	for !e.fin {
		e.agir()
	}
	log.Printf("Fin de l'employé %s", e.Id())
}

// Ce que l'employé fait à chaque tour
func (e *Employe) agir() {

	// Attend un message pour agir
	msg := <-e.chnl

	switch msg.Act {
	case NOOP: // Ne fait rien
		return
	case LIBRE: // Vie une année complète
		// e.logger.LogfType(LOG_EMPLOYE, "action libre %s", e.id)

		// Si l'agent est un agresseur, il agresse
		if e.Agresseur() {
			e.agresser()
		}

		// Participer à une formation
		i, _ := TrouverEmploye(*e.entreprise.formation, func(emp Employe) bool { return e.Id() == emp.Id() }, 0)
		if i >= 0 {
			e.seFormer()
		}

		// Vieillir
		e.gagnerAnciennete()

		// Avoir un enfant
		enfant := false
		if rand.Float64() < constantes.PROBA_ENFANT {
			e.avoirEnfant()
			enfant = true
		}

		// Demissionner apres congé maternité
		if e.Genre() == Femme && enfant {
			if rand.Float64() <= constantes.PROBA_DEPART_F {
				e.poserDemissionMaternite()
			}
		}

		// Depart à la retraite
		if e.anciennete >= constantes.ANCIENNETE_MAX {
			e.partirRetraite()
		}

		// Or, au sein d'un organisation inclusive, les employés ont une rétention supérieure de 20% (source: Catalyst)

		if e.entreprise.PourcentageFemmes() < constantes.SEUIL_IMPACT_FEMME {
			if rand.Float64() <= constantes.POURCENTAGE_DEM_SPONTANEE {
				e.poserDemission()
			}
		} else {
			if rand.Float64() <= constantes.POURCENTAGE_DEM_SPONTANEE*(1-constantes.RETENTION_PARITE) {
				e.poserDemission()
			}
		}

	case AGRESSION: // Se fait agresser par quelqu'un
		// e.logger.LogfType(LOG_EMPLOYE, "action agression %s", e.id)

		if msg.Payload != nil {
			e.etreAgresse(msg.Payload.(*Employe))
		} else {
			log.Printf("%s : je vais pas m'agresser moi-même", e.id)
		}

		// Si l'agent n'a plus de santé mentale, il pose sa démission
		if e.santeMentale <= 0 {
			e.partirDepression()
		}

	case FIN: // Arrêter l'employé
		// e.logger.LogfType(LOG_EMPLOYE, "action fin %s", e.id)
		e.fin = true
	}

	// Permet de notifier l'entreprise que l'agent vient de faire une action
	EnvoyerNotifActions(e.entreprise, LIBRE, nil)
}
