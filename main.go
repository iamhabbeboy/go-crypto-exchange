package main

import (
	"cryptox/app/clients"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load Env")
	}

	fiatCurrency := flag.String(
		"fiat", "USD", "NG",
	)

	nameOfCrypto := flag.String(
		"crypto", "BTC", "ETH",
	)

	flag.Parse()

	fmt.Println("Processing request.... 🧐")

	crypto, err := client.FetchCrypto(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
