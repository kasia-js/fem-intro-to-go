package utils

import "strings"

//Transforms sentence to upper case letters with ! at the end
func MakeExcited(sentence string) string {
	return strings.ToUpper(sentence) + "!"
}
