package tools

import (
	"strconv"
)

func convertbase(toconvert string, flag string) string {
	var base int
	if flag == "bin" {
		base = 2
	} else if flag == "hex" {
		base = 16
	}
	convertedbase := toconvert
	s, err := strconv.ParseInt(toconvert, base, 64)
	if err == nil {
		convertedbase = strconv.Itoa(int(s))
	}
	return convertedbase
}
