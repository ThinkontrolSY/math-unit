package mathunit

import (
	"errors"
	"reflect"
)

func (m Measure) Convert(to Unit) (Measure, error) {
	if reflect.TypeOf(m.unit) == reflect.TypeOf(to) {
		return Measure{to, m.value * m.unit.Coefficient() / to.Coefficient()}, nil
	}
	return Measure{}, errors.New("cannot convert between units")
}

func (m Measure) EQ(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value == other.value, e
}

func (m Measure) NE(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value != other.value, e
}

func (m Measure) GT(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value > other.value, e
}

func (m Measure) LT(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value < other.value, e
}

func (m Measure) GE(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value >= other.value, e
}

func (m Measure) LE(other Measure) (bool, error) {
	c, e := m.Convert(other.unit)
	return c.value <= other.value, e
}

func (m Measure) Add(other Measure) (Measure, error) {
	c, e := other.Convert(m.unit)
	return Measure{c.unit, c.value + m.value}, e
}

func (m Measure) Sub(other Measure) (Measure, error) {
	c, e := other.Convert(m.unit)
	return Measure{c.unit, c.value - m.value}, e
}

func (m Measure) Mul(other Measure) (Measure, error) {
	u, c := m.unit.Mul(other.unit)
	if u == nil {
		return Measure{}, errors.New("cannot multiply units")
	}
	return Measure{u, m.value * other.value * c}, nil
}

func (m Measure) Div(other Measure) (Measure, error) {
	u, c := m.unit.Div(other.unit)
	if u == nil {
		return Measure{}, errors.New("cannot divide units")
	}
	return Measure{u, m.value / other.value * c}, nil
}

func (m Measure) Multiply(v float64) Measure {
	return Measure{m.unit, m.value * v}
}
func (m Measure) Divide(v float64) Measure {
	return Measure{m.unit, m.value / v}
}
