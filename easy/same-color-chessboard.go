package easy

import (
	"strconv"
	"strings"
)

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	// a c e g if its odd number it will be black
	// a c e g if its even number it will be white
	// b d f h if its even number it will be white
	// b d f h if its odd number it will be black
	alpha1 := []rune(coordinate1)[0]
	alpha2 := []rune(coordinate2)[0]
	number1, _ := strconv.Atoi(string([]rune(coordinate1)[1]))
	number2, _ := strconv.Atoi(string([]rune(coordinate2)[1]))

	colorOfC1 := checkCharactor(string(alpha1), number1)
	colorOfC2 := checkCharactor(string(alpha2), number2)

	return colorOfC1 == colorOfC2
}

func checkCharactor(alpha string, number int) string {
	IsEven := true
	IsFirstAlpha := true
	y := "b d f h"
	if number%2 != 0 {
		//even number
		IsEven = false
	}
	if strings.Contains(y, alpha) {
		IsFirstAlpha = false
	}
	return checkColor(IsEven, IsFirstAlpha)
}

func checkColor(IsFirstAlpha, IsEven bool) string {
	if IsFirstAlpha {
		if IsEven {
			return "w"
		} else {
			return "b"
		}
	} else {
		if IsEven {
			return "b"
		} else {
			return "w"
		}
	}
}
