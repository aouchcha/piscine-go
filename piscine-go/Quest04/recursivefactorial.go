package piscine

/*import (
	"fmt"
)*/

/*func main() {
	arg := 21
	fmt.Println(RecursiveFactorial(arg))
}*/

func RecursiveFactorial(nb int) int {
	fact := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else if nb > 1 {
		fact = nb * RecursiveFactorial(nb-1)
	}
	return fact
}
