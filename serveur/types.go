package serveur

import "gitlab.utc.fr/mennynat/ia04-project/agt"

// -----------------------------
//       Types requêtes
// -----------------------------

type requeteNouvelleSimulation struct {
	ID                     string              `json:"id_simulation"`
	NbEmployes             int                 `json:"nb_employes"`
	PourcentageFemmes      float64             `json:"pourcentage_femmes"`
	Objectif               float64             `json:"objectif"`
	StratAvant             agt.StratParite     `json:"strat_avant"`
	StratApres             agt.StratParite     `json:"strat_apres"`
	TypeRecrutementAvant   agt.TypeRecrutement `json:"type_recrutement_avant"`
	TypeRecrutementApres   agt.TypeRecrutement `json:"type_recrutement_apres"`
	PourcentagePlacesAvant float64             `json:"pourcentage_places_avant"`
	PourcentagePlacesApres float64             `json:"pourcentage_places_apres"`
}

type reponseNouvelleSimulation struct {
	ID string `json:"id_simulation"`
}
