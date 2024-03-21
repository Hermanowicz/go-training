package main

import "fmt"

func main() {
	for n := range 100 {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
}
