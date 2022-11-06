package mathunit

type Area string

const (
	// SquareMeter is the SI unit of area
	SquareMeter Area = "m²"
	// SquareKilometer is the SI unit of area
	SquareKilometer Area = "km²"
	// SquareCentimeter is the SI unit of area
	SquareCentimeter Area = "cm²"
	// SquareMillimeter is the SI unit of area
	SquareMillimeter Area = "mm²"
	// SquareMicrometer is the SI unit of area
	SquareMicrometer Area = "µm²"
	// SquareNanometer is the SI unit of area
	SquareNanometer Area = "nm²"
	// SquareAngstrom is the SI unit of area
	SquareAngstrom Area = "Å²"
	// SquareInch is the imperial unit of area
	SquareInch Area = "in²"
	// SquareFoot is the imperial unit of area
	SquareFoot Area = "ft²"
	// SquareYard is the imperial unit of area
	SquareYard Area = "yd²"
	// SquareMile is the imperial unit of area
	SquareMile Area = "mi²"
	// SquareNauticalMile is the imperial unit of area
	SquareNauticalMile Area = "nmi²"
	// Acre is the imperial unit of area
	Acre Area = "ac"
	// Hectare is the imperial unit of area
	Hectare Area = "ha"
)

func (a Area) String() string {
	return string(a)
}

func (a Area) UnitType() string {
	return "area"
}

func (a Area) Coefficient() float64 {
	switch a {
	case SquareMeter:
		return 1
	case SquareKilometer:
		return 1000000
	case SquareCentimeter:
		return 0.0001
	case SquareMillimeter:
		return 0.000001
	case SquareMicrometer:
		return 0.000000000001
	case SquareNanometer:
		return 0.000000000000001
	case SquareAngstrom:
		return 0.0000000000000000001
	case SquareInch:
		return 0.00064516
	case SquareFoot:
		return 0.09290304
	case SquareYard:
		return 0.83612736
	case SquareMile:
		return 2589988.110336
	case SquareNauticalMile:
		return 3429904
	case Acre:
		return 4046.8564224
	case Hectare:
		return 10000
	}
	return 1
}

func (a Area) Values() []string {
	return []string{
		"m²",
		"km²",
		"cm²",
		"mm²",
		"µm²",
		"nm²",
		"Å²",
		"in²",
		"ft²",
		"yd²",
		"mi²",
		"nmi²",
		"ac",
		"ha",
	}
}

func (a Area) Valid() bool {
	for _, v := range a.Values() {
		if v == string(a) {
			return true
		}
	}
	return false
}

func (a Area) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case Length:
		return CubicMeter, a.Coefficient() * b.Coefficient()
	}
	return nil, 0
}

func (a Area) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case Area:
		return Dimensionless, a.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
