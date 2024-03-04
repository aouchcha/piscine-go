package piscine

func Max(a []int) int {
	max := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			max = a[i]
		}
	}
	return max
}
