package mathunit

type Dimension string

const (
	Dimensionless Dimension = ""
)

func (d Dimension) String() string {
	return string(d)
}

func (d Dimension) Coefficient() float64 {
	return 1
}

func (d Dimension) Values() []string {
	return []string{string(Dimensionless)}
}

func (d Dimension) Valid() bool {
	switch d {
	case Dimensionless:
		return true
	}
	return false
}

func (d Dimension) Mul(b Unit) (Unit, float64) {
	return b, 1
}

func (d Dimension) Div(b Unit) (Unit, float64) {
	return nil, 0
}
