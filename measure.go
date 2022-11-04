package mathunit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type Measure struct {
	unit  Unit
	value float64
}

func (m Measure) String() string {
	return fmt.Sprintf("%v %v", m.value, m.unit)
}

func (m Measure) Unit() Unit {
	return m.unit
}

func (m Measure) Value() float64 {
	return m.value
}

func (m Measure) Coefficient() float64 {
	return m.unit.Coefficient()
}

func (m Measure) Valid() bool {
	return m.unit.Valid()
}

func (m Measure) Values() []string {
	return m.unit.Values()
}

func (m Measure) Compatable(other Measure) bool {
	return reflect.TypeOf(m) == reflect.TypeOf(other)
}

func (m Measure) Convert(to Unit) Measure {
	return Measure{to, m.value * m.unit.Coefficient() / to.Coefficient()}
}

func (m Measure) Add(other Measure) Measure {
	return Measure{m.unit, m.value + other.Convert(m.unit).value}
}

func (m Measure) Sub(other Measure) Measure {
	return Measure{m.unit, m.value - other.Convert(m.unit).value}
}

func (m Measure) EQ(other Measure) bool {
	return m.value == other.Convert(m.unit).value
}

func (m Measure) NEQ(other Measure) bool {
	return m.value != other.Convert(m.unit).value
}

func (m Measure) LT(other Measure) bool {
	return m.value < other.Convert(m.unit).value
}

func (m Measure) LTE(other Measure) bool {
	return m.value <= other.Convert(m.unit).value
}

func (m Measure) GT(other Measure) bool {
	return m.value > other.Convert(m.unit).value
}

func (m Measure) GTE(other Measure) bool {
	return m.value >= other.Convert(m.unit).value
}

func (m Measure) Compare(other Measure) int {
	if m.value < other.Convert(m.unit).value {
		return -1
	} else if m.value > other.Convert(m.unit).value {
		return 1
	}
	return 0
}

func (v Measure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	}{
		Unit:  v.unit.String(),
		Value: v.value,
	})
}

func (d *Measure) UnmarshalJSON(v []byte) error {
	var tmp struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	}
	json.Unmarshal(v, &tmp)
	d.value = tmp.Value
	unit := Parse(tmp.Unit)
	if unit == nil {
		return errors.New("invalid unit")
	}
	d.unit = unit
	return nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (t *Measure) UnmarshalGQL(v interface{}) error {
	if m, ok := v.(map[string]interface{}); ok {
		if unit, ok := m["unit"]; ok {
			if ustr, ok := unit.(string); ok {
				u := Parse(ustr)
				if u == nil {
					return errors.New("invalid unit")
				}
				t.unit = u
			} else {
				return errors.New("invalid unit")
			}
		} else {
			return errors.New("unit is required")
		}
		if value, ok := m["value"]; ok {
			switch n := value.(type) {
			case float64:
				t.value = n
			case int:
				t.value = float64(n)
			case json.Number:
				f, err := n.Float64()
				if err != nil {
					return err
				}
				t.value = f
			}
		}
		return nil
	}
	return errors.New("map unmarshal from unknow type")
}

// MarshalGQL implements the graphql.Marshaler interface
func (t Measure) MarshalGQL(w io.Writer) {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
