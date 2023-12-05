package utils

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
