package main

import(
	"fmt"
	"strings"
	"os"
	"strconv"
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
	for i:=1; i< len(sli); i++ {
		if sli[i] == "(hex)" {
			new := ConvertBase(sli[i-1],16)
			sli[i-1] = new
		} else if sli[i] == "(bin)" {
			new := ConvertBase(sli[i-1],2)
			sli[i-1] = new
		}else if sli[i] == "(up)" {
			new := ToUpper(sli[i-1])
			sli[i-1] = new
		}else if sli[i] == "(low)" {
			new := ToLower(sli[i-1])
			sli[i-1] = new
		}else if sli[i] == "(cap)" {
			new := Capitalize(sli[i-1])
			sli[i-1] = new
		}else if sli[i] == "(up," {
			new := ToLowerNum(sli[i+1])
		}
	}
	count := 0
	for i := 1; i < len(sli); i++ {
		if sli[i] == "(hex)" || sli[i] == "(bin)" || sli[i] == "(low)" || sli[i] == "(up)" || sli[i] == "(cap)" || sli[i] == "(low, " || sli[i] == "(up, " || sli[i] == "(cap, " {
			count++
		} 
	}
	for i := 1; i < len(sli); i++ {
		for j := i; j < len(sli)-1; j++ {
			if sli[j] == "(hex)" || sli[j] == "(bin)" || sli[j] == "(low)" || sli[j] == "(up)" || sli[j] == "(cap)"{
				sli[j] = sli[j+1]
			}else if sli[j] == "(low, " || sli[j] == "(up, " || sli[j] == "(cap, " {
				sli[j] = sli[j+1]
			} else if sli[j] == sli[j-1] {
				sli[j] = sli[j+1]
			}
		}
	}

	for i:=0; i < len(sli)-count; i++ {
		if i < len(sli) { 
		fmt.Print(sli[i], " ")
		} else {
			fmt.Print(sli[i])
		}
	}
	fmt.Println()
	for i:=0; i < len(cc)-1 ; i++ {
	fmt.Println(cc[i]-'0')
	}
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
	for i := len(s)-1; i >= 0; i-- {
		result =result+ strings.IndexByte(base,s[i]) * power(l,j)
		j++
	}
	new := strconv.Itoa(result)
	return new
}

func power(x,y int) int {
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
		}else {
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
		}else {
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
		}else{
			word = word + string(char)
		}
	}
	return word
}
func ToLowerNum(s string) int {
	for _, char := range s {
		if char 
	}
}