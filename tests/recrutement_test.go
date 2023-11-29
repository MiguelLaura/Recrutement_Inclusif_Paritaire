package tests

import (
	"fmt"
	"testing"

	"gitlab.utc.fr/mennynat/ia04-project/agt"
)

func Test_EmployeMaxCompetences(t *testing.T) {
	var ent agt.Entreprise
	var e1 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 6, &ent)
	var e2 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 4, &ent)
	var e3 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 8, &ent)
	var e4 agt.Employe = *agt.NewEmploye(1, 0, 100, false, 0.25, 1, &ent)
	var cand []agt.Employe
	cand = append(cand, e1)
	cand = append(cand, e2)
	cand = append(cand, e3)
	cand = append(cand, e4)
	fmt.Println("candidats: ", cand)
	embauches := agt.EmployeMaxCompetences(cand)
	if embauches[0].id != "employe1"
}
