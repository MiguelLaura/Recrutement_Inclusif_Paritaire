package agt

type SituationActuelle struct {
	Annee  int     `json:"annee"`
	NbEmp  int     `json:"nbEmp"`
	Parite float64 `json:"parite"`
	Benef  float64 `json:"benefices"`
}
