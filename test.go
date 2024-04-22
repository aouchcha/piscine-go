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
	slice := strings.Split(str, " ")

	for i := 1; i < len(slice); i++ {
		if slice[i] == "(hex)" {
			temp, _ := strconv.ParseInt(slice[i-1], 16, 32)
			slice[i-1] = strconv.Itoa(int(temp))
			slice[i] = ""
		} else if slice[i] == "(bin)" {
			slice[i-1] = ConvertBase(slice[i-1], 2)
			slice[i] = ""
		} else if slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice[i] = ""
		} else if slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
		} else if slice[i] == "(cap)" {
			slice[i-1] = Capitalize(slice[i-1])
			slice[i] = ""
		} else if slice[i] == "(cap," || slice[i] == "(low," || slice[i] == "(up," {
			slice = numericalChanges(slice, i)
		}
	}

	str = strings.Join(slice, " ")
	slice = strings.Fields(str)

	result := []string{}
	for i := 0; i < len(slice); i++ {
		result = append(result, slice[i], " ")
	}


	//str = strings.Trim(str,"")
	str = string(handlPonctuel(result))
	fmt.Println(str)

	last_test := []rune(str)
	//slice = strings.Fields(str)
	word := ""
	for i := 0; i < len(last_test); i++ {
		if last_test[i] == '.' || last_test[i] == ',' || last_test[i] == '!' || last_test[i] == '?' || last_test[i] == ':' || last_test[i] == ';' && last_test[i+1] != ' ' {
			word = word + string(last_test[i]) + " "
		} else {
			word = word + string(last_test[i])
		}
	}
	str = word

	fmt.Println(str)
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
			slice[nbr-j-1] = strings.ToUpper(slice[nbr-j-1])
		}
	} else if slice[nbr] == "(low," {
		for j := 0; j < num; j++ {
			slice[nbr-j-1] = strings.ToLower(slice[nbr-j-1])
		}
	} else if slice[nbr] == "(cap," {
		for j := 0; j < num; j++ {
			slice[nbr-j-1] = Capitalize(slice[nbr-j-1])
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

func handlPonctuel(slice []string) []rune {
	for i := 1; i < len(slice)-1; i++ {
		if slice[i] == " " {
			for _, ele := range slice[i+1] {
				if ele == ',' || ele == '.' || ele == ';' || ele == ':' || ele == '!' || ele == '?' {
					slice[i] = ""
				} else if ele != '.' && ele != ',' && ele != '!' && ele != '?' && ele != ':' && ele != ';' {
					break
				}
			}
		}
	}
	var sliceofrunes []rune
	for _, str := range slice {
		for _, char := range str {
			sliceofrunes = append(sliceofrunes, char)
		}
	}

	for i := 0; i < len(sliceofrunes); i++ {

		if sliceofrunes[i] == '\'' {

			if sliceofrunes[i+1] == ' ' {
				for j := i + 1; j < len(sliceofrunes)-1; j++ {
					sliceofrunes[j] = sliceofrunes[j+1]
				}
			}
			if sliceofrunes[i-1] == ' ' && sliceofrunes[i-2] != ':' {
				for j := i; j < len(sliceofrunes)-1; j++ {
					sliceofrunes[j-1] = sliceofrunes[j]
				}
			}
		}

	}

	return sliceofrunes
}
