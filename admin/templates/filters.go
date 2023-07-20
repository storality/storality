package templates

import (
	"html/template"
	"unicode"
)

var functions = template.FuncMap{
	"capitalize": capitalize,
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}