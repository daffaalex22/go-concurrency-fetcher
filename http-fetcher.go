package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func animeFetcher(Url string, id int) Anime {
	request := http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   "kitsu.io",
			Path:   Url + strconv.Itoa(id),
		},
	}

	client := http.Client{
		Timeout: 10 * time.Minute,
	}

	response, err := client.Do(&request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	var body Body

	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	return body.Data
}
