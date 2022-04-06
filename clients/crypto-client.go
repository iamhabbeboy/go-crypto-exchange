package client

import (
	"cryptox/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func FetchCryto(fiat string, crypto string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat

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
