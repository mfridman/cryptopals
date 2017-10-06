package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	arg := os.Args

	if len(arg[1:]) == 0 {
		log.Fatalf("error: [%v] takes file as an Arg", os.Args[0])
	}

	// read string from file
	raw, err := ioutil.ReadFile(arg[1]) // 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
	if err != nil {
		log.Fatalln(err)
	}

	// decode hex to bytes
	msg := make([]byte, hex.DecodedLen(len(raw)))
	_, err = hex.Decode(msg, raw)
	if err != nil {
		log.Fatalln(err)
	}

	// encode bytes to string
	str := base64.StdEncoding.EncodeToString(msg)

	fmt.Println(str) // SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

}
