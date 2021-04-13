package main

import (
	"fmt"
	"strings"
)

//Decode ...
func Decode(cipherText string, key int, alphabet string) {

	calphabet := AlphabetDecoder(alphabet, key)

	var bytearr []int
	for _, v := range cipherText {
		result := strings.Index(calphabet, string(v))

		bytearr = append(bytearr, result)
	}

	for _, v := range bytearr {
		fmt.Printf("%v", string(alphabet[v]))
	}
	fmt.Println()
}
