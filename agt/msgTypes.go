package agt

type SituationActuelle struct {
	NbEmp        int     `json:"nbEmp"`
	Parite       float64 `json:"parite"`
	Benef        int     `json:"benefices"`
	Competence   float64 `json:"competence"`
	SanteMentale float64 `json:"sante-mentale"`
}

func NewSituationActuelle(nbemp int, parite float64, benef int, competence float64, santeMentale float64) *SituationActuelle {
	return &SituationActuelle{
		NbEmp:        nbemp,
		Parite:       parite,
		Benef:        benef,
		Competence:   competence,
		SanteMentale: santeMentale,
	}
}

type ReponseAuClient struct {
	Action string `json:"action"`
	Succes bool   `json:"succes"`
}
