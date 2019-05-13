package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	for _, v := range a {
		switch v {
		case 'a', 'i', 'u', 'e', 'o':
		default:
			fmt.Print(string(v))
		}
	}
	fmt.Println()
}
