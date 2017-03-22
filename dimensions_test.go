package dimensional

import (
	"fmt"
	"github.com/ryankurte/go-dimensional/scales"
	"github.com/ryankurte/go-dimensional/units"
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
		_, err := Parse("100m/s/s")
		if err != nil {
			t.Error(err)
		}

		_, err = Parse("1.74m*s*s")
		if err != nil {
			t.Error(err)
		}

	})

}
