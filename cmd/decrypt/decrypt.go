package main

import (
	"encoding/hex"
	"fmt"
	"log"

	goecies "github.com/ecies/go"
)

func main() {
	var privateKeyHex string
	fmt.Print("Enter the ecies private key: ")
	fmt.Scanf("%s", &privateKeyHex)
	eciesPrivKey, err := goecies.NewPrivateKeyFromHex(privateKeyHex)
	if err != nil {
		panic(err)
	}
	var inputHex string
	fmt.Printf("Enter the encrypted secret: ")
	_, _ = fmt.Scanf("%s", &inputHex)
	bz, err := hex.DecodeString(inputHex)
	if err != nil {
		log.Fatal("cannot decode hex string: ", err)
	}
	bz, err = goecies.Decrypt(eciesPrivKey, bz)
	if err != nil {
		log.Fatal("cannot decrypt: ", err)
	}
	println(string(bz))
}
