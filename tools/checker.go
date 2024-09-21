package tools

import "strings"

func Checktype(name string) bool { // check if its A txt
	is_it_a_txt_file := strings.HasSuffix(name, ".txt")
	return is_it_a_txt_file
}
