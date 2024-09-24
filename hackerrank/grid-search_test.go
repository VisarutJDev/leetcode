package hackerrank

import (
	"fmt"
	"testing"
)

func TestGridSearch(t *testing.T) {
	fmt.Println(gridSearch(
		[]string{
			"111111111111111",
			"111111111111111",
			"111111111111111",
			"111111011111111",
			"111111111111111",
			"111111111111111",
			"101010101010101",
		},
		[]string{
			"11111",
			"11111",
			"11111",
			"11110",
		},
	))
	fmt.Println(gridSearch(
		[]string{
			"7283455864",
			"6731158619",
			"8988242643",
			"3830589324",
			"2229505813",
			"5633845374",
			"6473530293",
			"7053106601",
			"0834282956",
			"4607924137",
		},
		[]string{
			"9505",
			"3845",
			"3530",
		},
	))
	fmt.Println(gridSearch(
		[]string{
			"456712",
			"561245",
			"123667",
			"781288",
		},
		[]string{
			"45",
			"67",
		},
	))
	fmt.Println(gridSearch(
		[]string{
			"123456",
			"567890",
			"234567",
			"194729",
		},
		[]string{
			"1234",
			"5678",
			"2345",
			"4729",
		},
	))
}
