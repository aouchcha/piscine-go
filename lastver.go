package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("You didn't enter the two important argumments in the terminal")
		return
	}

	InputFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error !!")
		return
	}

	file := string(InputFile)
	slice := strings.Fields(file)
	slice = AddWhiteSpaces(slice)
	slice = HandlVowels(slice)
	file = ConvertSliceintoString(slice)
	slice = strings.Fields(file)
	if slice[0] == "(hex)" || slice[0] == "(bin)" || slice[0] == "(up)" || slice[0] == "(cap)" || slice[0] == "(low)" {
		fmt.Println()
		fmt.Println("You give the program a commade to make a change in the first element check it")
		fmt.Println()
		return
	} else if slice[0] == "(up," || slice[0] == "(low," || slice[0] == "(cap," {
		fmt.Println()
		fmt.Println("You give the program a commande to make a numerical change in the first element check it")
		fmt.Println()
		return
	}

	slice = DealWithMarkers(slice)

	slice = DeletEmptyCellules(slice)

	slice = AddWhiteSpaces(slice)

	slice = HandlPonctuation(slice)

	file = HandlSingleQuotes(slice)

	file = strings.Trim(file, " ")

	OutputFile := os.Args[2]
	FinalResult := []byte(file)
	err = ioutil.WriteFile(OutputFile, FinalResult, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Go check the result file the progrqmme worked secssefuly !!!")
}

func DealWithMarkers(slice []string) []string {

	for i := 1; i < len(slice); i++ {
		if slice[i] == "(bin)" || slice[i] == "(hex)" {
			if IsBinaireOrHexa(slice[i], slice[i-1]) {
				if slice[i] == "(bin)" {
					num, _ := (strconv.ParseInt(slice[i-1], 2, 64))
					slice[i-1] = strconv.Itoa(int(num))
					slice[i] = ""
				} else if slice[i] == "(hex)" {
					num, _ := (strconv.ParseInt(slice[i-1], 16, 64))
					slice[i-1] = strconv.Itoa(int(num))
					slice[i] = ""
				}
			} else {
				fmt.Println()
				fmt.Println("The number that you want to transforme isn't binaire or not hex please check again !!!")

				fmt.Println("Check the result  it will show you the message thet you entered and modify it !!!!!!!!")
				fmt.Println()
				break
			}
		} else if slice[i] == "(up)" {
			if IsItValideWord(slice[i], slice[i-1]) {
				slice[i-1] = strings.ToUpper(slice[i-1])
				slice[i] = ""
			} else {
				fmt.Println()
				fmt.Println("The word that you want to upper contain a number")

				fmt.Println("Check the result  it will show you the message thet you entered and modify it !!!!!!!!")
				fmt.Println()
				break
			}
		} else if slice[i] == "(low)" {
			if IsItValideWord(slice[i], slice[i-1]) {
				slice[i-1] = strings.ToLower(slice[i-1])
				slice[i] = ""
			} else {
				fmt.Println()
				fmt.Println("The word that you want to lower contain a number")

				fmt.Println("Check the result  it will show you the message thet you entered and modify it !!!!!!!!")
				fmt.Println()
				break
			}
		} else if slice[i] == "(cap)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			if IsItValideWord(slice[i], slice[i-1]) {
				slice[i-1] = Capitalize(slice[i-1])
				slice[i] = ""
			} else {
				fmt.Println()
				fmt.Println("The word that you want to capitalize start with somting insteqd of letter")

				fmt.Println("Check the result  it will show you the message thet you entered and modify it !!!!!!!!")
				fmt.Println()
				break
			}
		} else if slice[i] == "(up," || slice[i] == "(cap," || slice[i] == "(low," {
			if IsItPossible(slice, slice[i], i) {
				num := TrimAtoi(slice[i+1])
				if slice[i] == "(up," {

					for j := 0; j < num; j++ {

						if slice[i-j-1] != "" {
							slice[i-j-1] = strings.ToUpper(slice[i-j-1])
						}
					}
				} else if slice[i] == "(low," {
					for j := 0; j < num; j++ {
						if slice[i-j-1] != "" {
							slice[i-j-1] = strings.ToLower(slice[i-j-1])
						}
					}
				} else if slice[i] == "(cap," {
					for j := 0; j < num; j++ {
						if slice[i-j-1] != "" {
							slice[i-j-1] = Capitalize(slice[i-j-1])
						}

					}
				}

				slice[i] = ""
				slice[i+1] = ""

			} else {
				fmt.Println()
				fmt.Println("We can't make the changes because either the number of words you wanna to change is greater than what we have or you didn't enter a number .")

				fmt.Println("Check the result  it will show you the message thet you entered and modify it !!!!!!!!")
				fmt.Println()
				break
			}
		}
	}
	return slice

}

func HandlSingleQuotes(slice []string) string {
	var temp []string
	var slirune []rune
	for _, str := range slice {
		for _, char := range str {
			if char == '\'' {
				slirune = append(slirune, ' ', char, ' ')
			} else {
				slirune = append(slirune, char)
			}
		}
	}
	word := string(slirune)
	slice = strings.Fields(word)
	for i := 0; i < len(slice)-1; i++ {
		if slice[i][len(slice[i])-1] == ':' {
			temp = append(temp, slice[i], " ")
		} else if slice[i+1] == "'" || slice[i] == "'" {
			temp = append(temp, slice[i])
		} else {
			temp = append(temp, slice[i], " ")
		}
	}
	if slice[len(slice)-1] == "'" {
		temp = append(temp, slice[len(slice)-1])
	} else {
		temp = append(temp, slice[len(slice)-1])
	}
	word = ConvertSliceintoString(temp)

	return word
}

func ConvertSliceintoString(slice []string) string {
	word := ""
	for _, str := range slice {
		for _, char := range str {
			word = word + string(char)
		}
	}
	return word
}

func AddWhiteSpaces(slice []string) []string {
	var temp []string
	for i := range slice {
		temp = append(temp, slice[i], " ")
	}
	return temp
}

func HandlPonctuation(slice []string) []string {

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == " " {
			if IsTheWordPonctuation(slice[i+1]) {
				for j := i; j < len(slice)-1; j++ {
					slice[j] = slice[j+1]
				}
			}
		}
	}
	word := ConvertSliceintoString((slice))

	//Handle the case when the ponctuation is in the middle or liee with the word
	temp := []rune(word)
	var ToReturn []rune
	for i := 0; i < len(temp)-1; i++ {
		if (temp[i+1] == ',' || temp[i+1] == '.' || temp[i+1] == ';' || temp[i+1] == ':' || temp[i+1] == '!' || temp[i+1] == '?') && temp[i] == ' ' {
			for j := i; j < len(temp)-1; j++ {
				temp[j] = temp[j+1]
			}
			ToReturn = append(ToReturn, temp[i], ' ')
		} else {
			ToReturn = append(ToReturn, temp[i])
		}
	}
	if temp[len(temp)-1] != ToReturn[len(ToReturn)-1] {
		ToReturn = append(ToReturn, temp[len(temp)-1])
	}
	slice = AddWhiteSpaces(strings.Fields(string(ToReturn)))
	return slice
}

func IsTheWordPonctuation(str string) bool {
	for _, char := range str {
		if char == ',' || char == '.' || char == ';' || char == ':' || char == '!' || char == '?' {
			continue
		} else {
			return false
		}
	}
	return true
}

func HandlVowels(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i+1][0] == 'a' || slice[i+1][0] == 'e' || slice[i+1][0] == 'u' || slice[i+1][0] == 'i' || slice[i+1][0] == 'o' || slice[i+1][0] == 'h' || slice[i+1][0] == 'A' || slice[i+1][0] == 'E' || slice[i+1][0] == 'U' || slice[i+1][0] == 'I' || slice[i+1][0] == 'O' || slice[i+1][0] == 'H' {
			if slice[i] == "a" {
				slice[i] = "an"
			} else if slice[i] == "A" {
				slice[i] = "AN"
			}
		}
	}
	return AddWhiteSpaces(slice)
}

func IsBinaireOrHexa(Marker, str string) bool {
	reponse := true
	if Marker == "(bin)" {
		for _, char := range str {
			if char == '0' || char == '1' {
				reponse = true
			} else {
				reponse = false
				break
			}
		}
	} else if Marker == "(hex)" {
		for _, char := range str {
			if char < '0' || char > '9' && char < 'a' || char > 'f' {
				reponse = false
				break
			}
		}
	}
	return reponse
}

func DeletEmptyCellules(slice []string) []string {
	var temp []string
	for i := range slice {
		if slice[i] != "" {
			temp = append(temp, slice[i])
		}
	}
	return temp
}

func IsItValideWord(Marker, str string) bool {
	if Marker == "(up)" || Marker == "(up," || Marker == "(low)" || Marker == "(low," {
		for _, char := range str {
			if char >= '0' && char <= '9' {
				return false
			}
		}
	} else if Marker == "(cap)" || Marker == "(cap," {
		if str[0] < 'a' || str[0] > 'z' {
			return false
		}
	}
	return true
}

func Capitalize(str string) string {
	temp := []rune(str)
	if temp[0] >= 'a' && temp[0] <= 'z' {
		temp[0] = temp[0] - 32
	}
	return string(temp)
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

func IsItPossible(slice []string, Marker string, index int) bool {
	num := TrimAtoi(slice[index+1])
	if num <= 0 || num > len(slice)-(len(slice)-index) {
		return false
	}
	return true
}
