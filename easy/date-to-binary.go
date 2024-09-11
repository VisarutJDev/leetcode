package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	// split string becuase we have dash in string
	// So we will got 3 value in array
	arrStr := strings.Split(date, "-")
	result := ""
	for i := range arrStr {
		number, err := strconv.Atoi(arrStr[i])
		if err != nil {
			// ... handle error
			panic(err)
		}
		result += fmt.Sprintf("%b", number)
		if i != len(arrStr)-1 {
			result += "-"
		}
	}
	return result
}
