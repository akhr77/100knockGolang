package main

import (
	"fmt"
	"strings"
)

/*
04. 元素記号
"Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，
取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
*/

func main() {
	line := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	words := strings.Fields(line)
	num := []int{1, 5, 6, 7, 8, 9, 15, 16}
	m := make(map[string]int)
	for i, word := range words {
		word = strings.TrimRight(word, ".")
		for _, n := range num {
			if i+1 == n {
				m[string(word[0])] = i + 1
			} else {
				m[string(word[0:2])] = i + 1
			}
		}
	}
	fmt.Println(m)
}
