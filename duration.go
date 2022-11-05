package mathunit

type Duration string

const (
	// Second is the SI unit of duration
	Second Duration = "s"
	// Millisecond is the SI unit of duration
	Millisecond Duration = "ms"
	// Microsecond is the SI unit of duration
	Microsecond Duration = "µs"
	// Nanosecond is the SI unit of duration
	Nanosecond Duration = "ns"
	// Picosecond is the SI unit of duration
	Picosecond Duration = "ps"
	// Minute is the SI unit of duration
	Minute Duration = "min"
	// Hour is the SI unit of duration
	Hour Duration = "h"
	// Day is the SI unit of duration
	Day Duration = "d"
	// Week is the SI unit of duration
	Week Duration = "wk"
	// Year is the SI unit of duration
	Year Duration = "yr"
)

func (d Duration) String() string {
	return string(d)
}

func (d Duration) Coefficient() float64 {
	switch d {
	case Second:
		return 1
	case Millisecond:
		return 0.001
	case Microsecond:
		return 0.000001
	case Nanosecond:
		return 0.000000001
	case Picosecond:
		return 0.000000000001
	case Minute:
		return 60
	case Hour:
		return 3600
	case Day:
		return 86400
	case Week:
		return 604800
	case Year:
		return 31536000
	}
	return 1
}

func (d Duration) Values() []string {
	return []string{
		string(Second),
		string(Millisecond),
		string(Microsecond),
		string(Nanosecond),
		string(Picosecond),
		string(Minute),
		string(Hour),
		string(Day),
		string(Week),
		string(Year),
	}
}

func (d Duration) Valid() bool {
	for _, v := range d.Values() {
		if string(d) == v {
			return true
		}
	}
	return false
}

func (d Duration) Mul(b Unit) (Unit, float64) {
	switch b.(type) {
	case Speed:
		return Meter, d.Coefficient() * b.Coefficient()
	case Frequency:
		return Hertz, d.Coefficient() * b.Coefficient()
	}
	return nil, 0
}

func (d Duration) Div(b Unit) (Unit, float64) {
	switch b.(type) {
	case Duration:
		return Dimensionless, d.Coefficient() / b.Coefficient()
	}
	return nil, 0
}
