package easy

import (
	"fmt"
	"testing"
)

func TestConvertDateToBinary(t *testing.T) {
	fmt.Println(convertDateToBinary("2080-02-29"))
	fmt.Println(convertDateToBinary("1900-01-01"))
}
