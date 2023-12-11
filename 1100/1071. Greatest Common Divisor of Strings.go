func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	index := GCD(len(str1), len(str2))
	return str1[:index]
}

func GCD(a int, b int) int {
	if a < b {
		temp := a
		a = b
		b = temp
	}
	for {
		temp := a % b
		a = b
		b = temp
		if temp == 0 {
			break
		}
	}
	return a
}