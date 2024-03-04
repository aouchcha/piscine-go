package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}*/

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		if arr[i] != max {
			arr[i] = min + i
		} else {
			break
		}
	}
	return arr
}
