package agt

import (
	"fmt"
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

type EmployeID string

var agtCnt int = 0

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
	// A FAIRE
	// Piste: loi normale avec mu=50 et sd=10 ? (voir premier lien note Laura)

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, compor, 0, ent)
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
//      Evenements
// ---------------------

func (e *Employe) Communiquer(act Action, payload any) {
	go func() {
		e.chnl <- Communicateur{act, payload}
	}()
}

func (e *Employe) gagnerAnciennete() {
	e.anciennete += 1
}

// ---------------------
//  Logique de simulation
// ---------------------

func (e *Employe) Start() {
	go func() {

		// Initialisation

		// Boucle de vie
		for {
			e.agir()
		}

	}()
}

func (e *Employe) agir() {

	// Attend un message pour agir
	msg := <-e.chnl

	switch msg.Act {
	case NOOP:
		return
	}

}
