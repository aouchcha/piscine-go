package piscine

func StrRev(s string) string {
	str := []rune(s)
	l := len(s)
	for i := 0; i < len(s); i++ {
		str[i] = rune(s[l-i-1])
	}
	return string(str)
}
