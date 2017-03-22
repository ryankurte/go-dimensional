package scales

type Scale struct {
	Short      string
	Long       string
	Multiplier float64
}

var Nano = &Scale{"n", "nano", 1e-9}
var Micro = &Scale{"u", "micro", 1e-6}
var Milli = &Scale{"m", "milli", 1e-3}
var None = &Scale{"", "", 1e+0}
var Kilo = &Scale{"K", "kilo", 1e+3}
var Mega = &Scale{"M", "mega", 1e+6}
var Giga = &Scale{"G", "giga", 1e+9}

var Scales = []*Scale{None, Kilo, Mega, Giga, Milli, Micro, Nano}
