package agt

import (
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

func GenererEmployeInit(ent Entreprise, genre Genre) *Employe {

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
