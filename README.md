# Go-Dimensional

An experiment in dimensional consistency in Golang.

The idea is that dimensions can be created (or parsed) with units that can then be used in calculations and to enforce dimensional consistency. 

Check out [pypi/units](https://pypi.python.org/pypi/units) for an example in Python.

## Status

Experimental / WIP. If you feel like adding a feature, pull requests would be greatly appreciated!

### Todos

- [ ] Compressor function to combine units & exponents (eg. `m * s * s -> m * s^2`)
- [ ] Multiplication and division methods that manage units (eg. `m/s * s = m`)
- [ ] Pretty printer with superscript notation (eg. `9.98 m/s²`)
- [ ] Exponent / Carat parsing (eg. `32.7e8 m/s^2`)
- [ ] Base type constructors (eg. `MHz(433.3)` )

