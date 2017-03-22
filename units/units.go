package units

// Unit is a physical unit type
type Unit struct {
	Short string
	Long  string
}

var Gram = &Unit{"g", "gram"}
var Second = &Unit{"s", "second"}
var Kelvin = &Unit{"K", "kelvin"}
var Amp = &Unit{"A", "ampere"}
var Mol = &Unit{"mol", "mol"}
var Candela = &Unit{"cd", "candela"}
var Meter = &Unit{"m", "meter"}
var Hertz = &Unit{"Hz", "hertz"}

var Units = []*Unit{Gram, Second, Kelvin, Amp, Mol, Candela, Meter, Hertz}
