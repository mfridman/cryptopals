# [Cryptopals crypto challenges](https://cryptopals.com/)

Work in progress...

## set 1: Basics

- [x] [Convert hex to base64](https://cryptopals.com/sets/1/challenges/1)
- [x] [Fixed XOR](https://cryptopals.com/sets/1/challenges/2)
- [x] [Single-byte XOR cipher](https://cryptopals.com/sets/1/challenges/3)
- [x] [Detect single-character XOR](https://cryptopals.com/sets/1/challenges/4)
- [x] [Implement repeating-key XOR](https://cryptopals.com/sets/1/challenges/5)
- [x] [Break repeating-key XOR](https://cryptopals.com/sets/1/challenges/6)

Here things start to get interesting. hamming distance/edit distance:

- minimum # of operations (ins, del, sub) needed to transform one str into the other 


- [ ] [AES in ECB mode](https://cryptopals.com/sets/1/challenges/7)
- [ ] [Detect AES in ECB mode](https://cryptopals.com/sets/1/challenges/8)


# Misc

```go
func factorial(n uint64) uint64 {
    if n > 0 {
        result := n * factorial(n-1)
        return result
    }
    return 1
}
```