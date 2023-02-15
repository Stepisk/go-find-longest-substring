package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sort"
	"time"
)

type Data struct {
	text                 []byte
	startIndex, endIndex int
}

func main() {
	filePath := flag.String("path", "data.txt", "path to the file to be processed")
	numOfWorkers := flag.Int("num-workers", runtime.NumCPU(), "amount of workers")
	flag.Parse()

	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jobs := make(chan Data, *numOfWorkers)
	results := make(chan []string, 1)

	for i := 0; i < *numOfWorkers; i++ {
		go worker(i, jobs, results)
	}

	timeStart := time.Now()
	for i := 0; i < len(data); i++ {
		jobs <- Data{data, i, len(data)}
	}
	close(jobs)

	result := make([]string, 0)
	for i := 0; i < *numOfWorkers; i++ {
		result = append(result, <-results...)
	}
	elapsed := time.Since(timeStart)

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	fmt.Println(result[0])
	fmt.Println(result)
	fmt.Printf("calculated in %s\n", elapsed)
}

func worker(id int, jobs <-chan Data, results chan<- []string) {
	maxSubstringLen := 0
	substringList := make([]string, 0)

	for job := range jobs {
		i := job.startIndex
		for j := i + 2; j < job.endIndex; j++ {
			substring := job.text[i:j]
			if bytes.Count(job.text, substring) > 1 {
				if len(substring) > maxSubstringLen {
					maxSubstringLen = len(substring)
					substringList = append(substringList, string(substring))
				}
			} else {
				break
			}
		}
	}

	results <- substringList
}
