package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) == 2 {
		file, err := os.Open("quest8.txt")
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			return
		}
		slice := make([]byte, 14)
		file.Read(slice)
		fmt.Println(string(slice))
	}
}
