benchmark:
	go test -bench . -benchmem -benchtime 60s -timeout 60m > benchmark.txt