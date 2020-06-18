package main

import "github.com/mastercard/oauth1-signer-go/utils"

func main() {
	createHeader()
}




func createHeader() string {
	signingKey, err := utils.LoadSigningKey(
		"C:\Users\wildc\Documents\mastercardProject\MCD_Sandbox_mastercard-tm-test_API_Keys",
		"keystorepassword" //will need to import these for prod
		if err != {
			log.Fatalln(err)
		}
		return signingKey
	)
}

