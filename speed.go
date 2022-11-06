package mathunit

type Speed string

const (
	// MeterPerSecond is the SI unit of speed
	MeterPerSecond Speed = "m/s"
	// KilometerPerHour is the SI unit of speed
	KilometerPerHour Speed = "km/h"
	// CentimeterPerSecond is the SI unit of speed
	CentimeterPerSecond Speed = "cm/s"
	// MillimeterPerSecond is the SI unit of speed
	MillimeterPerSecond Speed = "mm/s"
	// MicrometerPerSecond is the SI unit of speed
	MicrometerPerSecond Speed = "µm/s"
	// NanometerPerSecond is the SI unit of speed
	NanometerPerSecond Speed = "nm/s"
	// AngstromPerSecond is the SI unit of speed
	AngstromPerSecond Speed = "Å/s"
	// InchPerSecond is the imperial unit of speed
	InchPerSecond Speed = "in/s"
	// FootPerSecond is the imperial unit of speed
	FootPerSecond Speed = "ft/s"
	// YardPerSecond is the imperial unit of speed
	YardPerSecond Speed = "yd/s"
	// MilePerHour is the imperial unit of speed
	MilePerHour Speed = "mi/h"
	// Knot is the imperial unit of speed
	Knot Speed = "kn"
)

func (s Speed) String() string {
	return string(s)
}

func (s Speed) UnitType() string {
	return "speed"
}

func (s Speed) Coefficient() float64 {
	switch s {
	case MeterPerSecond:
		return 1
	case KilometerPerHour:
		return 3.6
	case CentimeterPerSecond:
		return 0.01
	case MillimeterPerSecond:
		return 0.001
	case MicrometerPerSecond:
		return 0.000001
	case NanometerPerSecond:
		return 0.000000001
	case AngstromPerSecond:
		return 0.000000000001
	case InchPerSecond:
		return 0.0254
	case FootPerSecond:
		return 0.3048
	case YardPerSecond:
		return 0.9144
	case MilePerHour:
		return 1.609344
	case Knot:
		return 1.852
	}
	return 1
}

func (s Speed) Values() []string {
	return []string{
		string(MeterPerSecond),
		string(KilometerPerHour),
		string(CentimeterPerSecond),
		string(MillimeterPerSecond),
		string(MicrometerPerSecond),
		string(NanometerPerSecond),
		string(AngstromPerSecond),
		string(InchPerSecond),
		string(FootPerSecond),
		string(YardPerSecond),
		string(MilePerHour),
		string(Knot),
	}
}

func (s Speed) Valid() bool {
	for _, v := range s.Values() {
		if v == string(s) {
			return true
		}
	}
	return false
}

func (s Speed) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case Duration:
		return Meter, s.Coefficient() * b.Coefficient()
	}
	return nil, 0
}

func (s Speed) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case Speed:
		return Dimensionless, s.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
