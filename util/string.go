package util

import "golang.org/x/text/width"

func ToLowerString(str string) string {
	return width.Narrow.String(str)
}
