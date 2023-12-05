package library

import (
	"fmt"
	"math"
	"strconv"
)

// interface{} to string
func ItString(value interface{}) string {
	if value == nil {
		return ""
	} else {
		data := fmt.Sprintf("%v", value)
		return data
	}

}

// interface{} to interface{}(integer)
func ItInt(i interface{}) int {
	value, err := strconv.Atoi(ItString(i))
	if err != nil {
		return 0
	}

	return value
}

// (*string to string)
func PointStringToString(value *string) string {
	var value_return string
	if value == nil {
		value_return = ""
	} else {
		valuestar := *value
		value_return = valuestar
	}
	return value_return
}

// (*string to *string)
func PointStringToPointString(value *string, insert ...string) *string {
	var value_return *string
	if value == nil {
		value_return = nil
	} else if len(insert) > 0 {
		for i := 0; i < len(insert); i++ {
			if i > 0 {
				value_next := *value_return
				valuenotnull := insert[i] + value_next
				value_return = &valuenotnull
			} else {
				valuestar := *value
				valuenotnull := insert[i] + valuestar
				value_return = &valuenotnull
			}
		}
	} else if len(insert) == 0 {
		valuestar := *value
		valuenotnull := valuestar
		value_return = &valuenotnull
	}
	return value_return
}

// https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
