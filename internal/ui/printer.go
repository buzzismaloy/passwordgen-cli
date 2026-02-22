package ui

import "fmt"

func PrintPassword(p string) {
	fmt.Println("Generated password:")
	fmt.Println(p)
}

func PrintError(err error) {
	fmt.Println(err)
}
