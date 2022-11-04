package mathunit

import "fmt"

type Length string

const (
	// Meter is the SI unit of length
	Meter Length = "m"
	// Kilometer is the SI unit of length
	Kilometer Length = "km"
	// Centimeter is the SI unit of length
	Centimeter Length = "cm"
	// Millimeter is the SI unit of length
	Millimeter Length = "mm"
	// Micrometer is the SI unit of length
	Micrometer Length = "µm"
	// Nanometer is the SI unit of length
	Nanometer Length = "nm"
	// Angstrom is the SI unit of length
	Angstrom Length = "Å"
	// Inch is the imperial unit of length
	Inch Length = "in"
	// Foot is the imperial unit of length
	Foot Length = "ft"
	// Yard is the imperial unit of length
	Yard Length = "yd"
	// Mile is the imperial unit of length
	Mile Length = "mi"
	// NauticalMile is the imperial unit of length
	NauticalMile Length = "nmi"
	// AstronomicalUnit is the astronomical unit of length
	AstronomicalUnit Length = "au"
	// LightYear is the astronomical unit of length
	LightYear Length = "ly"
	// Parsec is the astronomical unit of length
	Parsec Length = "pc"

	// The following units are not SI or imperial units, but are included for completeness

)

func (l Length) String() string {
	return string(l)
}

func (l Length) Values() []string {
	return []string{
		"m",
		"km",
		"cm",
		"mm",
		"µm",
		"nm",
		"Å",
		"in",
		"ft",
		"yd",
		"mi",
		"nmi",
		"au",
		"ly",
		"pc",
	}
}

func (l Length) Valid() bool {
	_, err := ParseLength(l.String())
	return err == nil
}

func (l Length) Coefficient() float64 {
	switch l {
	case Meter:
		return 1
	case Kilometer:
		return 1000
	case Centimeter:
		return 0.01
	case Millimeter:
		return 0.001
	case Micrometer:
		return 0.000001
	case Nanometer:
		return 0.000000001
	case Angstrom:
		return 0.0000000001
	case Inch:
		return 0.0254
	case Foot:
		return 0.3048
	case Yard:
		return 0.9144
	case Mile:
		return 1609.344
	case NauticalMile:
		return 1852
	case AstronomicalUnit:
		return 149597870700
	case LightYear:
		return 9460730472580800
	case Parsec:
		return 30856775814913600
	}
	return 1
}

func ParseLength(s string) (Length, error) {
	switch s {
	case "m":
		return Meter, nil
	case "km":
		return Kilometer, nil
	case "cm":
		return Centimeter, nil
	case "mm":
		return Millimeter, nil
	case "µm":
		return Micrometer, nil
	case "nm":
		return Nanometer, nil
	case "Å":
		return Angstrom, nil
	case "in":
		return Inch, nil
	case "ft":
		return Foot, nil
	case "yd":
		return Yard, nil
	case "mi":
		return Mile, nil
	case "nmi":
		return NauticalMile, nil
	case "au":
		return AstronomicalUnit, nil
	case "ly":
		return LightYear, nil
	case "pc":
		return Parsec, nil
	}
	return "", fmt.Errorf("unknown length unit: %s", s)
}
