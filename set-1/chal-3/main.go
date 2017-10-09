// Single-byte XOR cipher

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"
	"sync"
)

type message struct {
	key   byte
	words []byte
	score float64
}

type messages []message

func (m messages) Len() int           { return len(m) }
func (m messages) Less(i, j int) bool { return m[i].score < m[j].score }
func (m messages) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

var wg sync.WaitGroup

func main() {

	// lowerA := []byte("abcdefghijklmnopqrstuvwxyz")
	upperA := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// numbers := []byte("0123456789")
	ASCII := bytes.Join([][]byte{upperA}, []byte(""))

	s1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736" // hex string
	encoded := hexStringToBytes(s1)

	wg.Add(len(ASCII))

	c := make(chan message)

	for _, j := range ASCII {
		go func(b byte) {
			msg := decode(b, encoded)
			c <- msg
			wg.Done()
		}(j)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var scores messages

	for n := range c {
		scores = append(scores, n)
	}

	sort.Sort(scores)

	for _, j := range scores[len(scores)-5:] {
		fmt.Printf("[%c] %.2f: %s\n", j.key, j.score, j.words)
	}
	// displays top 5 scores
	// [V]: 114.09: Maaeg`i.CM)}.bgek.o.~a{`j.ah.loma`
	// [Z]: 114.29: Ammikle"OA%q"nkig"c"rmwlf"md"`caml
	// [Y]: 119.50: Bnnjhof!LB&r!mhjd!`!qntoe!ng!c`bno
	// [X]: 139.35: Cooking MC's like a pound of bacon
	// [R]: 151.68: Ieeacdm*GI-y*fcao*k*zedn*el*hkied
}

func hexStringToBytes(s string) []byte {
	// decode hex string to bytes
	str := []byte(s)
	msg := make([]byte, hex.DecodedLen(len(str)))
	hex.Decode(msg, str)
	return msg
}

func decode(k byte, enc []byte) message {
	var dec []byte

	for _, j := range enc {
		dec = append(dec, j^k)
	}

	out := message{
		words: dec,
		score: charFreq(dec),
		key:   k,
	}
	return out
}

func charFreq(b []byte) float64 {
	freq := map[byte]float64{'E': 12.02, 'T': 9.10, 'A': 8.12, 'O': 7.68, 'I': 7.31, 'N': 6.95, 'S': 6.28, 'R': 6.02, 'H': 5.92, 'D': 4.32, 'L': 3.98, 'U': 2.88, 'C': 2.71, 'M': 2.61, 'F': 2.30, 'Y': 2.11, 'W': 2.09, 'G': 2.03, 'P': 1.82, 'B': 1.49, 'V': 1.11, 'K': 0.69, 'X': 0.17, 'Q': 0.11, 'J': 0.10, 'Z': 0.07}

	var t float64

	for _, j := range bytes.ToUpper(b) {
		if v, ok := freq[j]; ok {
			t += v
		}
	}
	return t
}
