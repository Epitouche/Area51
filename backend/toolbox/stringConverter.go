package toolbox

import (
	"fmt"
	"strconv"
)

type StringConverter interface {
	StringToFloat64(string) (float64, error)
	StringToBoolean(string) (bool, error)
}

func StringToFloat64(str string) (float64, error) {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: '%s' cannot be converted to float64", str)
	}
	return result, nil

}

func StringToBoolean(str string) (bool, error) {
	result, err := strconv.ParseBool(str)
	if err != nil {
		return false, fmt.Errorf("invalid input: '%s' cannot be converted to boolean", str)
	}
	return result, nil
}
