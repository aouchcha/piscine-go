package piscine

import (
	"github.com/01-edu/z01"
)

/*func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}*/

func PrintNbrInOrder(n int) {
	var arr []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap := arr[i]
				arr[i] = arr[j]
				arr[j] = swap
			}
		}
	}
	for k := 0; k < len(arr); k++ {
		z01.PrintRune(rune(arr[k] + '0'))
	}
}
