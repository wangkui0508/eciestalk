package main

import (
	"encoding/hex"
	"fmt"

	goecies "github.com/ecies/go"
)

func main() {
	var pubkeyHex string
	fmt.Print("Enter the ecies pubkey: ")
	fmt.Scanf("%s", &pubkeyHex)
	pubkey, err := goecies.NewPublicKeyFromHex(pubkeyHex)
	if err != nil {
		fmt.Println("Cannot decode ecies pubkey")
		panic(err)
	}

	var secret string
	fmt.Print("Enter your secret: ")
	fmt.Scanf("%s", &secret)

	encrypted, err := goecies.Encrypt(pubkey, []byte(secret))
	if err != nil {
		fmt.Println("Cannot encrypt sBCH key")
		panic(err)
	}

	fmt.Printf("The encrypted Secret: %s\n", hex.EncodeToString(encrypted))
}
