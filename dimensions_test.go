package dimensional

import (
	"encoding/json"
	"fmt"
	"github.com/ryankurte/go-dimensional/scales"
	"github.com/ryankurte/go-dimensional/units"
	"reflect"
	"testing"
)

func TestDimensional(t *testing.T) {

	t.Run("Parses short scales", func(t *testing.T) {
		for _, scale := range scales.Scales {
			str := fmt.Sprintf("%s%s", scale.Short, units.Meter.Short)

			u, err := ParseBaseUnit(str)
			if err != nil {
				t.Error(err)
			}

			if u.Unit != units.Meter {
				t.Errorf("Invalid unit (actual %s, expected %s)", u.Unit.Long, units.Meter.Long)
			}

			if u.Scale != scale {
				t.Errorf("Invalid scale (actual %s, expected %s)", u.Scale.Short, scale.Short)
			}
		}
	})

	t.Run("Parses long scales", func(t *testing.T) {
		for _, scale := range scales.Scales {
			str := fmt.Sprintf("%s%s", scale.Long, units.Meter.Long)

			u, err := ParseBaseUnit(str)
			if err != nil {
				t.Error(err)
			}

			if u.Unit != units.Meter {
				t.Errorf("Invalid unit (actual %s, expected %s)", u.Unit.Long, units.Meter.Long)
			}

			if u.Scale != scale {
				t.Errorf("Invalid scale (actual %s, expected %s)", u.Scale.Long, scale.Long)
			}
		}
	})

	t.Run("Parses long units", func(t *testing.T) {
		for _, unit := range units.Units {
			str := fmt.Sprintf("%s", unit.Long)

			u, err := ParseBaseUnit(str)
			if err != nil {
				t.Error(err)
			}

			if u.Unit != unit {
				t.Errorf("Invalid unit (actual %s, expected %s)", u.Unit.Long, unit.Long)
			}

			if u.Scale != scales.None {
				t.Errorf("Invalid scale (actual %s, expected %s)", u.Scale.Long, scales.None.Long)
			}
		}
	})

	t.Run("Parses strings", func(t *testing.T) {
		d, err := Parse("100m/s/s")
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("Dimension: %s\n", d.String())

		d, err = Parse("1.74 Km*s*s")
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("Dimension: %s\n", d.String())
	})

	type dimensionContainer struct {
		Value Dimension
	}

	t.Run("Parses dimensions from generic decoders", func(t *testing.T) {

		var dc dimensionContainer
		err := json.Unmarshal([]byte(`{"value": "100 Km*s/mg"}`), &dc)
		if err != nil {
			t.Error(err)
		}

		d := dc.Value

		fmt.Printf("Dimension: %s\n", d.String())

		if expected := 100.0; d.float64 != expected {
			t.Errorf("Invalid value (actual %f, expected %f)", d.float64, expected)
		}

		if expected := []BaseUnit{{units.Meter, scales.Kilo}, {units.Second, scales.None}, {units.Gram, scales.Milli}}; !reflect.DeepEqual(d.unit.units, expected) {
			t.Errorf("Invalid units (actual %+v, expected %+v)", d.unit.units, expected)
		}

		if expected := []int{1, 1, -1}; !reflect.DeepEqual(d.unit.exponents, expected) {
			t.Errorf("Invalid exponents (actual %+v, expected %+v)", d.unit.exponents, expected)
		}

	})

}
