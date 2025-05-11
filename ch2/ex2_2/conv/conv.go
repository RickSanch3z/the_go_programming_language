package conv

import (
	"fmt"
	"errors"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Kilogram float64
type Pound float64
type Meter float64
type Feet float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

// methods for types --------
func (c Celsius) String() string {
	return fmt.Sprintf("%.6g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.6g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.6g°K", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.6glb", p)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%.6gkg", kg)
}

func (m Meter) String() string {
	return fmt.Sprintf("%.6gm", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%.6gft", f)
}
// --------------------------

// Convertion functions -----
// CToF converts Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// FToC converts Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// KToC converts Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

// FToK converts Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}

// KToF converts Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}

// KgToP converts Kilogram to Pound 
func KgToP(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}

// PToKg converts Pound to Kilogram
func PToKg(p Pound) Kilogram {
	return Kilogram(p / 2.20462)
}

// MToFt converts Meter to Feet 
func MToFt(m Meter) Feet {
	return Feet(m / 3.28084)
}

// FtToM converts Feet to Meter
func FtToM(ft Feet) Meter {
	return Meter(ft * 3.28084)
}
// --------------------------

func ConvertUnit(convT string, value string) (string, error) {
	retValue := ""

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", err
	}

	switch convT {
	case "CToF":
		retValue = fmt.Sprintf("%.6g°C : %s", v, CToF(Celsius(v)).String())
	case "FToC":
		retValue = fmt.Sprintf("%.6g°F : %s", v, FToC(Fahrenheit(v)).String())
	case "CToK":
		retValue = fmt.Sprintf("%.6g°C : %s", v, CToK(Celsius(v)).String())
	case "KToC":
		retValue = fmt.Sprintf("%.6g°K : %s", v, KToC(Kelvin(v)).String())
	case "FToK":
		retValue = fmt.Sprintf("%.6g°F : %s", v, FToK(Fahrenheit(v)).String())
	case "KToF":
		retValue = fmt.Sprintf("%.6g°K : %s", v, KToF(Kelvin(v)).String())
	case "KgToP":
		retValue = fmt.Sprintf("%.6gkg : %s", v, KgToP(Kilogram(v)).String())
	case "PToKg":
		retValue = fmt.Sprintf("%.6glb : %s", v, PToKg(Pound(v)).String())
	case "MToFt":
		retValue = fmt.Sprintf("%.6gm : %s", v, MToFt(Meter(v)).String())
	case "FtToM":
		retValue = fmt.Sprintf("%.6gft : %s", v, FtToM(Feet(v)).String())
	default:
		return "", errors.New("no convertion has been found")
	}

	return retValue, nil
}

func ValidConv(convT string) error {
	switch convT {
	case "CToF":
	case "FToC":
	case "CToK":
	case "KToC":
	case "FToK":
	case "KToF":
	case "KgToP":
	case "PToKg":
	case "MToFt":
	case "FtToM":
	default:
		return errors.New("convertion is not valid")
	}

	return nil
}
