package piscine

func StringToIntSlice(str string) []int {
	var sli []int
	for _, char := range str {
		sli = append(sli, int(char))
	}
	return sli
}
