package piscine

func ReverseMenuIndex(menu []string) []string {
	sli := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		sli[len(menu)-1-i] += menu[i]
	}
	return sli
}
