package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var URL = "/api/edge/anime/"

func ConcurrentFetch(numOfJobs int, numOfWorkers int) {
	var (
		jobs              = make(chan Job, numOfWorkers)
		results           = make(chan Anime, numOfWorkers)
		resultsCollection []Anime
	)

	var allocateJobs = func(numOfJobs int) {
		for i := 1; i <= numOfJobs; i++ {
			jobs <- Job{i}
		}

		close(jobs)
	}

	var createWorker = func(wg *sync.WaitGroup) {
		for job := range jobs {
			anime := animeFetcher(URL, job.number)

			results <- anime
		}

		wg.Done()
	}

	var createWorkerPool = func(numOfWorkers int) {
		var wg sync.WaitGroup
		for i := 0; i < numOfWorkers; i++ {
			go createWorker(&wg)
			wg.Add(1)
		}
		wg.Wait()
		close(results)
	}

	var getResults = func(resultsReceived chan bool) {
		for result := range results {
			resultsCollection = append(resultsCollection, result)
		}

		resultsReceived <- true
	}

	go allocateJobs(numOfJobs)

	resultsReceived := make(chan bool)
	go getResults(resultsReceived)

	createWorkerPool(numOfWorkers)

	<-resultsReceived

	resultsJSON, err := json.MarshalIndent(resultsCollection, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	writeToFile(resultsJSON)
}

func SequentialFetch(numOfAnimes int) {
	var animeCollection []Anime
	for i := 1; i <= numOfAnimes; i++ {
		animeCollection = append(animeCollection, animeFetcher(URL, i))
	}

	resultsJSON, err := json.MarshalIndent(animeCollection, "", " ")
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
