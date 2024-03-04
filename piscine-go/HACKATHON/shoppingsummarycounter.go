package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	MP := make(map[string]int)
	word := ""
	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else {
			MP[word]++
			word = ""
		}
	}
	MP[word]++
	return MP
}
