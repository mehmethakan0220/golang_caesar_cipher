package main

var template = `

usage examples

encoding 
	./sezar -o encode -a EN/en/TR/tr  -k 5   -t plainText

decoding
	./sezar -o decode -a EN/en/TR/tr  -k 5	 -t cipherText

bruteforce
	./sezar -o brute  -t cipherText


tip:
	./sezar -o encode -a en -k 5	-t lowecasemessage
	./sezar -o encode -a EN -k 5	-t UPPERCASEMESSAGE

`
