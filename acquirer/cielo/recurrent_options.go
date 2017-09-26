package cielo

type Recurrent uint32

const (
	Monthly Recurrent = 1 + iota
	Bimonthly
	Quarterly
	SemiAnnual
	Annual
)

var recurrentOptions = [...]string{
	"Monthly",
	"Bimonthly",
	"Quarterly",
	"SemiAnnual",
	"Annual",
}

// Get the value for the recurrent option
func (r Recurrent) getOption() string {
	return recurrentOptions[r-1]
}
