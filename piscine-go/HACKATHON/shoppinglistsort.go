package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				swap := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = swap
			}
		}
	}
	return slice
}
