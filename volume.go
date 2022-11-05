package mathunit

type Volume string

const (
	// Liter is the SI unit of volume
	Liter Volume = "l"
	// Milliliter is the SI unit of volume
	Milliliter Volume = "ml"
	// CubicMeter is the SI unit of volume
	CubicMeter Volume = "m³"
	// CubicCentimeter is the SI unit of volume
	CubicCentimeter Volume = "cm³"
	// CubicMillimeter is the SI unit of volume
	CubicMillimeter Volume = "mm³"
	// CubicInch is the imperial unit of volume
	CubicInch Volume = "in³"
	// CubicFoot is the imperial unit of volume
	CubicFoot Volume = "ft³"
	// CubicYard is the imperial unit of volume
	CubicYard Volume = "yd³"
	// CubicMile is the imperial unit of volume
	CubicMile Volume = "mi³"
	// CubicNauticalMile is the imperial unit of volume
	CubicNauticalMile Volume = "nmi³"
	// AcreFoot is the imperial unit of volume
	AcreFoot Volume = "ac ft"
	// AcreInch is the imperial unit of volume
	AcreInch Volume = "ac in"
	// Barrel is the imperial unit of volume
	Barrel Volume = "bbl"
	// Bushel is the imperial unit of volume
	Bushel Volume = "bu"
	// Cord is the imperial unit of volume
	Cord Volume = "cd"
	// Cup is the imperial unit of volume
	Cup Volume = "cup"
	// FluidOunce is the imperial unit of volume
	FluidOunce Volume = "fl oz"
	// Gallon is the imperial unit of volume
	Gallon Volume = "gal"
	// Peck is the imperial unit of volume
	Peck Volume = "pk"
	// Pint is the imperial unit of volume
	Pint Volume = "pt"
	// Quart is the imperial unit of volume
	Quart Volume = "qt"
	// Tablespoon is the imperial unit of volume
	Tablespoon Volume = "tbsp"
	// Teaspoon is the imperial unit of volume
	Teaspoon Volume = "tsp"
)

func (v Volume) String() string {
	return string(v)
}

func (v Volume) Values() []string {
	return []string{
		"l",
		"ml",
		"m³",
		"cm³",
		"mm³",
		"in³",
		"ft³",
		"yd³",
		"mi³",
		"nmi³",
		"ac ft",
		"ac in",
		"bbl",
		"bu",
		"cd",
		"cup",
		"fl oz",
		"gal",
		"pk",
		"pt",
		"qt",
		"tbsp",
		"tsp",
	}
}

func (v Volume) Valid() bool {
	for _, value := range v.Values() {
		if value == v.String() {
			return true
		}
	}
	return false
}

func (v Volume) Coefficient() float64 {
	switch v {
	case Liter:
		return 1
	case Milliliter:
		return 0.001
	case CubicMeter:
		return 1000
	case CubicCentimeter:
		return 0.001
	case CubicMillimeter:
		return 0.000001
	case CubicInch:
		return 0.016387064
	case CubicFoot:
		return 28.316846592
	case CubicYard:
		return 764.554857984
	case CubicMile:
		return 4168181825.44058
	case CubicNauticalMile:
		return 6352182208
	case AcreFoot:
		return 1233484800
	case AcreInch:
		return 61024
	case Barrel:
		return 158.987294928
	case Bushel:
		return 35.23907016688
	case Cord:
		return 128
	case Cup:
		return 0.2365882365
	case FluidOunce:
		return 0.0295735295625
	case Gallon:
		return 3.785411784
	case Peck:
		return 8.809767541696
	case Pint:
		return 0.473176473
	case Quart:
		return 0.946352946
	case Tablespoon:
		return 0.01478676478125
	case Teaspoon:
		return 0.00492892159375
	}
	return 1
}

func (v Volume) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case PackageUnit:
		return v, b.Coefficient()
	}
	return nil, 0
}

func (v Volume) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case PackageUnit:
		return v, 1 / b.Coefficient()
	case Volume:
		return Bottle, v.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
