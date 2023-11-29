package agt

import (
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

func GenererEmployeInit(ent *Entreprise, genre Genre) *Employe {
	// Génération aléatoire de l'attribut agresseur
	agg := genAgresseur(genre)

	// Génération aléatoire de l'ancienneté de l'employé entre 0 et ANCIENNETE_MAX
	anc := rand.Intn(constantes.ANCIENNETE_MAX)

	// Génération aléatoire du comportement de l'employé
	// On considère une proba égale d'avoir les différents comportements
	compor := genComportement()

	// Génération aléatoire de la compétence de l'employé
	// A FAIRE
	// Piste: loi normale avec mu=50 et sd=10 ? (voir premier lien note Laura)

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, compor, 0, ent)
}
