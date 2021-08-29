package util

import (
	"math"
	"strconv"
)

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

func ToDecimalString(precision int, num int) string {
	ret := NewArrayByteFilledWith(len(strconv.Itoa(precision-1)), 48)
	numStr := []byte(strconv.Itoa(num))
	retLength := len(ret)
	for i := 0; i < len(numStr); i++ {
		ret[retLength-1-i] = numStr[i]
	}
	return string(ret)
}
