package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"

	res := fixedXOR(
		hexStringToBytes(s1),
		hexStringToBytes(s2),
	)

	// https://golang.org/pkg/fmt/#pkg-overview
	fmt.Printf("%x\n", res) // 746865206b696420646f6e277420706c6179
}

func hexStringToBytes(s string) []byte {
	// decode hex string to bytes
	str := []byte(s)
	msg := make([]byte, hex.DecodedLen(len(str)))
	hex.Decode(msg, str)
	return msg
}

func fixedXOR(a, b []byte) []byte {
	if len(a) != len(b) {
		return nil
	}

	var out []byte
	for i := range a {
		out = append(out, a[i]^b[i])
	}
	return out
}
