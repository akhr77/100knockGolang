package main

import (
	"fmt"
	"strings"
)

/*
03. 円周率
"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，
各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
*/

func main() {
	line := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	words := strings.Fields(line)
	list := []int{}
	for _, word := range words {
		word = strings.TrimRight(word, ",")
		word = strings.TrimRight(word, ".")
		list = append(list, len([]rune(word)))
	}
	fmt.Println(list)
}
