package agt

import (
	"fmt"
	"math"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
	"gonum.org/v1/gonum/stat/distuv"
)

type EmployeID string

var agtCnt int = 0

// On veut que les compétences tournent autour de 5 sans trop s’éparpiller autour
var loiNormale = distuv.Normal{
	Mu:    5,
	Sigma: 3,
}

type Employe struct {
	id           EmployeID
	genre        Genre
	anciennete   int //entre 0 et 40
	santeMentale int //entre 0 et 100
	agresseur    bool
	comportement Comportement
	competence   int //entre 0 et 10
	entreprise   *Entreprise
	chnl         chan Communicateur
}

// ---------------------
//        Général
// ---------------------

// Permet d'envoyer un certain message à un Employe. Ce message contient une action qu'il va effectuer
// ainsi qu'un payload optionnel permettant de transmettre des informations en plus à l'agent.
func EnvoyerMessage(dest *Employe, act Action, payload any) {
	go func() {
		dest.chnl <- Communicateur{act, payload}
	}()
}

// ---------------------
//     Constructeurs
// ---------------------

func genererIDEmploye() EmployeID {
	res := fmt.Sprintf("employe%d", agtCnt)
	agtCnt++

	return EmployeID(res)
}

func GenererEmployeInit(ent *Entreprise, genre Genre) *Employe {

	var agg bool // false par défaut
	// Génération aléatoire de l'attribut agresseur
	if genre == Homme {
		if rand.Float64() <= constantes.POURCENTAGE_AGRESSEUR_H {
			agg = true
		}
	} else {
		if rand.Float64() <= constantes.POURCENTAGE_AGRESSEUR_H {
			agg = true
		}
	}

	// Génération aléatoire de l'ancienneté de l'employé entre 0 et ANCIENNETE_MAX
	anc := rand.Intn(constantes.ANCIENNETE_MAX)

	// Génération aléatoire du comportement de l'employé
	// On considère une proba égale d'avoir les différents comportements
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

	// Génération aléatoire de la compétence de l'employé

	// Permet de ne pas avoir de compétence négative et de ne pas aller au dessus du seuil max de compétence
	competence := int(math.Abs(loiNormale.Rand())) % (constantes.COMPETENCE_MAX + 1)

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, compor, competence, ent)
}

func NewEmploye(gen Genre, anc int, san int, ag bool, compor Comportement, compet int, ent *Entreprise) *Employe {
	return &Employe{
		id:           genererIDEmploye(),
		genre:        gen,
		anciennete:   anc,
		santeMentale: san,
		agresseur:    ag,
		comportement: compor,
		competence:   compet,
		entreprise:   ent,
		chnl:         make(chan Communicateur),
	}
}

// ---------------------
//        Getters
// ---------------------

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

// L'employé porte plainte à son entreprise au sujet d'un autre employé.
func (plaignant *Employe) porterPlainte(accuse *Employe) {
	plaignant.entreprise.RecevoirPlante(plaignant, accuse)
}

// L'Employé pose sa démission auprès de son entreprise
func (e *Employe) poserDemission() {
	e.entreprise.RecevoirDemission(e)
}

// L'Employé arrive à la fin de sa carrière et pose donc sa retraite
func (e *Employe) partirRetraite() {
	e.entreprise.RecevoirDemission(e)
}

// L'Employé travaille sur cette année
func (e *Employe) travailler() {
	e.entreprise.MettreAJourCA(e.santeMentale, e.competence)
}

// ---------------------
//      Evenements
// ---------------------

// L'employé est agressé par quelqu'un
func (agresse *Employe) etreAgresse(agresseur *Employe) {

	// Selon son comportement, il va porter plainte ou non
	if rand.Float32() < float32(agresse.comportement) {
		agresse.porterPlainte(agresseur)
	}

	agresse.santeMentale -= constantes.DEGATS_TRAUMATISME
}

// ---------------------
//  Actions sur autres
// ---------------------

// L'agent agresse quelqu'un pris au hasard dans son entreprise
func (agresseur *Employe) agresser() {
	cible := agresseur.entreprise.EnvoyerEmploye()

	// S'assure de ne pas s'agresser lui-même
	for cible == agresseur {
		cible = agresseur.entreprise.EnvoyerEmploye()
	}

	EnvoyerMessage(cible, AGRESSION, agresseur)
}

// ---------------------
//  Logique de simulation
// ---------------------

// Lance la vie de l'agent
func (e *Employe) Start() {
	go func() {

		// Initialisation

		// Boucle de vie
		for {
			e.agir()
		}

	}()
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

		e.gagnerAnciennete()

		if e.anciennete >= constantes.ANCIENNETE_MAX {
			e.partirRetraite()
		}

	case AGRESSION: // Se fait agressé par quelqu'un

		e.etreAgresse(msg.Payload.(*Employe))

		// Si l'agent n'a plus de santé mentale, il pose sa démission
		if e.santeMentale <= 0 {
			e.poserDemission()
		}

	}

	// Permet de notifier l'entreprise que l'agent vient de faire une action
	e.entreprise.NotifierAction()
}
