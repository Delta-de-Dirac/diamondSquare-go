package utils

import "strings"

func IsPowerOf2(x int) bool {
	if x <= 0{
		return false
	}
	for x % 2 == 0{
		x /= 2
		if x == 1{
			return true
		}
	}
	return false
}

func FilterString(s string, r string) string{
	out := ""
	for _, v := range(s){
		if strings.ContainsRune(r, v){
			out += string(v)
		}
	}
	return out
}

func Min(x int32, y ...int32) int32{
	min := x
	for _, v := range(y){
		if v < min{
			min = v
		}
	}
	return min
}
