package tasks

import (
	"fmt"
	"strings"
)

func TextToTitle(text string) (string,error) {
	if len(text) <= 0 {
		return "",fmt.Errorf("error: empty text")
	}

	isUpper := func(y rune) bool {
		if y >= 'A' && y <= 'Z' {
			return true
		} else {
			return false
		}
	}
	isLower := func(x rune) bool {
		if x >= 'a' && x <= 'z' {
			return true
		} else {
			return false
		}
	}

	found := false
	for _, v := range text {
		if isUpper(v) || isLower(v) {
			found = true
			break
		}
	}
	if !found {
		return "",fmt.Errorf("error: no letter")
	}
	
	var res string
	var words []string = strings.Split(text," ")
	for _,w := range words {
		for i, v := range w {
			if i == 0 && isLower(v) {
				res += string(v-32)
			} else {
				if isUpper(v) {
					res += string(v+32)
				} else {
					res += string(v)
				}
			}
		}
		res += " "
	}
	return res,nil
}