package mathunit

type Angle string

const (
	// Radian is the SI unit of angle
	Radian Angle = "rad"
	// Degree is the SI unit of angle
	Degree Angle = "deg"
	// Grad is the SI unit of angle
	Grad Angle = "grad"
	// ArcMinute is the SI unit of angle
	ArcMinute Angle = "arcmin"
	// ArcSecond is the SI unit of angle
	ArcSecond Angle = "arcsec"
	// Revolution is the SI unit of angle
	Revolution Angle = "rev"
)

func (a Angle) String() string {
	return string(a)
}

func (a Angle) UnitType() string {
	return "angle"
}

func (a Angle) Coefficient() float64 {
	switch a {
	case Radian:
		return 1
	case Degree:
		return 0.017453292519943295
	case Grad:
		return 0.015707963267948966
	case ArcMinute:
		return 0.0002908882086657216
	case ArcSecond:
		return 4.84813681109536e-06
	case Revolution:
		return 6.283185307179586
	}
	return 1
}

func (a Angle) Values() []string {
	return []string{
		"rad",
		"deg",
		"grad",
		"arcmin",
		"arcsec",
		"rev",
	}
}

func (a Angle) Valid() bool {
	switch a {
	case Radian, Degree, Grad, ArcMinute, ArcSecond, Revolution:
		return true
	}
	return false
}

func (a Angle) Mul(b Unit) (Unit, float64) {
	return nil, 0
}

func (a Angle) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case Angle:
		return Dimensionless, a.Coefficient() / b.Coefficient()
	case Duration:
		return RadianPerSecond, a.Coefficient() / b.Coefficient()
	case Frequency:
		return Radian, a.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
