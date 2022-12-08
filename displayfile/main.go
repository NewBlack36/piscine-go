package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	samba := os.Args[1:]

	length := 0
	for i := range samba {
		length = i + 1
	}

	if length > 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if samba[0] == "quest8.txt" {

		content, err := ioutil.ReadFile(samba[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf(string(content))
	}
}
