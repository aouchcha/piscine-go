package piscine

/*import (
	"os"

	//"github.com/01-edu/z01"
	"fmt"
)*/

/*func main() {
	const (
		R, C = 9, 9
	)

	args := os.Args[1:]
	answer := ""
	// arg := []string(args)
	// Check Errors!
	for _, arg := range args {
		for range arg {
			if len(arg) != 9 {
				answer = "Error"
				fmt.Println(answer)
				return
			}
		}
	}
	// Remplir the table
	var table [R][C]int
	for i := 0; i < R; i++ {
		for j, char := range args[i] {
			if char >= '1' && char <= '9' {
				table[i][j] = int(char - '0')
			} else {
				table[i][j] = 0
			}
		}
	}
	// fmt.Println(table[0][1])

	// Check Empty cellule

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if table[i][j] == 0 {
			} else {
				continue
			}
		}
	}
}*/
