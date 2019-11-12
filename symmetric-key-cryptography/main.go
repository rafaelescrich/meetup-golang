package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// AES-GCM should be used because the operation is an authenticated encryption
// algorithm designed to provide both data authenticity (integrity) as well as
// confidentiality.

// Merged into Golang in https://go-review.googlesource.com/#/c/18803/

func encrypt() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("meetup-golang123")
	plaintext := []byte("Este é o texto plano a ser cifrado")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := []byte("meetupgolang")

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Texto cifrado: %x\n", ciphertext)
}

func decrypt() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("meetup-golang123")
	ciphertext, _ := hex.DecodeString("8749ecbe2067faa0dff0456cd7edb147b0121ca151ef9bd44461a64fad684564bc0dd4899b65121380c851bc4209a7631f1ee9")

	nonce := []byte("meetupgolang")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Texto plano original: %s\n", string(plaintext))
}

func main() {

	// cifrar texto

	encrypt()

	// decifrar texto

	decrypt()
}
