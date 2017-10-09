// Implement repeating-key XOR

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const key = "ICE"

func main() {

	b, err := ioutil.ReadFile("chal5.txt")
	if err != nil {
		log.Fatalln(err)
	}

	msg := encryptDecrypt(b, key)

	// hex encoded string
	fmt.Printf("%x\n", msg)
	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
	// a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

	// decrypt and get back original msg
	fmt.Printf("%s\n", encryptDecrypt(msg, key))
	/*
		Burning 'em, if you ain't quick and nimble
		I go crazy when I hear a cymbal
	*/

}

func encryptDecrypt(b []byte, k string) []byte {

	var out []byte
	for i := range b {
		out = append(out, b[i]^k[i%len(k)])
	}
	// to return a hex string use hex.EncodeToString(b),
	// otherwise return byets and let caller determine output format
	return out
}
