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
