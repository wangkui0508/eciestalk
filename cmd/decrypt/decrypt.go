package main

import (
	"encoding/hex"
	"fmt"
	"log"

	goecies "github.com/ecies/go"
)

func main() {
	eciesPrivKey, err := goecies.GenerateKey()
	if err != nil {
		log.Fatal("failed to gen ecies key: ", err)
	}
	fmt.Println("The ecies pubkey:",
		hex.EncodeToString(eciesPrivKey.PublicKey.Bytes(true)))

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
