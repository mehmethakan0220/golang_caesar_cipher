package main

import (
	"fmt"
	"strings"
)

func BruteForce(cipherText string) {
	for i := 0; i < 26; i++ {
		alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		pre := alphabet[0:i]
		post := alphabet[i:]

		calphabet := post + pre

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
}
