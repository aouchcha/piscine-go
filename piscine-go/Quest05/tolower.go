package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}*/

func ToLower(s string) string {
	tolower := ""
	for _, char := range s {
		if char >= 65 && char <= 90 {
			char = char + 32
		}
		tolower += string(char)
	}
	return tolower
}
