# [Cryptopals](https://cryptopals.com/) [![Report Card](https://goreportcard.com/badge/github.com/mfridman/cryptopals)](https://goreportcard.com/report/github.com/mfridman/cryptopals)

> crypto challenges made fun

Chipping away at em', work in progress...

Indeed there is overlap between the challenges and one could build a standard lib. But, instead, I wanted each challenge to be its own contained unit.

## Usage

`go get -u github.com/mfridman/cryptopals/...`
```bash
# assuming you are in your GOPATH
cd src/github.com/mfridman/cryptopals/set-1/chal-6
go run main.go
```

## Set 1: Basics

- [x] **[solution](set-1/chal-1)** ........ [Convert hex to base64](https://cryptopals.com/sets/1/challenges/1)
- [x] **[solution](set-1/chal-2)** ........ [Fixed XOR](https://cryptopals.com/sets/1/challenges/2)
- [x] **[solution](set-1/chal-3)** ........ [Single-byte XOR cipher](https://cryptopals.com/sets/1/challenges/3)
- [x] **[solution](set-1/chal-4)** ........ [Detect single-character XOR](https://cryptopals.com/sets/1/challenges/4)
- [x] **[solution](set-1/chal-5)** ........ [Implement repeating-key XOR](https://cryptopals.com/sets/1/challenges/5)
- [x] **[solution](set-1/chal-6)** ........ [Break repeating-key XOR](https://cryptopals.com/sets/1/challenges/6)
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