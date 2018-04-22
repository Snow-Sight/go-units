package units

import "fmt"

// Element represents a quantity and unit
type Element struct {
	Quantity float64
	Unit     Unit
}

// Unit represents a unit
type Unit struct {
	FullName           string
	Abbreviation       string
	order              int32
	RatioToCounterpart float64
	SICounterPart      string
}

// UnitMap maps strings to their base
var UnitMap = map[string]Unit{
	"Centimetre": Centimetre,
	"Inch":       Inch,
}

// Errors

// ErrUnitMustBeLength returns the error for when the unit has to be centimetre or inch
var ErrUnitMustBeLength = fmt.Errorf("Unit must be centimetre or inch")

// CreateElement returns an element comprised of the quanitity and unit
func CreateElement(q float64, u Unit) *Element {
	return &Element{q, u}
}

// ConvertToSIBase takes an element and converts it to it's SI counterpart
func (e *Element) ConvertToSIBase() *Element {
	return &Element{e.Quantity * e.Unit.RatioToCounterpart, UnitMap[e.Unit.SICounterPart]}
}

// Centimetre is the SI base -3 unit for length
var Centimetre = Unit{
	"Centimetre",
	"cm",
	-3,
	1,
	"Centimetre",
}

// Inch is the smallest imperial length
var Inch = Unit{
	"Inch",
	"inch",
	1,
	2.54,
	"Centimetre",
}
