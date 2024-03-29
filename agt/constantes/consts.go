package constantes

const (
	// GENERAL
	GEN_F_H_CANDIDATS       = .5  // ...
	ANCIENNETE_MAX          = 43  // ...
	SANTE_MENTALE_MAX       = 100 // ...
	COMPETENCE_MIN          = 0
	COMPETENCE_MAX          = 10
	POURCENTAGE_AGRESSEUR_F = .02  // ...
	POURCENTAGE_AGRESSEUR_H = .12  // ...
	SEUIL_IMPACT_FEMME      = 0.35 // D'après Rosabeth Kanter, une minorité doit représenter au minimum 35% du groupe dominant pour avoir une influence

	// AGRESSION
	DEGATS_TRAUMATISME = 20 // ...
	TIMEOUT_AGRESSION  = 4
	PROBA_PLAINTE      = .3
	PROBA_LICENCIEMENT = .175

	// RECRUTEMENT
	POURCENTAGE_RECRUTEMENT = .05
	NB_CANDIDATS_PAR_OFFRE  = 18 // Source : https://blog.flatchr.io/barometre-des-entreprises-qui-recrutent-deuxieme-semestre-2022

	// DEPARTS
	POURCENTAGE_DEM_SPONTANEE = 0.027 // Taux de démission spontanée en France au 1er trimestre 2022
	RETENTION_PARITE          = 0.2   // Au sein d'un organisation inclusive, les employés ont une rétention supérieure de 20% (source: Catalyst)

	// BENEFICES
	CA_PAR_EMPLOYE_MIC       = 225000 // Source : https://letowski.fr/quelques-grands-indicateurs-par-tranche-de-taille-dentreprise/
	CA_PAR_EMPLOYE_PME       = 208000
	CA_PAR_EMPLOYE_ETI       = 338000
	COUT_EMPLOYE             = 42927 // Source : https://mon-entreprise.urssaf.fr/simulateurs/salaire-brut-net
	COUT_RECRUTEMENT         = 6500
	BOOST_PRODUCTIVITE_FEMME = 0.2849 // Source : https://www.cairn.info/revue-travail-genre-et-societes-2010-1-page-181.htm
	SEUIL_AMENDE             = 0.4
	POURCENTAGE_AMENDE       = 0.01

	// CONGE PARENTAL
	PROBA_ENFANT       = .04575   // Indice de natalité français / Années de fertilité femme
	PROBA_CONGE_F      = 1        // Probabilité de partir en congé parental pour les femmes
	PROBA_CONGE_H      = .71      // Probabilité de partir en congé parental pour les femmes
	PROPORTION_ARRET_F = 4. / 12. // Pour le calcul du benefice, proportion du congé maternité sur l'année = 4 mois / 12 mois
	PROPORTION_ARRET_H = 1. / 12. // Pour le calcul du benefice, proportion du congé paternité sur l'année = 1 mois / 12 mois
	PROBA_DEPART_F     = .3       // % de femmes à quitter leur emploi après grossesse

	// FORMATION
	POURCENTAGE_FORMATION  = .32 // Pourcentage d'employés réalisant une formation par an
	PRIX_FORMATION         = 500 // Coût de la formation par jour et par employé
	NB_JOURS_FORMATION     = 7
	CONTRIBUTION_FORMATION = .45 // Entreprise participe à 45% des coûts de formation
	COUNT_FORMATION        = 5   // Nombre de formations pour monter en compétences

	// TEAM BUILDING
	BOOST_TEAM_BUILDING = 1
	COUT_TB_PAR_EMPLOYE = 100
)
