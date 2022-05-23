package main

import (
	"flag"
	"fmt"
)

var (
	txt = flag.String("txt", "jeketek itu hebat", "-txt=text_yg_akan_dihitung_brp_charnya")
)

func main() {
	flag.Parse()

	fmt.Println("Kata yang dimasukkan adalah:", *txt)

	// var word = "zanur prihatna"
	// for _, w := range word {
	// 	fmt.Printf("%q %t\n", w, isVocal(w))
	// }

	// fmt.Println(getVocalInWord(word))
	// fmt.Println(getConsonantInWord(word))

	// v, k := GetCharCount(word)
	// fmt.Println("vocals:")
	// for k, d := range v {
	// 	fmt.Println(k, d)
	// }

	// fmt.Println("consonant:")
	// for k, d := range k {
	// 	fmt.Println(k, d)
	// }

	vocals, consonans := GetCharCount(*txt)

	fmt.Println("vocals:", vocals)
	fmt.Println("consonans:", consonans)
}

func isVocal(c rune) bool {
	switch c {
	case 'A', 'I', 'U', 'E', 'O', 'a', 'i', 'u', 'e', 'o':
		return true
	default:
		return false
	}
}

func getVocalInWord(txt string) int {
	n := 0
	for _, v := range txt {
		if isVocal(v) {
			n++
		}
	}
	return n
}

func getConsonantInWord(txt string) int {
	n := 0
	for _, v := range txt {
		if !isVocal(v) && v != ' ' {
			n++
		}
	}
	return n
}

type CharCounts map[string]int

//-- exercise1: implementasikan GetCharCount, buat test scriptnya dan buat console-app
//-- exercise2: palindrome
func GetCharCount(txt string) (vocals CharCounts, consonans CharCounts) {
	//-- implementations
	vocals = make(CharCounts)
	// consonans = make(CharCounts)
	consonans = map[string]int{}
	// consonans = *new(map[string]int)
	for _, v := range txt {
		if isVocal(v) {
			vocals[string(v)] = vocals[string(v)] + 1
		} else if !isVocal(v) && v != ' ' {
			consonans[string(v)] = consonans[string(v)] + 1
		}
	}

	return vocals, consonans
}
