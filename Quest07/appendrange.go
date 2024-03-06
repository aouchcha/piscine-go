package piscine

/*import "fmt"*/

/*unc main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}*/

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var arr []int
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}

	return arr
}
