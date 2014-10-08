package cpe

type Part rune

var (
	Application      = Part('a')
	OperationgSystem = Part('o')
	Hardware         = Part('h')
)

func (m Part) String() string {
	if m.IsValid() {
		return string(m)
	} else {
		panic("\"%v\" is not valid as part attribute.")
	}
}

func (m Part) IsValid() bool {
	switch m {
	case Application, OperationgSystem, Hardware:
		return true
	default:
		return false
	}
}
