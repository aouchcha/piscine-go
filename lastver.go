package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("There is an error in the commande line .")
		return
	}

	InputFile := os.Args[1]

	data, err := ioutil.ReadFile(InputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)
	sli := strings.Split(str, "\n")
	str = joinSliceOfStrings(sli)
	sli = strings.Split(str, " ")


	for i := 1; i < len(sli); i++ {
		if sli[0] == "(hex)" || sli[0] == "(bin)" || sli[0] == "(cap)" || sli[0] == "(low)" || sli[0] == "(up)" || sli[0] == "(cap," || sli[0] == "(low," || sli[0] == "(up," {
			sli[0] = ""
		} else if sli[i] == "(hex)" {
			temp, _ := strconv.ParseInt(sli[i-1], 16, 32)
			sli[i-1] = strconv.Itoa(int(temp))
			sli[i] = ""
		} else if sli[i] == "(bin)" {
			temp, _ := strconv.ParseInt(sli[i-1], 2, 32)
			sli[i-1] = strconv.Itoa(int(temp))
			sli[i] = ""
		} else if sli[i] == "(up)" {
			sli[i-1] = strings.ToUpper(sli[i-1])
			sli[i] = ""
		} else if sli[i] == "(low)" {
			sli[i-1] = strings.ToLower(sli[i-1])
			sli[i] = ""
		} else if sli[i] == "(cap)" {
			sli[i-1] = Capitalize(sli[i-1])
			sli[i] = ""
		} else if sli[i] == "(cap," || sli[i] == "(low," || sli[i] == "(up," {
			sli = numericalChanges(sli, i)
		}
	}
	sli = splitWhitSpacesInSliceOfStrings(sli)

	//Add whitSpaces
	sli = addWhithSpaces(sli)
	sli = handelPonctuel(sli)
	sli = Handl_Voyel(sli)

	// handel the case when the "," is in the middel of the word
	var final_result string
	final_result = joinSliceOfStrings(sli)
	final_test := []rune(final_result)
	var slice []rune
	for i := 0; i < len(final_test)-1; i++ {
		if (final_test[i] == '.' || final_test[i] == ',' || final_test[i] == '!' || final_test[i] == '?' || final_test[i] == ':' || final_test[i] == ';') /*&& (final_test[i+1] != '.' && final_test[i+1] != ',' && final_test[i+1] != '!' && final_test[i+1] != '?' && final_test[i+1] != ':' && final_test[i+1] != ';' && final_test[i] != ' ') */{
			slice = append(slice, final_test[i], ' ')
		} else {
			slice = append(slice, final_test[i])
		}
	}

	//transform slice of runes into string
	final_result = string(slice)
	sli = strings.Split(final_result, " ")
	sli = splitWhitSpacesInSliceOfStrings(sli)

	//slice = []rune{}
	count := 0
	slice = Handl_singlequotes(sli, &count)

	// //transform slice of runes into string
	final_result = string(slice)
	fmt.Println(final_result)

	//transform slice of runes into string
	final_result = ""
	for i := 0; i < len(slice)-count; i++ {
		final_result = final_result + string(slice[i])
	}

	ResultFile := os.Args[2]
	Final := []byte(final_result)
	err = ioutil.WriteFile(ResultFile, Final, 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

}

func splitWhitSpacesInSliceOfStrings(s []string) []string {
	var sli []string
	word := ""
	for _, str := range s {
		for _, char := range str {
			if char != ' ' {
				word = word + string(char)
			} else {
				sli = append(sli, word," ")
				word = ""
			}
		}
		if word != "" {
			sli = append(sli, word)
			word = ""
		}
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
func numericalChanges(slice []string, nbr int) []string {
	num := TrimAtoi(slice[nbr+1])
	if slice[nbr] == "(up," {
		for j := 0; j < num; j++ {
			slice[nbr-j] = strings.ToUpper(slice[nbr-j])
		}
	} else if slice[nbr] == "(low," {
		for j := 0; j < num; j++ {
			slice[nbr-j] = strings.ToLower(slice[nbr-j])
		}
	} else if slice[nbr] == "(cap," {
		for j := 0; j < num; j++ {
			slice[nbr-j] = Capitalize(slice[nbr-j])
		}
	}
	slice[nbr] = ""
	slice[nbr+1] = ""
	return slice
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

func Handl_singlequotes(finish []string, count *int) []rune {
	Product := []string{}
	Product = append(Product, finish[0])
	for i := 1; i < len(finish); i++ {
		if finish[i] == "'" || finish[i-1] == "'" {
			Product = append(Product, finish[i])
		} else {
			Product = append(Product, " ", finish[i])
		}
	}

	final_result := joinSliceOfStrings(Product)
	slice := []rune{}
	for i := 0; i < len(final_result); i++ {
		if final_result[i] == '\'' && final_result[i-1] == ':' {
			slice = append(slice, ' ', rune(final_result[i]))
		} else {
			slice = append(slice, rune(final_result[i]))
		}
	}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == ' ' && slice[i+1] == '\'' && slice[i-1] != ':' {
			for j := i; j < len(slice)-1; j++ {
				slice[j] = slice[j+1]
			}
			*count++
		} else {
			continue
		}
	}

	return slice
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
