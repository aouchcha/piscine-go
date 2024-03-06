package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, s := range *ptr {
		if s != "" {
			count++
		}
	}
	i := 0
	arr := make([]string, count)
	for _, s := range *ptr {
		if s != "" {
			arr[i] = s
			i++
		}
	}
	*ptr = arr
	return count
}
