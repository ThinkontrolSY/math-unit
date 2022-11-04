package mathunit

import "fmt"

// Mass represents a SI unit of mass (in grams, G)
type Mass string

const (
	// Gram is the SI unit of mass
	Gram Mass = "g"
	// Kilogram is the SI unit of mass
	Kilogram Mass = "kg"
	// Ton is the SI unit of mass
	Ton Mass = "t"

	Microgram Mass = "µg"
	Milligram Mass = "mg"

	// Pound is the imperial unit of mass
	Pound Mass = "lb"
	// Ounce is the imperial unit of mass
	Ounce Mass = "oz"

	// Stone is the imperial unit of mass
	Stone Mass = "st"

	// Slug is the imperial unit of mass
	Slug Mass = "slug"

	// Carat is the imperial unit of mass
	Carat Mass = "ct"
)

func (m Mass) String() string {
	return string(m)
}

func (m Mass) Values() []string {
	return []string{
		"g",
		"kg",
		"t",
		"µg",
		"mg",
		"lb",
		"oz",
		"st",
		"slug",
		"ct",
	}
}

func (m Mass) Valid() bool {
	switch m {
	case Gram, Kilogram, Ton, Microgram, Milligram, Pound, Ounce, Stone, Slug, Carat:
		return true
	}
	return false
}

func (m Mass) Coefficient() float64 {
	switch m {
	case Gram:
		return 1
	case Kilogram:
		return 1000
	case Ton:
		return 1000000
	case Microgram:
		return 0.000001
	case Milligram:
		return 0.001
	case Pound:
		return 453.59237
	case Ounce:
		return 28.349523125
	case Stone:
		return 6350.29318
	case Slug:
		return 14593.90294
	case Carat:
		return 0.2
	}
	return 1
}

func ParseMass(s string) (Mass, error) {
	switch s {
	case "g":
		return Gram, nil
	case "kg":
		return Kilogram, nil
	case "t":
		return Ton, nil
	case "µg":
		return Microgram, nil
	case "mg":
		return Milligram, nil
	case "lb":
		return Pound, nil
	case "oz":
		return Ounce, nil
	case "st":
		return Stone, nil
	case "slug":
		return Slug, nil
	case "ct":
		return Carat, nil
	}
	return "", fmt.Errorf("unknown mass unit %q", s)
}
