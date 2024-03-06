package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}*/

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 97 && char <= 122 {
			char = (char - 32)
		}
		result = result + string(char)
	}
	return result
}
