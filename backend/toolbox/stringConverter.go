package toolbox

import (
	"fmt"
	"strconv"
)

type StringConverter interface {
	StringToFloat64(string) (float64, error)
}

func StringToFloat64(str string) (float64, error) {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: '%s' cannot be converted to float64", str)
	}
	return result, nil

}
