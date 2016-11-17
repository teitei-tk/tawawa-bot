package util

import "regexp"

var re, _ = regexp.Compile(`(月曜朝の社畜諸兄にたわわをお届けします|月曜日のたわわ) その([0-9\.]*)`)

func IsTawawaString(text string) bool {
	return re.MatchString(ToLowerString(text))
}
