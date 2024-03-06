package piscine

func JumpOver(str string) string {
	s := ""
	if len(str) >= 3 {
		for i, char := range str {
			if i%3 == 2 {
				s = s + string(char)
			}
		}
	}
	return s + "\n"
}
