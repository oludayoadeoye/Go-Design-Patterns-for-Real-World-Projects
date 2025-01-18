package main

import "fmt"

func main() {
	ctr := 5
	for ctr > 0 {
		fmt.Println("say, thank you: ", ctr)
		ctr -= 1
	}
}
