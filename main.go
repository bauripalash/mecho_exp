package main

import (
	"fmt"
	"mechobiral/ipa"
	"os"
	"strings"
)

func ipaTestWords(){
	f , _ := os.ReadFile("testwords/sample.txt")
	data := strings.Split(string(f), "\n")

	for _, item := range data{
		ipa.JustPrint(item)
	}

}

func main() {
//	fmt.Println(bengali.BN_VOWELS_MAP["aa"] == rune('আ'))
	ipaTestWords()
//	fmt.Println(bengali.GetMergedMap(true , true))
	fmt.Println(1)
}
