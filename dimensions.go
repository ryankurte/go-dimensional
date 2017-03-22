package dimensional

import (
	"fmt"
	"github.com/ryankurte/go-dimensional/scales"
	"github.com/ryankurte/go-dimensional/units"
	"regexp"
	"strconv"
	"strings"
)

// BaseUnit is a physical unit with associated scale
// For example "km"
type BaseUnit struct {
	Unit  *units.Unit
	Scale *scales.Scale
}

// ComplexUnit is a chain of base units and exponents
// For example, 'km/s*s'
type ComplexUnit struct {
	units     []BaseUnit
	exponents []int
}

// Dimension is a dimensional object with a value and complex unit
type Dimension struct {
	float64
	unit ComplexUnit
}

func ParseBaseUnit(unitString string) (*BaseUnit, error) {

	// Match against base units
	var unit *units.Unit
	suffixString := ""

	// Attempt to match long units
	for _, u := range units.Units {
		if strings.HasSuffix(unitString, u.Long) {
			unit = u
			suffixString = u.Long
			break
		}
	}

	// If there was no unit match, match short units
	if unit == nil {
		for _, u := range units.Units {
			if strings.HasSuffix(unitString, u.Short) {
				unit = u
				suffixString = u.Short
				break
			}
		}
	}

	// Check units matched OK
	if unit == nil {
		return nil, fmt.Errorf("Unrecognised unit parsing: '%s'", unitString)
	}

	// Remove unit suffix to leave prefix only
	prefixString := strings.Replace(unitString, suffixString, "", 1)

	// Attempt to match long scales
	scale := scales.None
	for _, s := range scales.Scales {
		if s.Long != "" && strings.HasPrefix(prefixString, s.Long) {
			scale = s
			break
		}
	}

	// If there was no scale match, match short scales
	if scale == scales.None {
		for _, s := range scales.Scales {
			if s.Short != "" && strings.HasPrefix(prefixString, s.Short) {
				scale = s
				break
			}
		}
	}

	return &BaseUnit{unit, scale}, nil
}

var parserRegex = regexp.MustCompile(`(?m)^([0-9\.e*\-\^]+)[ ]*([a-zA-Z\/\*]+)$`)
var unitRegex = regexp.MustCompile(`(?m)^(([\/\*]{0,1})([a-zA-Z]+))+$`)

func Parse(s string) (*Dimension, error) {
	matches := parserRegex.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("Parser regex failed for input: '%s'", s)
	}

	if len(matches) != 3 {
		return nil, fmt.Errorf("Invalid parser input: '%s', input should be a string containing a number then a unit", s)
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return nil, err
	}

	dimension := Dimension{float64: value}

	unitStrings := unitRegex.FindStringSubmatch(matches[2])
	fmt.Printf("Unit strings: %+v\n", unitStrings)

	/*
		for i, unitBlock := range unitStrings {

			fmt.Printf("Unit block: `%s`\n", unitBlock)

			exponent := 1
			if unitBlock[1] == "/" {
				exponent = -1
			}

			unit, err := ParseBaseUnit(unitBlock[2])
			if err != nil {
				return err
			}

			dimension.unit.units[i] = *unit
			dimension.unit.exponents[i] = exponent

		}

		fmt.Printf("Object: %+v", dimension)
	*/
	return &dimension, nil
}
