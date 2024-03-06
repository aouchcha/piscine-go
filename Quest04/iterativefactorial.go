package piscine

/*import (
	"fmt"
)*/

func IterativeFactorial(nb int) int {
	fact := 1
	if nb == 0 || nb == 1 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	} else {
		for i := nb; i >= 1; i-- {
			fact *= i
		}
	}
	return fact
}

/*func main() {
	arg := 20
	fmt.Println(IterativeFactorial(arg))
}*/
