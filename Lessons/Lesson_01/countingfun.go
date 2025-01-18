package main

import "fmt"

func main() {
	ctr := 0
	for ctr < 12 {
		fmt.Println("say I love you: ", ctr)
		ctr += 1
	}
}
