package mathunit

type PackageUnit string

const (
	// SingleUnit
	Dozen       PackageUnit = "doz"
	Doubledozen PackageUnit = "ddoz"
	Halfdozen   PackageUnit = "hdoz"
	Threedozens PackageUnit = "3doz"
	Fourdozens  PackageUnit = "4doz"
	Gross       PackageUnit = "grs"
	Greatgross  PackageUnit = "ggrs"

	// MultipleUnits
	Box    PackageUnit = "box"
	Bag    PackageUnit = "bag"
	Bottle PackageUnit = "btl"
	Can    PackageUnit = "can"
)

func (p PackageUnit) String() string {
	return string(p)
}

func (p PackageUnit) Values() []string {
	return []string{
		"doz",
		"ddoz",
		"hdoz",
		"3doz",
		"4doz",
		"grs",
		"ggrs",
		"box",
		"bag",
		"btl",
		"can",
	}
}

func (p PackageUnit) Valid() bool {
	switch p {
	case Dozen, Doubledozen, Halfdozen, Threedozens, Fourdozens, Gross, Greatgross, Box, Bag, Bottle, Can:
		return true
	}
	return false
}

func (p PackageUnit) Coefficient() float64 {
	switch p {
	case Dozen:
		return 12
	case Doubledozen:
		return 24
	case Halfdozen:
		return 6
	case Threedozens:
		return 36
	case Fourdozens:
		return 48
	case Gross:
		return 144
	case Greatgross:
		return 1728
	case Box:
		return 1
	case Bag:
		return 1
	case Bottle:
		return 1
	case Can:
		return 1
	}
	return 1
}

func (p PackageUnit) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case Mass, Volume:
		return b, p.Coefficient()
	}
	return nil, 0
}

func (p PackageUnit) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case PackageUnit:
		return Dimensionless, p.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
