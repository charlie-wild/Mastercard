package main

import "github.com/mastercard/oauth1-signer-go/utils"

func createHeader() string {
	signingKey, err := utils.LoadSigningKey()
	//test
}
