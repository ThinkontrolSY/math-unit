package mathunit

import "fmt"

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

func ParsePackageUnit(s string) (PackageUnit, error) {
	switch s {
	case "doz":
		return Dozen, nil
	case "ddoz":
		return Doubledozen, nil
	case "hdoz":
		return Halfdozen, nil
	case "3doz":
		return Threedozens, nil
	case "4doz":
		return Fourdozens, nil
	case "grs":
		return Gross, nil
	case "ggrs":
		return Greatgross, nil
	case "box":
		return Box, nil
	case "bag":
		return Bag, nil
	case "btl":
		return Bottle, nil
	case "can":
		return Can, nil
	}
	return "", fmt.Errorf("unknown package unit: %s", s)
}
