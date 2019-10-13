package main

import "fmt"

/*
01. 「パタトクカシーー」
「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．
*/

func main() {
	word := "パタトクカシーー"
	runWord := []rune(word)

	for i := 0; i < len(runWord); i++ {
		if i%2 == 0 {
			fmt.Println(string(runWord[i]))
		}
	}
}
