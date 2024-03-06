package piscine

/*import (
	"fmt"
)*/

/*func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}*/

func ConcatParams(args []string) string {
	str := ""
	for i, arg := range args {
		if i < len(args) {
			str = str + arg
		}
		if i < len(args)-1 {
			str = str + "\n"
		}
	}
	return str
}
