package piscine

func Join(strs []string, sep string) string {
	word := ""
	for i, s := range strs {
		if i < len(strs)-1 {
			word = word + string(s) + sep
		} else {
			word = word + string(s)
		}
	}
	return word
}
