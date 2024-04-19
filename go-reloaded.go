package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	var sli []string
	var final_product []string
	sli = SplitWhiteSpaces(args)
	for i := 1; i < len(sli); i++ {
		if sli[i] == "(hex)" {
			new := ConvertBase(sli[i-1], 16)
			sli[i-1] = new
		} else if sli[i] == "(bin)" {
			new := ConvertBase(sli[i-1], 2)
			sli[i-1] = new
		} else if sli[i] == "(up)" {
			new := ToUpper(sli[i-1])
			sli[i-1] = new
		} else if sli[i] == "(low)" {
			new := ToLower(sli[i-1])
			sli[i-1] = new
		} else if sli[i] == "(cap)" {
			new := Capitalize(sli[i-1])
			sli[i-1] = new
		} else if sli[i] == "(up," {

			new := sli[i+1]
			num := TrimAtoi(new)
			sli = ToUpperNum(sli, num)
		} else if sli[i] == "(low," {

			new := sli[i+1]
			num := TrimAtoi(new)
			sli = ToLowrNum(sli, num)
		} else if sli[i] == "(cap," {
			new := sli[i+1]
			num := TrimAtoi(new)
			sli = CapNum(sli, num)
		}
	}
	count := 0
	for i := 1; i < len(sli); i++ {
		if sli[i] == "(hex)" || sli[i] == "(bin)" || sli[i] == "(low)" || sli[i] == "(up)" || sli[i] == "(cap)" {
			count++
		} else if sli[i] == "(low," || sli[i] == "(up," || sli[i] == "(cap," {
			count += 2
		}
	}
	for i := 1; i < len(sli); i++ {
		for j := i; j < len(sli)-1; j++ {
			if sli[j] == "(hex)" || sli[j] == "(bin)" || sli[j] == "(low)" || sli[j] == "(up)" || sli[j] == "(cap)" {
				sli[j] = ""
			} else if sli[j] == "(low," || sli[j] == "(up," || sli[j] == "(cap," {
				sli[j] = ""
				sli[j+1] = ""
			} else if sli[j] == sli[j-1] {
				sli[j] = ""
			}
		}
	}
	sli = SplitWhiteSpaces(sli)
	for i := 0; i < len(sli); i++ {
		final_product = append(final_product, sli[i], " ")
	}
	
	for i := 0; i < len(final_product); i++ {
		if i < len(final_product) {
			fmt.Print(final_product[i])
		} else {
			fmt.Print(final_product[i])
		}
	}
	fmt.Println()
}

func SplitWhiteSpaces(s []string) []string {
	var sli []string
	word := ""
	for _, str := range s {
		for _, char := range str {
			if char != ' ' {
				word = word + string(char)
			} else {
				sli = append(sli, word)
				word = ""
				continue
			}
		}
		if word != "" {
			sli = append(sli, word)
			word = ""
		}
	}
	return sli
}

func ConvertBase(s string, l int) string {
	var base string
	if l == 16 {
		base = "0123456789ABCDEF"
	} else {
		base = "01"
	}
	result := 0
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		result = result + strings.IndexByte(base, s[i])*power(l, j)
		j++
	}
	new := strconv.Itoa(result)
	return new
}

func power(x, y int) int {
	pow := 1
	if y > 0 {
		pow = x * power(x, y-1)
	}
	return pow
}

func ToUpper(s string) string {
	word := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			word = word + string(char-32)
		} else {
			word = word + string(char)
		}
	}
	return word
}
func ToLower(s string) string {
	word := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			word = word + string(char+32)
		} else {
			word = word + string(char)
		}
	}
	return word
}
func Capitalize(s string) string {
	word := ""
	for i, char := range s {
		if char >= 'a' && char <= 'z' && i == 0 {
			word = word + string(char-32)
		} else {
			word = word + string(char)
		}
	}
	return word
}
func ToUpperNum(sli []string, n int) []string {
	temp_slice := []string(sli)
	count := 0
	j := 0
	for i := 0; i < len(sli); i++ {
		if sli[i] == "(up," {
			j = i
		}

	}
	for i := j - 1; i > 0; i-- {
		if count < n {
			modify := ToUpper(sli[i])
			temp_slice[i] = modify
			count++
		} else {
			break
		}
	}
	return temp_slice

}

func ToLowrNum(sli []string, n int) []string {
	temp_slice := []string(sli)
	count := 0
	j := 0
	for i := 0; i < len(sli); i++ {
		if sli[i] == "(low," {
			j = i
		}

	}
	for i := j - 1; i > 0; i-- {
		if count < n {
			modify := ToLower(sli[i])
			temp_slice[i] = modify
			count++
		} else {
			break
		}
	}
	return temp_slice

}

func CapNum(sli []string, n int) []string {
	temp_slice := []string(sli)
	count := 0
	j := 0
	for i := 0; i < len(sli); i++ {
		if sli[i] == "(cap," {
			j = i
		}
	}
	for i := j - 1; i > 0; i-- {
		if count < n {
			modify := Capitalize(sli[i])
			temp_slice[i] = modify
			count++
		} else {
			break
		}
	}
	return temp_slice
}

func TrimAtoi(s string) int {
	word := ""
	result := 0
	signe := 1
	for _, char := range s {
		if char == '-' && word == "" {
			signe = -1
			continue
		}
		if char >= '0' && char <= '9' {
			word = word + string(char)
		}
	}
	for _, val := range word {
		result = result*10 + int(val-'0')
	}

	return result * signe
}
