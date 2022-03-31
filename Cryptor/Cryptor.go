package Cryptor

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	crand "math/rand"
	"time"
)

const capletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	ErrInvalidBlockSize = errors.New("[-] Invalid Blocksize")

	ErrInvalidPKCS7Data = errors.New("[-] Invalid PKCS7 Data (Empty or Not Padded)")

	ErrInvalidPKCS7Padding = errors.New("[-] Invalid Padding on Input")
)

func Pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func RandomBuffer(size int) []byte {
	buffer := make([]byte, size)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return buffer
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[crand.Intn(len(letters))]

	}
	return string(b)
}

// this is how the stuff is stored when called
// generate seed using the current time in UnixNano measurement
//generate randomNum using INTERVAL (max-min)
// then add minimum to it
// store the value as n
// this is where we call RandStringBytes (above)
// generate n into random string of bytes
    // within RandStringBytes
	// create a byte array for the BUFFER the size of number
		// THIS IS IMPORTANT - THIS IS HOW IT KNOWS
		// by doing this it seems like - okay, I can figure out what N should be based on the size of the randomByteString
	// now assigning each array slot to a random entry from variable 'letters', this is done by calculating the length, converting to Int, then using crand.

func VarNumberLength(min, max int) string {
	var r string
	crand.Seed(time.Now().UnixNano())
	num := crand.Intn(max-min) + min
	n := num
	r = RandStringBytes(n)
	return r
}

func printHexOutput(input ...[]byte) {
	for _, i := range input {
		fmt.Println(hex.EncodeToString(i))
	}
}

func GenerateNumer(min, max int) int {
	crand.Seed(time.Now().UnixNano())
	num := crand.Intn(max-min) + min
	n := num
	return n

}

func CapLetter() string {
	n := 1
	b := make([]byte, n)
	for i := range b {
		b[i] = capletters[crand.Intn(len(capletters))]

	}
	return string(b)
}
