package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// RemoveDecimalPoint removes the decimal point from a float number with 2 decimal and returns the resulting integer
func RemoveDecimalPoint(number float64) (int, error) {
	numberStr := fmt.Sprintf("%.2f", number)

	numberStr = strings.Replace(numberStr, ".", "", -1)

	result, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// AddDecimalPoint adds a decimal point to an integer and returns the resulting float number with 2 decimal
func AddDecimalPoint(number int) (float64, error) {
	numberStr := fmt.Sprintf("%d", number)

	if len(numberStr) < 2 {
		numberStr = "0" + numberStr
	}

	result, err := strconv.ParseFloat(numberStr[:len(numberStr)-2]+"."+numberStr[len(numberStr)-2:], 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
