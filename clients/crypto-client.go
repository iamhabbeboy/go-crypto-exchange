package client

import (
	"cryptox/app/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func FetchCrypto(fiat string, crypto string) (string, error) {
	api := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")
	URL := api + "/v1/currencies/ticker?key=" + apiKey + "&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal("Error occured while fetch request")
	}
	defer resp.Body.Close()

	var cResp model.Cryptoresponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("Error occured while decoding data")
	}

	return cResp.TextOutput(), nil
}
