// Break repeating-key XOR

package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"sort"
	"strings"
	"unicode/utf8"
)

// A KeyChar represents a single-byte Char and its freq score.
type KeyChar struct {
	Char  byte
	Score float64 // Frequency score
}

// KeyChars holds slice of KeyChar.
type KeyChars []KeyChar

func (k KeyChars) Len() int           { return len(k) }
func (k KeyChars) Less(i, j int) bool { return k[i].Score < k[j].Score }
func (k KeyChars) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

// HighestScore returns KeyChar, i.e., the single-byte XOR, with the highest frequency score.
func (k KeyChars) HighestScore() KeyChar {
	sort.Sort(k)
	return k[len(k)-1]
}

// A KeySize represents the length and normalized, averaged, Hamming distance of a key.
type KeySize struct {
	Length int
	Score  float64 // Hamming distance
}

// Keys holds a slice of KeySize.
type Keys []KeySize

func (k Keys) Len() int { return len(k) }
func (k Keys) Less(i, j int) bool {
	return k[i].Score < k[j].Score || isNaN(k[i].Score) && !isNaN(k[j].Score)
}
func isNaN(f float64) bool   { return f != f }
func (k Keys) Swap(i, j int) { k[i], k[j] = k[j], k[i] }

// LowestScore returns KeySize with the lowest normalized, averaged, Hamming distance score
func (k Keys) LowestScore() KeySize {
	sort.Sort(k)
	return k[0]
}

func main() {

	b, err := ioutil.ReadFile("chal6.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// decode base64 encoded string to get ciphertex as []byte
	ciphertext, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		log.Fatalln(err)
	}

	var ks Keys
	// test key lengths between 2 and 50
	for len := 2; len <= 50; len++ {
		// TODO think about adding concurrency here
		ks = append(ks,
			KeySize{
				Score:  KeyScore(ciphertext, len),
				Length: len,
			})
	}

	// display key with the loweset averaged normalized edit distance/Hamming distance
	fmt.Printf("Likely key size is [%d] with normalized Hamming distance score of [%.4f]\n", ks.LowestScore().Length, ks.LowestScore().Score)

	tb := TransposeBlocks(ciphertext, ks.LowestScore().Length)

	var key []byte
	for _, j := range tb {
		// solve each block as if it was single-character XOR and choose the one with the highest score,
		// usually the single-byte XOR that yields the highest score is the repeating-key XOR key byte for that block
		kc := ScoreBlock(j)
		key = append(key, kc.HighestScore().Char)
	}

	// display key
	fmt.Println(strings.Repeat("-", utf8.RuneCount(key)))
	fmt.Printf("%s\n", key)
	fmt.Println(strings.Repeat("-", utf8.RuneCount(key)))

	// finally, decrypt original ciphertext with the key
	final := EncryptDecryptXOR(ciphertext, key)
	fmt.Printf("%s\n", final)

	// that was fun!
}

// ScoreBlock takes block of bytes and scores each block against a single-character XOR.
// Uses char frequency of first 127 standard chars. Returns slice of KeyChar.
func ScoreBlock(b []byte) KeyChars {
	var kc KeyChars

	for i := 0; i <= 127; i++ {
		var xor []byte
		for _, k := range b {
			xor = append(xor, byte(i)^k)
		}
		kc = append(kc, KeyChar{
			Char:  byte(i),
			Score: CharFreqScore(xor)})
	}

	return kc
}

// CharFreqScore scores bytes of English text and returns sum of each character's frequency score.
func CharFreqScore(b []byte) float64 {
	// freq := map[byte]float64{'E': 12.02, 'T': 9.10, 'A': 8.12, 'O': 7.68, 'I': 7.31, 'N': 6.95, 'S': 6.28, 'R': 6.02, 'H': 5.92, 'D': 4.32, 'L': 3.98, 'U': 2.88, 'C': 2.71, 'M': 2.61, 'F': 2.30, 'Y': 2.11, 'W': 2.09, 'G': 2.03, 'P': 1.82, 'B': 1.49, 'V': 1.11, 'K': 0.69, 'X': 0.17, 'Q': 0.11, 'J': 0.10, 'Z': 0.07}
	/*
		switched to this frequency distro
		http://www.fitaly.com/board/domper3/posts/136.html
	*/

	freqMap := map[byte]float64{9: 0.0057, 23: 0.0000, 32: 17.1662, 33: 0.0072, 34: 0.2442, 35: 0.0179, 36: 0.0561, 37: 0.0160, 38: 0.0226, 39: 0.2447, 40: 0.2178, 41: 0.2233, 42: 0.0628, 43: 0.0215, 44: 0.7384, 45: 1.3734, 46: 1.5124, 47: 0.1549, 48: 0.5516, 49: 0.4594, 50: 0.3322, 51: 0.1847, 52: 0.1348, 53: 0.1663, 54: 0.1153, 55: 0.1030, 56: 0.1054, 57: 0.1024, 58: 0.4354, 59: 0.1214, 60: 0.1225, 61: 0.0227, 62: 0.1242, 63: 0.1474, 64: 0.0073, 65: 0.3132, 66: 0.2163, 67: 0.3906, 68: 0.3151, 69: 0.2673, 70: 0.1416, 71: 0.1876, 72: 0.2321, 73: 0.3211, 74: 0.1726, 75: 0.0687, 76: 0.1884, 77: 0.3529, 78: 0.2085, 79: 0.1842, 80: 0.2614, 81: 0.0316, 82: 0.2519, 83: 0.4003, 84: 0.3322, 85: 0.0814, 86: 0.0892, 87: 0.2527, 88: 0.0343, 89: 0.0304, 90: 0.0076, 91: 0.0086, 92: 0.0016, 93: 0.0088, 94: 0.0003, 95: 0.1159, 96: 0.0009, 97: 5.1880, 98: 1.0195, 99: 2.1129, 100: 2.5071, 101: 8.5771, 102: 1.3725, 103: 1.5597, 104: 2.7444, 105: 4.9019, 106: 0.0867, 107: 0.6753, 108: 3.1750, 109: 1.6437, 110: 4.9701, 111: 5.7701, 112: 1.5482, 113: 0.0747, 114: 4.2586, 115: 4.3686, 116: 6.3700, 117: 2.0999, 118: 0.8462, 119: 1.3034, 120: 0.1950, 121: 1.1330, 122: 0.0596, 123: 0.0026, 124: 0.0007, 125: 0.0026, 126: 0.0003, 131: 0.0000, 149: 0.6410, 183: 0.0010, 223: 0.0000, 226: 0.0000, 229: 0.0000, 230: 0.0000, 237: 0.0000}

	var sum float64
	for _, j := range b {
		if f, ok := freqMap[j]; ok {
			sum += f
		}
	}

	return sum
}

// TransposeBlocks breaks ciphertext into blocks of key length
// and transposes the blocks, returning a 2D slice of bytes.
// e.g., block that is the first byte of every block, block that is the second byte of every block, etc.
func TransposeBlocks(b []byte, l int) [][]byte {

	p := len(b) / l // INFO this will not include the remainder of len(b) mod (2*l)

	var c int
	blocks := make([][]byte, l)
	for i := 0; i < p; i++ {
		for j := range b[c : c+l] {
			blocks[j] = append(blocks[j], b[c : c+l][j])
		}
		c += l
	}

	return blocks
}

// KeyScore takes ciphertext and key length
// returns the normalized, averaged, edit distance, i.e., Hamming distance
// lower numbers have higher likelihood of being the correct key length
func KeyScore(b []byte, l int) float64 {
	/*
		tried using k-(n-1) where k is total and n is key size
		based on https://math.stackexchange.com/q/1611927, but this obviously didn't work

		may be worth truingtrying all combinations, e.g., !n / r!(n-r)!
		but this may be computationally expensive
	*/

	p := len(b) / (2 * l) // INFO this will not include the remainder; len(b) mod (2*l)

	var sum float64
	var c int
	for i := 0; i < p; i++ {
		dist := HammingDistance(b[c:c+l], b[c+l:c+(l*2)])
		c = c + (l * 2)
		sum += (float64(dist) / float64(l))
	}

	return float64(sum) / float64(p)
}

// HammingDistance takes a pair of equal-length bytes and returns the edit distance, i.e., Hamming distance
// which is just the number of differing bits
// NOTE for binary strings a and b, the Hamming distance is equal to the number of ones (population count) in a XOR b
func HammingDistance(n, m []byte) int {
	if len(n) != len(m) {
		log.Fatalln(errors.New("cannot compute Hamming distance, len(n) != len(m)"))
	}

	var dist int
	for i := range n {
		if n[i] != m[i] {
			dist += bits.OnesCount64(uint64(n[i] ^ m[i]))
		}
	}

	return dist
}

// EncryptDecryptXOR takes ciphertext and performs XOR encryption or decryption
// based on provided key. Returns bytes and caller determines output format
func EncryptDecryptXOR(b, k []byte) []byte {
	var out []byte
	for i := range b {
		out = append(out, b[i]^k[i%len(k)])
	}
	// to return a hex string use hex.EncodeToString(out)
	return out
}
