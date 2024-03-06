package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IterativePower(2, 3))
}*/

func IterativePower(nb int, power int) int {
	result := 1
	if power <= 0 {
		return 0
	}
	for i := 0; i < power; i++ {
		result = nb * result
	}
	return result
}
