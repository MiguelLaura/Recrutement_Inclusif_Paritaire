package agt

type SituationActuelle struct {
	NbEmp  int     `json:"nbEmp"`
	Parite float64 `json:"parite"`
	Benef  float64 `json:"benefices"`
}

type ReponseAuClient struct {
	Action string `json:"action"`
	Succes bool   `json:"succes"`
}
