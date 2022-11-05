package mathunit

type DigitalStorage string

const (
	// Bit is the SI unit of digital storage
	Bit DigitalStorage = "b"
	// Kibibit is the SI unit of digital storage
	Kibibit DigitalStorage = "Kib"
	// Mebibit is the SI unit of digital storage
	Mebibit DigitalStorage = "Mib"
	// Gibibit is the SI unit of digital storage
	Gibibit DigitalStorage = "Gib"
	// Tebibit is the SI unit of digital storage
	Tebibit DigitalStorage = "Tib"
	// Pebibit is the SI unit of digital storage
	Pebibit DigitalStorage = "Pib"
	// Exbibit is the SI unit of digital storage
	Exbibit DigitalStorage = "Eib"
	// Zebibit is the SI unit of digital storage
	Zebibit DigitalStorage = "Zib"
	// Yobibit is the SI unit of digital storage
	Yobibit DigitalStorage = "Yib"

	// Byte is the SI unit of digital storage
	Byte DigitalStorage = "B"
	// Kibibyte is the SI unit of digital storage
	Kibibyte DigitalStorage = "KiB"
	// Mebibyte is the SI unit of digital storage
	Mebibyte DigitalStorage = "MiB"
	// Gibibyte is the SI unit of digital storage
	Gibibyte DigitalStorage = "GiB"
	// Tebibyte is the SI unit of digital storage
	Tebibyte DigitalStorage = "TiB"
	// Pebibyte is the SI unit of digital storage
	Pebibyte DigitalStorage = "PiB"
	// Exbibyte is the SI unit of digital storage
	Exbibyte DigitalStorage = "EiB"
	// Zebibyte is the SI unit of digital storage
	Zebibyte DigitalStorage = "ZiB"
	// Yobibyte is the SI unit of digital storage
	Yobibyte DigitalStorage = "YiB"

	// Kilobit is the SI unit of digital storage
	Kilobit DigitalStorage = "kb"
	// Megabit is the SI unit of digital storage
	Megabit DigitalStorage = "Mb"
	// Gigabit is the SI unit of digital storage
	Gigabit DigitalStorage = "Gb"
	// Terabit is the SI unit of digital storage
	Terabit DigitalStorage = "Tb"
	// Petabit is the SI unit of digital storage
	Petabit DigitalStorage = "Pb"
	// Exabit is the SI unit of digital storage
	Exabit DigitalStorage = "Eb"
	// Zettabit is the SI unit of digital storage
	Zettabit DigitalStorage = "Zb"
	// Yottabit is the SI unit of digital storage
	Yottabit DigitalStorage = "Yb"

	// Kilobyte is the SI unit of digital storage
	Kilobyte DigitalStorage = "kB"
	// Megabyte is the SI unit of digital storage
	Megabyte DigitalStorage = "MB"
	// Gigabyte is the SI unit of digital storage
	Gigabyte DigitalStorage = "GB"
	// Terabyte is the SI unit of digital storage
	Terabyte DigitalStorage = "TB"
	// Petabyte is the SI unit of digital storage
	Petabyte DigitalStorage = "PB"
	// Exabyte is the SI unit of digital storage
	Exabyte DigitalStorage = "EB"
	// Zettabyte is the SI unit of digital storage
	Zettabyte DigitalStorage = "ZB"
	// Yottabyte is the SI unit of digital storage
	Yottabyte DigitalStorage = "YB"
)

func (d DigitalStorage) String() string {
	return string(d)
}

func (d DigitalStorage) Values() []string {
	return []string{
		string(Bit),
		string(Kibibit),
		string(Mebibit),
		string(Gibibit),
		string(Tebibit),
		string(Pebibit),
		string(Exbibit),
		string(Zebibit),
		string(Yobibit),

		string(Byte),
		string(Kibibyte),
		string(Mebibyte),
		string(Gibibyte),
		string(Tebibyte),
		string(Pebibyte),
		string(Exbibyte),
		string(Zebibyte),
		string(Yobibyte),

		string(Kilobit),
		string(Megabit),
		string(Gigabit),
		string(Terabit),
		string(Petabit),
		string(Exabit),
		string(Zettabit),
		string(Yottabit),

		string(Kilobyte),
		string(Megabyte),
		string(Gigabyte),
		string(Terabyte),
		string(Petabyte),
		string(Exabyte),
		string(Zettabyte),
		string(Yottabyte),
	}
}

func (d DigitalStorage) Valid() bool {
	for _, v := range d.Values() {
		if v == string(d) {
			return true
		}
	}
	return false
}

func (d DigitalStorage) Coefficient() float64 {
	switch d {
	case Bit:
		return 1
	case Kibibit:
		return 1024
	case Mebibit:
		return 1048576
	case Gibibit:
		return 1073741824
	case Tebibit:
		return 1099511627776
	case Pebibit:
		return 1125899906842624
	case Exbibit:
		return 1152921504606846976
	case Zebibit:
		return 1180591620717411303424
	case Yobibit:
		return 1208925819614629174706176

	case Byte:
		return 8
	case Kibibyte:
		return 8192
	case Mebibyte:
		return 8388608
	case Gibibyte:
		return 8589934592
	case Tebibyte:
		return 8796093022208
	case Pebibyte:
		return 9007199254740992
	case Exbibyte:
		return 9223372036854775808
	case Zebibyte:
		return 9444732965739290427392
	case Yobibyte:
		return 966367641612104137216

	case Kilobit:
		return 1000
	case Megabit:
		return 1000000
	case Gigabit:
		return 1000000000
	case Terabit:
		return 1000000000000
	case Petabit:
		return 1000000000000000
	case Exabit:
		return 1000000000000000000
	case Zettabit:
		return 1000000000000000000000
	case Yottabit:
		return 1000000000000000000000000

	case Kilobyte:
		return 8000
	case Megabyte:
		return 8000000
	case Gigabyte:
		return 8000000000
	case Terabyte:
		return 8000000000000
	case Petabyte:
		return 8000000000000000
	case Exabyte:
		return 8000000000000000000
	case Zettabyte:
		return 8000000000000000000000
	case Yottabyte:
		return 8000000000000000000000000
	}
	return 1
}

func (d DigitalStorage) Mul(b Unit) (Unit, float64) {
	return nil, 0
}

func (d DigitalStorage) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case DigitalStorage:
		return Dimensionless, d.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
