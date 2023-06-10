package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type FindStruct struct {
	text string
	start int
	end int
	maxLen *int
}

func FindLongestRepeatingSubstring(text string, numOfWorkers int) string {
	maxLen := 1

	jobs := make(chan FindStruct, numOfWorkers)
	results := make(chan string, numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		go worker(jobs, results);
	}

	for i := 0; i < len(text) - maxLen; i++ {
		jobs <- FindStruct{text: text, start: i, end: len(text) - maxLen, maxLen: &maxLen}
	}
	close(jobs)

	result := ""
	for i := 0; i < numOfWorkers; i++ {
		temp := <- results
		if len(temp) > len(result) {
			result = temp
		}
	}

	return result
}

func worker(jobs <-chan FindStruct, results chan<- string) {
	result := ""
	for job := range jobs {
		for i := job.start + *job.maxLen; i < job.end; i++ {
			if strings.Contains(job.text[i:], job.text[job.start:i]) {
				if i - job.start > len(result) {
					*job.maxLen = i - job.start
					result = job.text[job.start:i]
				}
			} else {
				break
			}
		}
	}
	results <- result
}

func main() {
	filePath := flag.String("path", "data.txt", "path to the file to be processed")
	numOfWorkers := flag.Int("num-workers", runtime.NumCPU(), "amount of workers")
	flag.Parse()

	data, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalln(err)
	}
	text := string(data)

	timeStart := time.Now()
	longestSubstr := FindLongestRepeatingSubstring(text, *numOfWorkers)
	elapsed := time.Since(timeStart)

	fmt.Println("The longest repeating substring is:", longestSubstr)
	fmt.Printf("Calculated in %s\n", elapsed)
}
