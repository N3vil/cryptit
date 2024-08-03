package main

import (
	"fmt"
	"github.com/N3vil/cryptit/encrypt"
	"github.com/N3vil/cryptit/decrypt"
)

func main() {

	text := "Hello World"
	secret := encrypt.Nimbus(text)
	fmt.Println(secret)
	fmt.Println(decrypt.Nimbus(secret))
}