package piscine

/*
import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", SplitWhiteSpaces("2gf+{,'\\[x*~T [/9bjYh@p)SN  g)xk`rAgbzjry \"i;C$D-M2?&Wo cO2&exWE`?G;w P)@~5fLvHgC*L ^'`]JEeQz~#\"` A=P8U^CG&6bbj"))
}
*/

func SplitWhiteSpaces(s string) []string {
	sli := []string{}
	word := ""
	for _, char := range s {
		if char != ' ' {
			word = word + string(char)
		} else if char == ' ' && word != "" {
			sli = append(sli, word)
			word = ""
		}
	}
	if word != "" {
		sli = append(sli, word)
		word = ""
	}
	return sli
}
