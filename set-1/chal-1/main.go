// Convert hex to base64

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {

	s1 := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	// decode hex to bytes
	msg := make([]byte, hex.DecodedLen(len(s1)))
	hex.Decode(msg, s1)

	// encode bytes to string
	str := base64.StdEncoding.EncodeToString(msg)

	fmt.Println(str) // SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}
