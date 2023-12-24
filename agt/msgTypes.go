package agt

type SituationActuelle struct {
	Annee  int     `json:"annee"`
	NbEmp  int     `json:"nbEmp"`
	Parite float64 `json:"parite"`
	Benef  int     `json:"benefices"`
}

func NewSituationActuelle(annee int, nbemp int, parite float64, benef int) *SituationActuelle {
	return &SituationActuelle{
		Annee:  annee,
		NbEmp:  nbemp,
		Parite: parite,
		Benef:  benef,
	}
}

type ReponseAuClient struct {
	Action string `json:"action"`
	Succes bool   `json:"succes"`
}

type InformationsInitiales struct {
	PariteInit           float64 `json:"pariteInit"`
	NbEmployesInit       int     `json:"nbEmployesInit"`
	Status               string  `json:"status"`
	ObjectifRecrut       float64 `json:"objectif"`
	InfoRecrutementAvant string  `json:"recrutAvant"`
	InfoRecrutementApres string  `json:"recrutApres"`
}
