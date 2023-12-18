package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AnimeAttribute struct {
	Synopsis       string `json:"synopsis"`
	Description    string `json:"description"`
	Rating         string `json:"averageRating"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	EpisodeCount   int32  `json:"episodeCount"`
	EpisodeLength  int32  `json:"episodeLength"`
	ShowType       string `json:"showType"`
	RatingRank     int32  `json:"ratingRank"`
	PopularityRank int32  `json:"popularityRank"`
}

type Anime struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes AnimeAttribute `json:"attributes"`
	// Relationships map[string]interface{} `json:"-"`
}

type Body struct {
	Data Anime `json:"data"`
}

func main() {
	URL := "/api/edge/anime/"
	ID := 45857

	anime := animeFetcher(URL, ID)

	fmt.Printf("%+v\n", anime)
}

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
		Timeout: 5 * time.Minute,
	}

	response, error := client.Do(&request)
	if error != nil {
		fmt.Println(error)
	}

	var body Body

	err := json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}

	return body.Data
}
