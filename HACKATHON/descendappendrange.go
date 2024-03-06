package piscine

func DescendAppendRange(max, min int) []int {
	sli := []int{}
	if max <= min {
		return sli
	} else {
		for i := max; i > min; i-- {
			sli = append(sli, i)
		}
	}
	return sli
}
