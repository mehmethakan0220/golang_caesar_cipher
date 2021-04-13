package main

func AlphabetEncoder(alphabet string, key int) string {

	pre := alphabet[0:key]
	post := alphabet[key:]
	result := post + pre
	return result
}
