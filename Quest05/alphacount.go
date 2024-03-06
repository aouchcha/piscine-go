package piscine

/*import (
	"fmt"
)*/

/*func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}*/

func AlphaCount(s string) int {
	str := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if str[i] >= 65 && str[i] <= 90 || str[i] >= 97 && str[i] <= 122 {
			count++
		}
	}
	return count
}
