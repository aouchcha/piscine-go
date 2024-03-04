package main

import "github.com/01-edu/z01"

func main() {
	DescendComb()
}

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for z := '8'; z >= '0'; z-- {
					if i == k && j <= z || k > i {
						continue
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(z)
						if j != '0' && j != '1' && k != '0' && z != '0' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
