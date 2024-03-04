package piscine

/*
import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/

func IsSorted(f func(a, b int) int, a []int) bool {
	resultup := true
	resultdown := true
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			resultup = false
		}
		if f(a[i], a[i+1]) < 0 {
			resultdown = false
		}
	}
	return resultup || resultdown
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
