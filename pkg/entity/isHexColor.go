package entity

import "regexp"

func IsHex(h string) bool {
	color := regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
	return color.MatchString(h)
}
