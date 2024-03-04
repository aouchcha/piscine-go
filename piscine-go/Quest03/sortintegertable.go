package piscine

/*import (
	"fmt"
)*/

/*func main() {
	s := []int{18, 4, -20, 1000, -19785, 0} //[0 1 2 3 4 5]
	SortIntegerTable(s)
	fmt.Println(s)
}*/

func SortIntegerTable(table []int) {
	swap := 0
	for i := 0; i < len(table)-1; i++ { // ghadi nfexiw biha bsseh mankhedmouch lindice dyalha , lgharad mnha howa nfekssiw i bash numero lwahed ytverifa kaml 3ad ndiro i++
		for j := 0; j < len(table)-1; j++ { // we use len(table)-1 because we will have access to tha last case with table[j+1] else we'll have an error because we will have g out of range of the the array
			if table[j] > table[j+1] {
				swap = table[j]
				table[j] = table[j+1]
				table[j+1] = swap
			}
		}
	}
}
