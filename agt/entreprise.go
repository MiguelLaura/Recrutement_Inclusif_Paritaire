package agt

func (ent Entreprise) nbEmployes() int {
	return len(ent.employes)
}

func (ent Entreprise) PourcentageFemmes() float64 {
	femmes := FiltreFemme(ent.employes)
	return float64(len(femmes)) / float64(len(ent.employes))
}
