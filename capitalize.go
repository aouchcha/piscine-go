package piscine

func Capitalize(s string) string {
	str := []rune(s)
	for i := 1; i < len(s)-1; i++ {
		if (s[i] == ' ' || s[i] == '+') && s[i+1] >= 'a' && s[i+1] <= 'z' {
			str[i+1] = str[i+1] - 32
		}
	}
	return string(str)
}
