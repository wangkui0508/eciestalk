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
	fmt.Println("The ecies private key:",
		hex.EncodeToString(eciesPrivKey.Bytes()))
}
