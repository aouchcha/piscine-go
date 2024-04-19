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
	word := ""
	for _, str := range args {
		for _, char := range str {
			if char != ' ' {
				word = word + string(char)
			} else {
				sli = append(sli, word)
				word = ""
			}
		}
		if word != "" {
			sli = append(sli, word)
			word = ""
		}
	}
	cc := ""
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
			cc = sli[i+1]
			new := sli[i+1]
			num := TrimAtoi(new)
			sli = ToUpperNum(sli, num)
		}
	}
	fmt.Println("cc is", cc)
	count := 0
	for i := 1; i < len(sli); i++ {
		if sli[i] == "(hex)" || sli[i] == "(bin)" || sli[i] == "(low)" || sli[i] == "(up)" || sli[i] == "(cap)" || sli[i] == "(low, " || sli[i] == "(up, " || sli[i] == "(cap, " {
			count++
		}
	}
	for i := 1; i < len(sli); i++ {
		for j := i; j < len(sli)-1; j++ {
			if sli[j] == "(hex)" || sli[j] == "(bin)" || sli[j] == "(low)" || sli[j] == "(up)" || sli[j] == "(cap)" {
				sli[j] = sli[j+1]
			} else if sli[j] == "(low," || sli[j] == "(up," || sli[j] == "(cap," {
				sli[j] = sli[j+1]
			} else if sli[j] == sli[j-1] {
				sli[j] = sli[j+1]
			}
		}
	}

	for i := 0; i < len(sli)-count; i++ {
		if i < len(sli) {
			fmt.Print(sli[i], " ")
		} else {
			fmt.Print(sli[i])
		}
	}
	fmt.Println()
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
	for i := len(sli) - 3; i > 0; i-- {
		if count < n {
			modify := ToUpper(sli[i])
			temp_slice[i] = modify
			count++
		} else {
			break
		}
	}
	// fmt.Println(temp_slice)
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
