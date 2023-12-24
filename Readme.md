# Basic Concurrency Project: API Fetcher
This project is built for the purpose of understanding The Go Programming Language main sellpoints: Concurrency with Go Routines.

This is intended to acommodate my learning of Concurrency in Go Programming Language.

## Objectives
This personal project was built with these questions to be answered.

- Are _concurrent_ requests actually faster than _sequential_ requests?
- What are someother usecases for using Go Routines?
- Are _concurrent_ writing to a file faster than doing it sequentially/manually?
- How many number of workers in a Worker Pool are the most efficient?

## How it goes
The following are description of main files used during my learning proccess.
- `http-fetcher.go` : stores a basic function for sending `GET` request to [Kitsu API](https://kitsu.docs.apiary.io/).
- `main.go` : stores 3 different function, later benchmarked for comparison. Each of these functions sends multiple API requests, then write the response to a file. The difference are how these concluded:
    
    - Concurrent requests
    - Concurrent requests and write to file
    - Sequential requests

- `main_test.go` : contains the test functions used for benchmarking
- `types.go` : defines types used 

## Results
> Are _concurrent_ requests actually faster than _sequential_ requests?
- Concurrent _requests_ are, definitely, faster and less resource intensive than Sequential one. This is proven by the benchmark result stored in `benchmark.txt`
> What are someother usecases for using Go Routines?
- Go routines use cases, then, includes something that can be blocking such as API calls and Database calls.
> Are _concurrent_ writing to a file faster than doing it sequentially/manually?
- In this example, concurrently writing to a file doesn't really show that much of advantage. This is probably because we still need mutex to make sure only one go routine is accessing the file at a time as seen on this [blog post](https://www.markhneedham.com/blog/2017/01/31/go-multi-threaded-writing-csv-file/).
> How many number of workers in a Worker Pool are the most efficient?
- According to [this Stackoverflow answer](https://stackoverflow.com/questions/47958826/optimal-size-of-worker-pool), there is no objective number to how many number of workers are the most efficient for a Worker Pool. The only way to find one is to benchmark every configuration.