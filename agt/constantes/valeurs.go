package constantes

type ObjectifPourcentageFemmes float64
type ComportementVSS float64
type Interet string

const (
	Macho       ObjectifPourcentageFemmes = .0
	Indifferent ObjectifPourcentageFemmes = .2 // 20% de femmes dans le monde du travail (ref. Mathilde)
	RespectLoi  ObjectifPourcentageFemmes = .4 // C'est la loi !! (Dura lex sed lex)
	GirlPower   ObjectifPourcentageFemmes = .75
)

const (
	Inactif ComportementVSS = .0
)

type Valeurs struct {
}
