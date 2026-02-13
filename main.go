package main

import (
	"sync"
)

type Recipent struct {
	Name  string
	Email string
}

func main() {

	recipentChannel := make(chan Recipent)

	go func() {
		loadRecipent("./email.csv", recipentChannel)

	}()
	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipentChannel, &wg)

	}

	wg.Wait()	

}
