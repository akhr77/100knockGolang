package main

import "fmt"

/*
00. 文字列の逆順
文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
*/

func main() {
	word := "stressed"

	for i := len(word) - 1; i >= 0; i-- {
		fmt.Print(string(word[i]))
	}
	fmt.Print("\n")
}
