package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

var folder = ".chaucha"
var genesis = "6e27bffd2a104bea1c870be76aab1cce13bebb0db40606773827517da9528174"

var magicBytes = []byte{170, 162, 38, 169}
var readHash = make([]byte, 32)
var read8 = make([]byte, 8)
var read4 = make([]byte, 4)
var read2 = make([]byte, 2)
var read1 = make([]byte, 1)
var scriptLength = 0
var readScript = make([]byte, 0)

func doubleSHA256(slice string) string {
	h1 := sha256.New()
	h2 := sha256.New()
	h1.Write([]byte(slice))
	h2.Write(h1.Sum(nil))
	return hex.EncodeToString(reverse(h2.Sum(nil)))

}

func byte2int(slice []byte) int {
	data := uint64(0)
	for _, b := range slice {
		data = (data << 8) | uint64(b)
	}
	return int(data)
}

func reverse(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func readVariableInt(f *os.File) int {
	f.Read(read1)
	txCount := byte2int(read1)
	switch txCount {
	case 253:
		f.Read(read2)
		txCount = byte2int(reverse(read2))
	case 254:
		f.Read(read4)
		txCount = byte2int(reverse(read4))
	case 255:
		f.Read(read8)
		txCount = byte2int(reverse(read8))
	}
	return txCount
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
