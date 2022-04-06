package main

import (
	"cryptox/app/clients"
	"flag"
	"fmt"
	"log"
)

func main() {
	fiatCurency := flag.String(
		"fiat", "USD", "NG",
	)

	nameOfCrypto := flag.String(
		"crypto", "BTC", "ETH",
	)

	flag.Parse()

	crypto, err := client.FetchCryto(*fiatCurency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
