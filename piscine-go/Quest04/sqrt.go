package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(Sqrt(-411437))
	fmt.Println(Sqrt(9))
}*/

func Sqrt(nb int) int {
	x := 1
	if nb <= 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		x = i * i
		if x == nb && nb%i == 0 {
			x = i
			break
		} else {
			x = 0
		}
	}
	return x
}
