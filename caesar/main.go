package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	operation := flag.String("o", "encode", "./sezar -o(encode/decode/brute)")
	alphabet := flag.String("a", "en", "en/tr")
	text := flag.String("t", "", "plain or cipher text")
	key := flag.Int("k", 0, "key")
	flag.Parse()

	fmt.Printf("%s,%s,%s,%d \n", *operation, *alphabet, *text, *key)

	context, err := ioutil.ReadFile("./alphabets.json")
	if err != nil {
		log.Fatal("read file error", err)
	}

	alf := make(map[string]string)
	err = json.Unmarshal(context, &alf)
	if err != nil {
		log.Fatal("json convert error", err)
	}

	if *operation == "encode" {
		Encode(text, *key, alf[*alphabet])
	} else if *operation == "decode" {
		Decode(*text, *key, alf[*alphabet])
	} else if *operation == "brute" {
		BruteForce(*text)
	} else {
		fmt.Println(template)
	}

}
