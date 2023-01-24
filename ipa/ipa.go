package ipa

import (
	"fmt"
	"mechobiral/bengali"
)

func GetUnicodeName(c rune) string{
	x := bengali.GetMergedMap(true , true)
	for t , v := range x{
		if v == c{
			return t
		}
	}
	
	return string(c)
}

func JustPrint(word string){
	wordr := []rune(word)
	fmt.Println(word)
	for _ , item := range wordr{
		
		fmt.Printf("%U -> %s\n" , item , GetUnicodeName(item))
	}
}
