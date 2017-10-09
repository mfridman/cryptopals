// Detect single-character XOR

package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

type message struct {
	encrypted []byte
	decrypted []byte
	key       byte
	score     float64
}

type messages []message

func (m messages) Len() int           { return len(m) }
func (m messages) Less(i, j int) bool { return m[i].score < m[j].score }
func (m messages) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {
	f, err := os.Open("chal4.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var items []string
	for sc.Scan() {
		items = append(items, sc.Text())
	}

	var wg sync.WaitGroup
	wg.Add(len(items))

	c := make(chan messages)

	for _, j := range items {
		go func(s string) {
			raw := hexStringToBytes(s)
			msg := decrypt(raw)
			c <- msg
			wg.Done()
		}(j)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var aggr messages

	for n := range c {
		aggr = append(aggr, n...)
	}

	sort.Sort(aggr)

	for _, j := range aggr[len(aggr)-5:] {
		fmt.Printf("score: %.2f\t[%d=%q]\t%s\n",
			j.score,
			j.key,
			j.key,
			j.decrypted,
		)
	}

}

func hexStringToBytes(s string) []byte {
	// decode hex string to bytes
	str := []byte(s)
	msg := make([]byte, hex.DecodedLen(len(str)))
	hex.Decode(msg, str)

	return msg
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

func decrypt(b []byte) messages {

	var msgs messages

	for i := 0; i <= 127; i++ {
		var dec []byte
		for _, k := range b {
			dec = append(dec, byte(i)^k)
		}
		msgs = append(msgs, message{
			key:       byte(i),
			encrypted: b,
			decrypted: dec,
			score:     charFreq(dec)})
	}

	return msgs
}
