package agt

import (
	"math/rand"

	"gitlab.utc.fr/mennynat/ia04-project/agt/constantes"
)

func GenererEmployeInit(ent Entreprise, genre Genre) *Employe {

	// Génération aléatoire de l'attribut agresseur avec une loi de Bernouilli
	// La probabilité d'être un agresseur.e dépend du genre de la personne
	var b Bernoulli
	// La loi de distribution b renvoie la valeur 1 avec une probabilité P, 0 sinon
	if genre == Homme {
		b.P = constantes.POURCENTAGE_AGRESSEUR_H
	} else {
		b.P = constantes.POURCENTAGE_AGRESSEUR_F
	}
	var agg bool // false par défaut
	if b.Rand() == 1 {
		agg = true
	}

	// Génération aléatoire de l'ancienneté de l'employé entre 0 et ANCIENNETE_MAX
	anc := rand.Intn(constantes.ANCIENNETE_MAX)

	// Génération aléatoire du comportement de l'employé
	// A FINIR APRES DISCUSSION
	// Loi Bernouilli si deux comportements possibles, binomiale si plus

	return NewEmploye(genre, anc, constantes.SANTE_MENTALE_MAX, agg, 0, ent)
}
