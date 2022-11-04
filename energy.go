package mathunit

import "fmt"

type Energy string

const (
	// Joule is the SI unit of energy
	Joule Energy = "J"
	// Kilojoule is the SI unit of energy
	Kilojoule Energy = "kJ"
	// Megajoule is the SI unit of energy
	Megajoule Energy = "MJ"
	// Gigajoule is the SI unit of energy
	Gigajoule Energy = "GJ"
	// Terajoule is the SI unit of energy
	Terajoule Energy = "TJ"
	// Petajoule is the SI unit of energy
	Petajoule Energy = "PJ"
	// Exajoule is the SI unit of energy
	Exajoule Energy = "EJ"
	// Zettajoule is the SI unit of energy
	Zettajoule Energy = "ZJ"
	// Yottajoule is the SI unit of energy
	Yottajoule Energy = "YJ"

	// WattHour is the SI unit of energy
	WattHour Energy = "Wh"
	// KilowattHour is the SI unit of energy
	KilowattHour Energy = "kWh"
	// MegawattHour is the SI unit of energy
	MegawattHour Energy = "MWh"
	// GigawattHour is the SI unit of energy
	GigawattHour Energy = "GWh"
	// TerawattHour is the SI unit of energy
	TerawattHour Energy = "TWh"
	// PetawattHour is the SI unit of energy
	PetawattHour Energy = "PWh"
	// ExawattHour is the SI unit of energy
	ExawattHour Energy = "EWh"
	// ZettawattHour is the SI unit of energy
	ZettawattHour Energy = "ZWh"
	// YottawattHour is the SI unit of energy
	YottawattHour Energy = "YWh"

	// Calorie is the imperial unit of energy
	Calorie Energy = "cal"
	// Kilocalorie is the imperial unit of energy
	Kilocalorie Energy = "kcal"
	// Megacalorie is the imperial unit of energy
	Megacalorie Energy = "Mcal"
	// Gigacalorie is the imperial unit of energy
	Gigacalorie Energy = "Gcal"
	// Teracalorie is the imperial unit of energy
	Teracalorie Energy = "Tcal"
	// Petacalorie is the imperial unit of energy
	Petacalorie Energy = "Pcal"
	// Exacalorie is the imperial unit of energy
	Exacalorie Energy = "Ecal"
	// Zettacalorie is the imperial unit of energy
	Zettacalorie Energy = "Zcal"
	// Yottacalorie is the imperial unit of energy
	Yottacalorie Energy = "Ycal"
)

func (e Energy) String() string {
	return string(e)
}

func (e Energy) Values() []string {
	return []string{
		string(Joule),
		string(Kilojoule),
		string(Megajoule),
		string(Gigajoule),
		string(Terajoule),
		string(Petajoule),
		string(Exajoule),
		string(Zettajoule),
		string(Yottajoule),

		string(WattHour),
		string(KilowattHour),
		string(MegawattHour),
		string(GigawattHour),
		string(TerawattHour),
		string(PetawattHour),
		string(ExawattHour),
		string(ZettawattHour),
		string(YottawattHour),

		string(Calorie),
		string(Kilocalorie),
		string(Megacalorie),
		string(Gigacalorie),
		string(Teracalorie),
		string(Petacalorie),
		string(Exacalorie),
		string(Zettacalorie),
		string(Yottacalorie),
	}
}

func (e Energy) Valid() bool {
	for _, v := range e.Values() {
		if v == string(e) {
			return true
		}
	}
	return false
}

func (e Energy) Coefficient() float64 {
	switch e {
	case Joule:
		return 1
	case Kilojoule:
		return 1000
	case Megajoule:
		return 1000000
	case Gigajoule:
		return 1000000000
	case Terajoule:
		return 1000000000000
	case Petajoule:
		return 1000000000000000
	case Exajoule:
		return 1000000000000000000
	case Zettajoule:
		return 1000000000000000000000
	case Yottajoule:
		return 1000000000000000000000000

	case WattHour:
		return 3600
	case KilowattHour:
		return 3600000
	case MegawattHour:
		return 3600000000
	case GigawattHour:
		return 3600000000000
	case TerawattHour:
		return 3600000000000000
	case PetawattHour:
		return 3600000000000000000
	case ExawattHour:
		return 3600000000000000000000
	case ZettawattHour:
		return 3600000000000000000000000
	case YottawattHour:
		return 3600000000000000000000000000

	case Calorie:
		return 4.1868
	case Kilocalorie:
		return 4186.8
	case Megacalorie:
		return 4186800
	case Gigacalorie:
		return 4186800000
	case Teracalorie:
		return 4186800000000
	case Petacalorie:
		return 4186800000000000
	case Exacalorie:
		return 4186800000000000000
	case Zettacalorie:
		return 4186800000000000000000
	case Yottacalorie:
		return 4186800000000000000000000
	}
	return 1
}

func ParseEnergy(s string) (Energy, error) {
	switch s {
	case "J":
		return Joule, nil
	case "kJ":
		return Kilojoule, nil
	case "MJ":
		return Megajoule, nil
	case "GJ":
		return Gigajoule, nil
	case "TJ":
		return Terajoule, nil
	case "PJ":
		return Petajoule, nil
	case "EJ":
		return Exajoule, nil
	case "ZJ":
		return Zettajoule, nil
	case "YJ":
		return Yottajoule, nil

	case "Wh":
		return WattHour, nil
	case "kWh":
		return KilowattHour, nil
	case "MWh":
		return MegawattHour, nil
	case "GWh":
		return GigawattHour, nil
	case "TWh":
		return TerawattHour, nil
	case "PWh":
		return PetawattHour, nil
	case "EWh":
		return ExawattHour, nil
	case "ZWh":
		return ZettawattHour, nil
	case "YWh":
		return YottawattHour, nil

	case "cal":
		return Calorie, nil
	case "kcal":
		return Kilocalorie, nil
	case "Mcal":
		return Megacalorie, nil
	case "Gcal":
		return Gigacalorie, nil
	case "Tcal":
		return Teracalorie, nil
	case "Pcal":
		return Petacalorie, nil
	case "Ecal":
		return Exacalorie, nil
	case "Zcal":
		return Zettacalorie, nil
	case "Ycal":
		return Yottacalorie, nil
	}
	return "", fmt.Errorf("unknown energy unit %q", s)
}
