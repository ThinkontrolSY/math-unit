package mathunit

import "fmt"

type Pressure string

const (
	// Pascal is the SI unit of pressure
	Pascal Pressure = "Pa"
	// KiloPascal is the SI unit of pressure
	KiloPascal Pressure = "kPa"
	// MegaPascal is the SI unit of pressure
	MegaPascal Pressure = "MPa"
	// GigaPascal is the SI unit of pressure
	GigaPascal Pressure = "GPa"
	// TeraPascal is the SI unit of pressure
	TeraPascal Pressure = "TPa"
	// PetaPascal is the SI unit of pressure
	PetaPascal Pressure = "PPa"
	// ExaPascal is the SI unit of pressure
	ExaPascal Pressure = "EPa"
	// ZettaPascal is the SI unit of pressure
	ZettaPascal Pressure = "ZPa"
	// YottaPascal is the SI unit of pressure
	YottaPascal Pressure = "YPa"
	// HectoPascal is the SI unit of pressure
	HectoPascal Pressure = "hPa"
	// DekaPascal is the SI unit of pressure
	DekaPascal Pressure = "daPa"
	// DeciPascal is the SI unit of pressure
	DeciPascal Pressure = "dPa"
	// CentiPascal is the SI unit of pressure
	CentiPascal Pressure = "cPa"
	// MilliPascal is the SI unit of pressure
	MilliPascal Pressure = "mPa"
	// MicroPascal is the SI unit of pressure
	MicroPascal Pressure = "µPa"
	// NanoPascal is the SI unit of pressure
	NanoPascal Pressure = "nPa"
	// PicoPascal is the SI unit of pressure
	PicoPascal Pressure = "pPa"
	// FemtoPascal is the SI unit of pressure
	FemtoPascal Pressure = "fPa"
	// AttoPascal is the SI unit of pressure
	AttoPascal Pressure = "aPa"
	// ZeptoPascal is the SI unit of pressure
	ZeptoPascal Pressure = "zPa"
	// YoctoPascal is the SI unit of pressure
	YoctoPascal Pressure = "yPa"

	// Bar is the SI unit of pressure
	Bar Pressure = "bar"
	// MilliBar is the SI unit of pressure
	MilliBar Pressure = "mbar"
	// MicroBar is the SI unit of pressure
	MicroBar Pressure = "µbar"
	// NanoBar is the SI unit of pressure
	NanoBar Pressure = "nbar"
	// PicoBar is the SI unit of pressure
	PicoBar Pressure = "pbar"
	// FemtoBar is the SI unit of pressure
	FemtoBar Pressure = "fbar"
	// AttoBar is the SI unit of pressure
	AttoBar Pressure = "abar"
	// ZeptoBar is the SI unit of pressure
	ZeptoBar Pressure = "zbar"
	// YoctoBar is the SI unit of pressure
	YoctoBar Pressure = "ybar"

	// Atmosphere is the SI unit of pressure
	Atmosphere Pressure = "atm"
	// TechnicalAtmosphere is the SI unit of pressure
	TechnicalAtmosphere Pressure = "at"
	// MilliAtmosphere is the SI unit of pressure
	MilliAtmosphere Pressure = "matm"
	// MicroAtmosphere is the SI unit of pressure
	MicroAtmosphere Pressure = "µatm"
	// NanoAtmosphere is the SI unit of pressure
	NanoAtmosphere Pressure = "natm"
	// PicoAtmosphere is the SI unit of pressure
	PicoAtmosphere Pressure = "patm"
	// FemtoAtmosphere is the SI unit of pressure
	FemtoAtmosphere Pressure = "fatm"
	// AttoAtmosphere is the SI unit of pressure
	AttoAtmosphere Pressure = "aatm"
	// ZeptoAtmosphere is the SI unit of pressure
	ZeptoAtmosphere Pressure = "zatm"
	// YoctoAtmosphere is the SI unit of pressure
	YoctoAtmosphere Pressure = "yatm"

	// Torr is the SI unit of pressure
	Torr Pressure = "Torr"
	// MilliTorr is the SI unit of pressure
	MilliTorr Pressure = "mTorr"
	// MicroTorr is the SI unit of pressure
	MicroTorr Pressure = "µTorr"
	// NanoTorr is the SI unit of pressure
	NanoTorr Pressure = "nTorr"
	// PicoTorr is the SI unit of pressure
	PicoTorr Pressure = "pTorr"
	// FemtoTorr is the SI unit of pressure
	FemtoTorr Pressure = "fTorr"
	// AttoTorr is the SI unit of pressure
	AttoTorr Pressure = "aTorr"
	// ZeptoTorr is the SI unit of pressure
	ZeptoTorr Pressure = "zTorr"
	// YoctoTorr is the SI unit of pressure
	YoctoTorr Pressure = "yTorr"
)

func (p Pressure) String() string {
	return string(p)
}

func (p Pressure) Values() []string {
	return []string{
		string(Pascal),
		string(KiloPascal),
		string(MegaPascal),
		string(GigaPascal),
		string(TeraPascal),
		string(PetaPascal),
		string(ExaPascal),
		string(ZettaPascal),
		string(YottaPascal),
		string(HectoPascal),
		string(DekaPascal),
		string(DeciPascal),
		string(CentiPascal),
		string(MilliPascal),
		string(MicroPascal),
		string(NanoPascal),
		string(PicoPascal),
		string(FemtoPascal),
		string(AttoPascal),
		string(ZeptoPascal),
		string(YoctoPascal),
		string(Bar),
		string(MilliBar),
		string(MicroBar),
		string(NanoBar),
		string(PicoBar),
		string(FemtoBar),
		string(AttoBar),
		string(ZeptoBar),
		string(YoctoBar),
		string(Atmosphere),
		string(TechnicalAtmosphere),
		string(MilliAtmosphere),
		string(MicroAtmosphere),
		string(NanoAtmosphere),
		string(PicoAtmosphere),
		string(FemtoAtmosphere),
		string(AttoAtmosphere),
		string(ZeptoAtmosphere),
		string(YoctoAtmosphere),
		string(Torr),
		string(MilliTorr),
		string(MicroTorr),
		string(NanoTorr),
		string(PicoTorr),
		string(FemtoTorr),
		string(AttoTorr),
		string(ZeptoTorr),
		string(YoctoTorr),
	}
}

func (p Pressure) Valid() bool {
	_, err := ParsePressure(p.String())
	return err == nil
}

func (p Pressure) Coefficient() float64 {
	switch p {
	case Pascal:
		return 1
	case KiloPascal:
		return 1e3
	case MegaPascal:
		return 1e6
	case GigaPascal:
		return 1e9
	case TeraPascal:
		return 1e12
	case PetaPascal:
		return 1e15
	case ExaPascal:
		return 1e18
	case ZettaPascal:
		return 1e21
	case YottaPascal:
		return 1e24
	case HectoPascal:
		return 1e2
	case DekaPascal:
		return 1e1
	case DeciPascal:
		return 1e-1
	case CentiPascal:
		return 1e-2
	case MilliPascal:
		return 1e-3
	case MicroPascal:
		return 1e-6
	case NanoPascal:
		return 1e-9
	case PicoPascal:
		return 1e-12
	case FemtoPascal:
		return 1e-15
	case AttoPascal:
		return 1e-18
	case ZeptoPascal:
		return 1e-21
	case YoctoPascal:
		return 1e-24
	case Bar:
		return 1e5
	case MilliBar:
		return 1e2
	case MicroBar:
		return 1e-1
	case NanoBar:
		return 1e-4
	case PicoBar:
		return 1e-7
	case FemtoBar:
		return 1e-10
	case AttoBar:
		return 1e-13
	case ZeptoBar:
		return 1e-16
	case YoctoBar:
		return 1e-19
	case Atmosphere:
		return 101325
	case TechnicalAtmosphere:
		return 98066.5
	case MilliAtmosphere:
		return 101.325
	case MicroAtmosphere:
		return 0.101325
	case NanoAtmosphere:
		return 0.000101325
	case PicoAtmosphere:
		return 0.000000101325
	case FemtoAtmosphere:
		return 0.000000000101325
	case AttoAtmosphere:
		return 0.000000000000101325
	case ZeptoAtmosphere:
		return 0.000000000000000101325
	case YoctoAtmosphere:
		return 0.000000000000000000101325
	case Torr:
		return 133.322
	case MilliTorr:
		return 0.133322
	case MicroTorr:
		return 0.000133322
	case NanoTorr:
		return 0.000000133322
	case PicoTorr:
		return 0.000000000133322
	case FemtoTorr:
		return 0.000000000000133322
	case AttoTorr:
		return 0.000000000000000133322
	case ZeptoTorr:
		return 0.000000000000000000133322
	case YoctoTorr:
		return 0.000000000000000000000133322
	default:
		return 1
	}
}

func ParsePressure(s string) (Pressure, error) {
	switch s {
	case "Pa":
		return Pascal, nil
	case "kPa":
		return KiloPascal, nil
	case "MPa":
		return MegaPascal, nil
	case "GPa":
		return GigaPascal, nil
	case "TPa":
		return TeraPascal, nil
	case "PPa":
		return PetaPascal, nil
	case "EPa":
		return ExaPascal, nil
	case "ZPa":
		return ZettaPascal, nil
	case "YPa":
		return YottaPascal, nil
	case "hPa":
		return HectoPascal, nil
	case "daPa":
		return DekaPascal, nil
	case "dPa":
		return DeciPascal, nil
	case "cPa":
		return CentiPascal, nil
	case "mPa":
		return MilliPascal, nil
	case "µPa":
		return MicroPascal, nil
	case "nPa":
		return NanoPascal, nil
	case "pPa":
		return PicoPascal, nil
	case "fPa":
		return FemtoPascal, nil
	case "aPa":
		return AttoPascal, nil
	case "zPa":
		return ZeptoPascal, nil
	case "yPa":
		return YoctoPascal, nil
	case "bar":
		return Bar, nil
	case "mbar":
		return MilliBar, nil
	case "µbar":
		return MicroBar, nil
	case "nbar":
		return NanoBar, nil
	case "pbar":
		return PicoBar, nil
	case "fbar":
		return FemtoBar, nil
	case "abar":
		return AttoBar, nil
	case "zbar":
		return ZeptoBar, nil
	case "ybar":
		return YoctoBar, nil
	case "atm":
		return Atmosphere, nil
	case "at":
		return TechnicalAtmosphere, nil
	case "matm":
		return MilliAtmosphere, nil
	case "µatm":
		return MicroAtmosphere, nil
	case "natm":
		return NanoAtmosphere, nil
	case "patm":
		return PicoAtmosphere, nil
	case "fatm":
		return FemtoAtmosphere, nil
	case "aatm":
		return AttoAtmosphere, nil
	case "zatm":
		return ZeptoAtmosphere, nil
	case "yatm":
		return YoctoAtmosphere, nil
	case "torr":
		return Torr, nil
	case "mtorr":
		return MilliTorr, nil
	case "µtorr":
		return MicroTorr, nil
	case "ntorr":
		return NanoTorr, nil
	case "ptorr":
		return PicoTorr, nil
	case "ftorr":
		return FemtoTorr, nil
	case "atorr":
		return AttoTorr, nil
	case "ztorr":
		return ZeptoTorr, nil
	case "ytorr":
		return YoctoTorr, nil
	default:
		return "", fmt.Errorf("unknown pressure unit %q", s)
	}
}
