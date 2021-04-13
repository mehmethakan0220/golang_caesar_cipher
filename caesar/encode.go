package main

import (
	"fmt"
	"strings"
)

//Encode ...
func Encode(plainText *string, key int, alphabet string) {

	calphabet := AlphabetEncoder(alphabet, key)

	var bytearr []int
	for _, v := range *plainText {
		result := strings.Index(alphabet, string(v))

		bytearr = append(bytearr, result)
	}

	for _, v := range bytearr {
		fmt.Printf("%v", string(calphabet[v]))
	}
	fmt.Println()

}
