package piscine

/*import (
	"fmt"
)*/

/*func main() {
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(4))
}*/

func IsPrime(nb int) bool {
	reponse := false
	if nb <= 0 || nb == 1 {
		return false
	} else if nb == 2 {
		reponse = true
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			reponse = false
			break
		} else {
			reponse = true
		}
	}

	return reponse
}
