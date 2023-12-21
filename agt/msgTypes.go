package agt

type SituationActuelle struct {
	NbEmp  int     `json:"nbEmp"`
	Parite float64 `json:"parite"`
	Benef  float64 `json:"benefices"`
}

func NewSituationActuelle(nbemp int, parite float64, benef float64) *SituationActuelle {
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
