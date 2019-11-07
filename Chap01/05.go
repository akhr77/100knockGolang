package main

import (
	"fmt"
	"strings"
)

/*
05. n-gram
与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．
*/

func main() {
	line := "I am an NLPer"
	words := strings.Fields(line)

	for i := 0; len(words)-1 > i; i++ {
		fmt.Println(words[i], "-", words[i+1]) 
	}

}
