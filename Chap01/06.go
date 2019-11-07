package main

import "fmt"

/*
06. 集合
"paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．
さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．
*/
func main() {
	wordX := "paraparaparadise"
	wordY := "paragraph"
	bigramX := bigram(wordX)
	bigramY := bigram(wordY)
	fmt.Printf("bigramX=%v\n", bigramX)
	fmt.Printf("bigramY=%v\n", bigramY)
	// 和集合の計算
	m := shugo(bigramX, bigramY)
	for k, v := range m {
		fmt.Printf("和集合：%s,%d\n", k, v)
	}
	for k, v := range m {
		if v > 1 {
			fmt.Printf("積集合：%s,%d\n", k, v)
		}
	}

}

func bigram(s string) []string {
	bigram := []string{}
	for i := 0; i < len(s)-1; i++ {
		bigram = append(bigram, s[i:i+2])
	}
	return bigram
}

func shugo(bigX []string, bigY []string) map[string]int {
	m := make(map[string]int)
	for _, x := range bigX {
		m[x]++
	}
	return m

}
