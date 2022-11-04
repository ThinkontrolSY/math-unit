package mathunit

import (
	"reflect"
	"testing"
)

func TestAngle(t *testing.T) {
	u := Parse("kg")
	reflect.TypeOf(u)

	t.Logf("u: %v", u)
	t.Logf("u: %T", u)
	t.Log(reflect.TypeOf(u))

	t.Log(reflect.TypeOf(Centimeter) == reflect.TypeOf(Kilometer))
}
