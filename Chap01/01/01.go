package main

import "fmt"

func main() {
	word := "パタトクカシーー"
	runWord := []rune(word)

	for i := 0; i < len(runWord); i++ {
		if i%2 == 0 {
			fmt.Println(string(runWord[i]))
		}
	}
}
