package main

import "fmt"

/*
02. 「パトカー」＋「タクシー」＝「パタトクカシーー」
「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．
*/
func main() {
	wordA := "パトカー"
	wordB := "タクシー"
	runeWordA := []rune(wordA)
	runeWordB := []rune(wordB)
	var lenWord int
	if len(runeWordA) >= len(runeWordB) {
		lenWord = len(runeWordA)
	} else {
		lenWord = len(runeWordB)
	}
	for i := 0; i < lenWord; i++ {
		if string(runeWordA[i]) != " " {
			fmt.Println(string(runeWordA[i]))
		}
		if string(runeWordB[i]) != " " {
			fmt.Println(string(runeWordB[i]))
		}
	}

}
