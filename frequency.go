package mathunit

type Frequency string

const (
	// Hertz is the SI unit of frequency
	Hertz Frequency = "Hz"
	// Kilohertz is the SI unit of frequency
	Kilohertz Frequency = "kHz"
	// Megahertz is the SI unit of frequency
	Megahertz Frequency = "MHz"
	// Gigahertz is the SI unit of frequency
	Gigahertz Frequency = "GHz"
	// Terahertz is the SI unit of frequency
	Terahertz Frequency = "THz"
	// Petahertz is the SI unit of frequency
	Petahertz Frequency = "PHz"
	// Exahertz is the SI unit of frequency
	Exahertz Frequency = "EHz"
	// Zettahertz is the SI unit of frequency
	Zettahertz Frequency = "ZHz"
	// Yottahertz is the SI unit of frequency
	Yottahertz Frequency = "YHz"
	// RadianPerSecond is the SI unit of frequency
	RadianPerSecond Frequency = "rad/s"
	// RevolutionsPerMinute is the SI unit of frequency
	RevolutionsPerMinute Frequency = "rpm"
	// RevolutionsPerSecond is the SI unit of frequency
	RevolutionsPerSecond Frequency = "rps"
	// RevolutionsPerHour is the SI unit of frequency
	RevolutionsPerHour Frequency = "rph"
	// BeatsPerMinute is the SI unit of frequency
	BeatsPerMinute Frequency = "bpm"
	// BeatsPerSecond is the SI unit of frequency
	BeatsPerSecond Frequency = "bps"
	// BeatsPerHour is the SI unit of frequency
	BeatsPerHour Frequency = "bph"
	// CyclesPerSecond is the SI unit of frequency
	CyclesPerSecond Frequency = "cps"
	// CyclesPerMinute is the SI unit of frequency
	CyclesPerMinute Frequency = "cpm"
	// CyclesPerHour is the SI unit of frequency
	CyclesPerHour Frequency = "cph"
	// Decahertz is the SI unit of frequency
	Decahertz Frequency = "daHz"
	// Decihertz is the SI unit of frequency
	Decihertz Frequency = "dHz"
	// Centihertz is the SI unit of frequency
	Centihertz Frequency = "cHz"
	// Millihertz is the SI unit of frequency
	Millihertz Frequency = "mHz"
	// Microhertz is the SI unit of frequency
	Microhertz Frequency = "µHz"
	// Nanohertz is the SI unit of frequency
	Nanohertz Frequency = "nHz"
	// Picohertz is the SI unit of frequency
	Picohertz Frequency = "pHz"
	// Femtohertz is the SI unit of frequency
	Femtohertz Frequency = "fHz"
	// Attohertz is the SI unit of frequency
	Attohertz Frequency = "aHz"
	// Zeptohertz is the SI unit of frequency
	Zeptohertz Frequency = "zHz"
	// Yoctohertz is the SI unit of frequency
	Yoctohertz Frequency = "yHz"
)

func (f Frequency) String() string {
	return string(f)
}

func (f Frequency) Values() []string {
	return []string{
		"Hz",
		"kHz",
		"MHz",
		"GHz",
		"THz",
		"PHz",
		"EHz",
		"ZHz",
		"YHz",
		"rad/s",
		"rpm",
		"rps",
		"rph",
		"bpm",
		"bps",
		"bph",
		"cps",
		"cpm",
		"cph",
		"daHz",
		"dHz",
		"cHz",
		"mHz",
		"µHz",
		"nHz",
		"pHz",
		"fHz",
		"aHz",
		"zHz",
		"yHz",
	}
}

func (f Frequency) Valid() bool {
	for _, v := range f.Values() {
		if v == string(f) {
			return true
		}
	}
	return false
}

func (f Frequency) Coefficient() float64 {
	switch f {
	case Hertz:
		return 1
	case Kilohertz:
		return 1000
	case Megahertz:
		return 1000000
	case Gigahertz:
		return 1000000000
	case Terahertz:
		return 1000000000000
	case Petahertz:
		return 1000000000000000
	case Exahertz:
		return 1000000000000000000
	case Zettahertz:
		return 1000000000000000000000
	case Yottahertz:
		return 1000000000000000000000000
	case RadianPerSecond:
		return 159154943091895
	case RevolutionsPerMinute:
		return 1047197551197
	case RevolutionsPerSecond:
		return 62831853071795
	case RevolutionsPerHour:
		return 1047197551197
	case BeatsPerMinute:
		return 1047197551197
	case BeatsPerSecond:
		return 62831853071795
	case BeatsPerHour:
		return 1047197551197
	case CyclesPerSecond:
		return 62831853071795
	case CyclesPerMinute:
		return 1047197551197
	case CyclesPerHour:
		return 1047197551197
	case Decahertz:
		return 10
	case Decihertz:
		return 0.1
	case Centihertz:
		return 0.01
	case Millihertz:
		return 0.001
	case Microhertz:
		return 0.000001
	case Nanohertz:
		return 1e-09
	case Picohertz:
		return 1e-12
	case Femtohertz:
		return 1e-15
	case Attohertz:
		return 1e-18
	case Zeptohertz:
		return 1e-21
	case Yoctohertz:
		return 1e-24
	}
	return 1
}

func (f Frequency) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case Duration:
		return Radian, Radian.Coefficient() * f.Coefficient() / b.Coefficient()
	}
	return nil, 0
}

func (f Frequency) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case Frequency:
		return Dimensionless, f.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
