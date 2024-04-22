package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)
	sli := strings.Split(str, " ")

	for i := 1; i < len(sli); i++ {
		if sli[i] == "(hex)" {
			new := ConvertBase(sli[i-1], 16)
			sli[i-1] = new
			sli[i] = ""

		} else if sli[i] == "(bin)" {
			new := ConvertBase(sli[i-1], 2)
			sli[i-1] = new
			sli[i] = ""

		} else if sli[i] == "(up)" {
			new := ToUpper(sli[i-1])
			sli[i-1] = new
			sli[i] = ""

		} else if sli[i] == "(low)" {
			new := ToLower(sli[i-1])
			sli[i-1] = new
			sli[i] = ""
		} else if sli[i] == "(cap)" {
			new := Capitalize(sli[i-1])
			sli[i-1] = new
			sli[i] = ""
		} else if sli[i] == "(up," {
			new := sli[i+1]
			num := TrimAtoi(new)
			sli = ToUpperNum(sli, num)
			sli[i] = ""
			sli[i+1] = ""
		} else if sli[i] == "(low," {
			new := sli[i+1]
			num := TrimAtoi(new)
			sli = ToLowrNum(sli, num)
			sli[i] = ""
			sli[i+1] = ""
		} else if sli[i] == "(cap," {
			new := sli[i+1]
			num := TrimAtoi(new)
			sli = CapNum(sli, num)
			sli[i] = ""
			sli[i+1] = ""
		}
	}
	//fmt.Println(sli)
	sli = splitWhitSpacesInSliceOfStrings(sli)

	//Add whitSpaces
	sli = addWhithSpaces(sli)

	sli = handelPonctuel(sli)

	sli = Handl_Voyel(sli)

	var final_result string
	final_result = joinSliceOfStrings(sli)

	// handel the case when the "," is in the middel of the word
	final_test := []rune(final_result)
	var slice []rune
	for i := 0; i < len(final_test)-1; i++ {
		if (final_test[i] == '.' || final_test[i] == ',' || final_test[i] == '!' || final_test[i] == '?' || final_test[i] == ':' || final_test[i] == ';') && (final_test[i+1] != '.' && final_test[i+1] != ',' && final_test[i+1] != '!' && final_test[i+1] != '?' && final_test[i+1] != ':' && final_test[i+1] != ';' && final_test[i] != ' ') {
			slice = append(slice, final_test[i], ' ')
		} else {
			slice = append(slice, final_test[i])
		}
	}

	//transform slice of runes into string
	final_result = string(slice)

	sli = splitWhitSpacesOfStrings(final_result)
	sli = splitWhitSpacesInSliceOfStrings(sli)
	//fmt.Println(sli)
	sli = Handl_singlequotes(sli)

	//transform slice of strings into slice
	final_result = joinSliceOfStrings(sli)

	//handl the case when we have the "'" in the first element
	slice = []rune{}
	for i := 0; i < len(final_result); i++ {
		if final_result[i] == '\'' && final_result[i-1] == ':' {
			slice = append(slice, ' ', rune(final_result[i]))
		} else {
			slice = append(slice, rune(final_result[i]))
		}
	}

	//transform slice of runes into string
	final_result = string(slice)

	fmt.Println(final_result)

}

func splitWhitSpacesInSliceOfStrings(s []string) []string {
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

func splitWhitSpacesOfStrings(s string) []string {
	var sli []string
	word := ""

	for _, char := range s {
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

	return sli
}

func addWhithSpaces(sli []string) []string {
	var Product []string
	for i := 0; i < len(sli); i++ {
		if i < len(sli) {
			Product = append(Product, sli[i], " ")
		} else {
			Product = append(Product, sli[i])
		}
	}
	return Product
}

func joinSliceOfStrings(sli []string) string {
	word := ""
	for _, str := range sli {
		word = word + string(str)
	}
	return word
}

func ConvertBase(s string, l int) string {
	var base string
	if l == 16 {
		base = "0123456789abcdef"
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
	for i := j - 1; i >= 0; i-- {
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
	for i := j - 1; i >= 0; i-- {
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
	for i := j - 1; i >= 0; i-- {
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

func handelPonctuel(sli []string) []string {
	for i := 0; i < len(sli)-1; i++ {
		if sli[i] == " " {
			for _, char := range sli[i+1] {
				if char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';' {
					sli[i] = ""
				} else if char != '.' && char != ',' && char != '!' && char != '?' && char != ':' && char != ';' {
					break
				}
			}
		}
	}
	return sli
}

func Handl_singlequotes(finish []string) []string {
	Product := []string{}
	for i := 1; i < len(finish); i++ {
		if finish[i] == "'" || finish[i-1] == "'" {
			Product = append(Product, finish[i])
		} else {
			Product = append(Product, " ", finish[i])
		}
	}

	return Product
}
func Handl_Voyel(sli []string) []string {
	for i := 0; i < len(sli)-2; i++ {
		if sli[i] == "a" {
			for j, char := range sli[i+2] {
				if char == 'a' || char == 'e' || char == 'u' || char == 'o' || char == 'i' || char == 'A' || char == 'E' || char == 'U' || char == 'O' || char == 'I' && j == 0 {
					sli[i] = "an"
				} else {
					break
				}
			}
		}

		if sli[i] == "A" {
			for j, char := range sli[i+2] {
				if char == 'a' || char == 'e' || char == 'u' || char == 'o' || char == 'i' || char == 'A' || char == 'E' || char == 'U' || char == 'O' || char == 'I' && j == 0 {
					sli[i] = "AN"
				} else {
					break
				}
			}
		}
	}
	return sli
}
