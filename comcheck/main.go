package main

import (
	"fmt"
	"os"
)

func main() {
	mots := []string{"01", "galaxy", "galaxy 01"}
	cmpt := 0
	samba := os.Args[1:]

	for i := range samba {
		for _, s := range mots {
			if samba[i] == s {
				cmpt++
			}
		}
	}
	if cmpt >= 1 {
		fmt.Println("Alert!!!")
	}
}
