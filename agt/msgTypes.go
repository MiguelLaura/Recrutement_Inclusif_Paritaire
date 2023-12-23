package agt

type SituationActuelle struct {
	NbEmp  int     `json:"nbEmp"`
	Parite float64 `json:"parite"`
	Benef  int     `json:"benefices"`
}

func NewSituationActuelle(nbemp int, parite float64, benef int) *SituationActuelle {
	return &SituationActuelle{
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
	Status               string  `json:"status"`
	ObjectifRecrut       float64 `json:"objectif"`
	InfoRecrutementAvant string  `json:"recrutAvant"`
	InfoRecrutementApres string  `json:"recrutApres"`
}
