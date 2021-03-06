package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/relnod/calcgo/parser"
)

// Errors that can occur during calculation or conversion.
var (
	ErrorInvalidInteger     = errors.New("Invalid Integer")
	ErrorInvalidDecimal     = errors.New("Invalid Decimal")
	ErrorInvalidBinary      = errors.New("Invalid Binary")
	ErrorInvalidHexadecimal = errors.New("Invalid Hexadecimal")
	ErrorInvalidExponential = errors.New("Invalid Exponential")
	ErrorDivisionByZero     = errors.New("Division by zero")
)

// ConvertInteger converts an integer string to a float64.
// Returns an error if conversion failed.
func ConvertInteger(value string) (float64, error) {
	integer, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, ErrorInvalidInteger
	}
	return float64(integer), nil
}

// ConvertDecimal converts a decimal string to a float64.
// Returns an error if conversion failed.
func ConvertDecimal(value string) (float64, error) {
	decimal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, ErrorInvalidDecimal
	}
	return decimal, nil
}

// ConvertBin converts a binary string to a float64.
// Returns an error if conversion failed.
func ConvertBin(value string) (float64, error) {
	val := strings.Replace(value, "0b", "", 1)

	bin, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		return 0, ErrorInvalidBinary
	}

	return float64(bin), nil
}

// ConvertHex converts a hex string to a float64.
// Returns an error if conversion failed.
func ConvertHex(value string) (float64, error) {
	hexa, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, ErrorInvalidHexadecimal
	}

	return float64(hexa), nil
}

// ConvertExponential converts an exponential string to a float64.
// Returns an error if conversion failed.
func ConvertExponential(value string) (float64, error) {
	splitted := strings.Split(value, "^")
	if len(splitted) != 2 {
		return 0, ErrorInvalidExponential
	}

	base, err := strconv.Atoi(splitted[0])
	if err != nil {
		return 0, ErrorInvalidExponential
	}

	exponent, err := strconv.Atoi(splitted[1])
	if err != nil {
		return 0, ErrorInvalidExponential
	}

	res := math.Pow(float64(base), float64(exponent))
	if math.IsInf(res, 1) {
		return 0, ErrorInvalidExponential
	}

	return res, nil
}

// ConvertLiteral converts a atring literal to a float.
func ConvertLiteral(value string, nodeType parser.NodeType) (float64, error) {
	switch nodeType {
	case parser.NInt:
		return ConvertInteger(value)
	case parser.NDec:
		return ConvertDecimal(value)
	case parser.NBin:
		return ConvertBin(value)
	case parser.NHex:
		return ConvertHex(value)
	}

	return ConvertExponential(value)
}

// CalculateOperator calculates the result of an operator.
func CalculateOperator(left, right float64, nodeType parser.NodeType) (float64, error) {
	var result float64

	switch nodeType {
	case parser.NAdd:
		result = left + right
	case parser.NSub:
		result = left - right
	case parser.NMult:
		result = left * right
	case parser.NDiv:
		if right == 0 {
			return 0, ErrorDivisionByZero
		}
		result = left / right
	case parser.NMod:
		for {
			if left < right {
				break
			}

			left -= right
		}
		result = float64(left)
	case parser.NOr:
		result = float64(int(left) | int(right))
	case parser.NXor:
		result = float64(int(left) ^ int(right))
	case parser.NAnd:
		result = float64(int(left) & int(right))
	}

	return result, nil
}

// CalculateFunction calculates the result of a function.
func CalculateFunction(arg float64, nodeType parser.NodeType) (float64, error) {
	var result float64

	switch nodeType {
	case parser.NFnSqrt:
		result = math.Sqrt(arg)
	case parser.NFnSin:
		result = math.Sin(arg)
	case parser.NFnCos:
		result = math.Cos(arg)
	case parser.NFnTan:
		result = math.Tan(arg)
	case parser.NFnAbs:
		result = math.Abs(arg)
	case parser.NFnSignbit:
		if math.Signbit(arg) {
			result = 1
		} else {
			result = 0
		}
	case parser.NFnCeil:
		result = math.Ceil(arg)
	case parser.NFnFloor:
		result = math.Floor(arg)
	case parser.NFnTrunc:
		result = math.Trunc(arg)
	case parser.NFnCbrt:
		result = math.Cbrt(arg)
	case parser.NFnAsin:
		result = math.Asin(arg)
	case parser.NFnAcos:
		result = math.Acos(arg)
	case parser.NFnAtan:
		result = math.Atan(arg)
	case parser.NFnLog:
		result = math.Log(arg)
	case parser.NFnLog2:
		result = math.Log2(arg)
	case parser.NFnLog10:
		result = math.Log10(arg)
	}

	return result, nil
}
