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

func NewMeasure(unit Unit, value float64) Measure {
	return Measure{unit: unit, value: value}
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
	return reflect.TypeOf(m.unit) == reflect.TypeOf(other.unit)
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
