package transforms

import (
	"strings"
	"unicode"
)

func Slugify(name string) string {
	name = strings.ToLower(name)
	slug := ""
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			slug += string(r)
		} else if unicode.IsSpace(r) {
			slug += "-"
		}
	}
	slug = strings.ReplaceAll(slug, "--", "-")
	slug = strings.Trim(slug, "-")
	return slug
}