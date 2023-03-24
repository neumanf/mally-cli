package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const BaseUrl string = "https://api.mally.neumanf.com/api"

func PostRequest[Req any, Res any](path string, data Req, token *string) Res {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", BaseUrl+path, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")
	if token != nil {
		req.Header.Set("Cookie", "accessToken="+*token)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var res Res

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		log.Fatal("Could not decode API response", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode == 401 {
			log.Fatal("You are not authorized. If you are trying to login, check your credentials. Otherwise, please use 'mally-cli login' first to login and then try again.")
		}

		log.Fatal("Status code: ", resp.StatusCode, ", response: ", res)
	}

	return res
}
