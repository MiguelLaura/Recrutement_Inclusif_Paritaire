package agt

type SituationActuelle struct {
	Annee        int     `json:"annee"`
	NbEmp        int     `json:"nbEmp"`
	Parite       float64 `json:"parite"`
	Benef        int     `json:"benefices"`
	Competence   float64 `json:"competence"`
	SanteMentale float64 `json:"sante-mentale"`
}

func NewSituationActuelle(annee int, nbemp int, parite float64, benef int, competence float64, santeMentale float64) *SituationActuelle {
	return &SituationActuelle{
		Annee:        annee,
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

type InformationsInitiales struct {
	AnneesInit           int     `json:"anneesInit"`
	PariteInit           float64 `json:"pariteInit"`
	NbEmployesInit       int     `json:"nbEmployesInit"`
	Status               string  `json:"status"`
	ObjectifRecrut       float64 `json:"objectif"`
	InfoRecrutementAvant string  `json:"recrutAvant"`
	InfoRecrutementApres string  `json:"recrutApres"`
}
