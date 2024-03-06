package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}*/

func Index(s string, toFind string) int {
	index := 0
	str1 := []rune(s)
	str2 := []rune(toFind)
	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if str1[i] == str2[0] {
			index = i
			break
		} else {
			index = -1
		}
	}
	return index
}
