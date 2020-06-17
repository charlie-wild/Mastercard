package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func httpRead() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Add("Header", "Value")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return (string(body))

}
