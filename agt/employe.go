package agt

import (
	"fmt"
	"log"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
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
	res := fmt.Sprintf("employe%d", agtCnt)
	agtCnt++

	return EmployeID(res)
}

func GenererEmployeInit(ent **Entreprise, genre Genre) *Employe {
	// Génération aléatoire de l'attribut agresseur
	agg := genAgresseur(genre)

	// Génération aléatoire de l'ancienneté de l'employé entre 0 et ANCIENNETE_MAX
	anc := rand.Intn(constantes.ANCIENNETE_MAX)

	// Génération aléatoire du comportement de l'employé
	// On considère une proba égale d'avoir les différents comportements
	compor := genComportement()

	// Génération aléatoire de la compétence de l'employé
	competence := genCompetence()

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, compor, competence, *ent)
}

func NewEmploye(gen Genre, anc int, san int, ag bool, compor Comportement, compet int, ent *Entreprise) *Employe {
	return &Employe{
		id:              genererIDEmploye(),
		genre:           gen,
		anciennete:      anc,
		santeMentale:    san,
		agresseur:       ag,
		comportement:    compor,
		competence:      compet,
		cmpt_competence: 0,
		entreprise:      ent,
		chnl:            make(chan Communicateur),
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

func (e *Employe) Comportement() Comportement {
	return e.comportement
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
	return fmt.Sprintf("%s (%d)", e.id, e.genre)
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

// L'Employé pose sa démission auprès de son entreprise pour cause de dépression
func (e *Employe) partirDepression() {
	e.entreprise.RecevoirDepression(e)
}

// L'Employé arrive à la fin de sa carrière et pose donc sa retraite
func (e *Employe) partirRetraite() {
	e.entreprise.RecevoirRetraite(e)
}

// L'Employé travaille sur cette année
func (e *Employe) travailler() {
	e.entreprise.MettreAJourCA(e.santeMentale, e.competence)
}

// Peut-être à nuancer si trop de gains de compétences
func (e *Employe) seFormer() {
	e.cmpt_competence += 1
	if e.competence < 10 && e.cmpt_competence == 5 {
		log.Printf("Formation %s", e.Id())
		e.competence += 1
		e.cmpt_competence = 0
	}
	//log.Printf("Apres formation : %d", e.competence)
}

// ---------------------
//      Evenements
// ---------------------

// L'Employé est agressé par quelqu'un
func (agresse *Employe) etreAgresse(agresseur *Employe) {

	log.Printf("Employé %s agresse %s", agresseur.id, agresse.id)

	// Selon son comportement, il va porter plainte ou non
	if rand.Float64() < float64(agresse.comportement) {
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
	for {
		e.agir()
	}
}

// Ce que l'employé fait à chaque tour
func (e *Employe) agir() {

	// Attend un message pour agir
	msg := <-e.chnl

	switch msg.Act {
	case NOOP: // Ne fait rien
		return
	case LIBRE: // Vie une année complète

		// Si l'agent est un agresseur, il agresse
		if e.Agresseur() {
			e.agresser()
		}

		e.travailler()

		// Participer à une formation
		i, _ := TrouverEmploye(*e.entreprise.formation, func(emp Employe) bool { return e.Id() == emp.Id() }, 0)
		if i >= 0 {
			e.seFormer()
		}

		// Vieillir
		e.gagnerAnciennete()

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

		if msg.Payload != nil {
			e.etreAgresse(msg.Payload.(*Employe))
		} else {
			log.Printf("%s : je vais pas m'agresser moi-même", e.id)
		}

		// Si l'agent n'a plus de santé mentale, il pose sa démission
		if e.santeMentale <= 0 {
			e.partirDepression()
		}

	}

	// Permet de notifier l'entreprise que l'agent vient de faire une action
	EnvoyerNotifActions(e.entreprise, LIBRE, nil)
}
