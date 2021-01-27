package easy

import "strings"

// Interpret question
// You own a Goal Parser that can interpret a string command. The command consists of an alphabet of "G", "()" and/or "(al)" in some order. The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". The interpreted strings are then concatenated in the original order.
// Given the string command, return the Goal Parser's interpretation of command.
func Interpret(command string) string {
	return strings.Replace(strings.Replace(command, "()", "o", -1), "(al)", "al", -1)
}
