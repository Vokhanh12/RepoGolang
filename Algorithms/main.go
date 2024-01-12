package main

import (
	"fmt"
	"time"
)

// "github.com/vokhanh12/DataStructures"

const (
	numberOfURLs    = 100
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	queue := make(chan int, 10)

	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	return queue
}

func main() {
	queue := startQueue()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Minute * 5)
}

func test(numDBs int) {
	dbChan := getDatabaseChanels(numDBs)
	fmt.Println("Waiting for database", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Println("All database all online!")
}

func waitForDbs(numDbs int, dbchan chan struct{}) {
	for i := 0; i < numDbs; i++ {
		<-dbchan
	}
}

func getDatabaseChanels(numDBs int) chan struct{} {
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {

			ch <- struct{}{}
			fmt.Println("database %v is online", i+1)
		}
	}()

	return ch
}
