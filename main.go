package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	number int
}

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
}

type Body struct {
	Data Anime `json:"data"`
}

var jobs = make(chan Job, 5)
var results = make(chan Anime, 5)
var URL = "/api/edge/anime/"
var resultsCollection []Anime

func main() {
	const numOfJobs = 100
	go allocateJobs(numOfJobs)

	resultsReceived := make(chan bool)
	go getResults(resultsReceived)

	const numOfWorkers = 5
	createWorkerPool(numOfWorkers)

	<-resultsReceived

	resultsJSON, err := json.MarshalIndent(resultsCollection, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	writeToFile(resultsJSON)
}

func writeToFile(json []byte) {
	err := ioutil.WriteFile("results.json", json, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func allocateJobs(numOfJobs int) {
	for i := 1; i <= numOfJobs; i++ {
		jobs <- Job{i}
	}

	close(jobs)
}

func createWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		go createWorker(&wg)
		wg.Add(1)
	}
	wg.Wait()
	close(results)
}

func createWorker(wg *sync.WaitGroup) {
	for job := range jobs {
		anime := animeFetcher(URL, job.number)

		results <- anime
	}

	wg.Done()
}

func getResults(resultsReceived chan bool) {
	for result := range results {
		resultsCollection = append(resultsCollection, result)
	}

	resultsReceived <- true
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

	response, err := client.Do(&request)
	if err != nil {
		fmt.Println(err)
	}

	var body Body

	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}

	return body.Data
}
