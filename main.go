package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	writeToFile(resultsJSON)
}

func ConcurrentFetchWrite(numOfJobs int, numOfWorkers int) {
	var (
		jobs    = make(chan Job, numOfWorkers)
		results = make(chan Anime, numOfWorkers)
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
		CSVFile, err := os.Create("results.csv")
		if err != nil {
			log.Fatalf("Failed creating file %s", err)
		}

		CSVWriter := csv.NewWriter(CSVFile)
		var mu sync.Mutex

		for result := range results {
			go func(CSVWriter *csv.Writer, result Anime, mu *sync.Mutex) {
				mu.Lock()
				defer mu.Unlock()
				CSVWriter.Write([]string{
					result.ID,
					result.Type,
					result.Attributes.Rating,
					strconv.Itoa(result.Attributes.RatingRank),
					strconv.Itoa(result.Attributes.PopularityRank),
					result.Attributes.Description,
					result.Attributes.Synopsis,
					result.Attributes.StartDate,
					result.Attributes.EndDate,
					result.Attributes.ShowType,
					strconv.Itoa(result.Attributes.EpisodeCount),
					strconv.Itoa(result.Attributes.EpisodeLength),
				})
			}(CSVWriter, result, &mu)
		}

		resultsReceived <- true
	}

	go allocateJobs(numOfJobs)

	resultsReceived := make(chan bool)
	go getResults(resultsReceived)

	createWorkerPool(numOfWorkers)

	<-resultsReceived
}

func writeToFile(json []byte) {
	err := ioutil.WriteFile("results.json", json, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
