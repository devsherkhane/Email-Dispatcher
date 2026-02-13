package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipent, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipent := range ch {
		
		fmt.Println(id, recipent)
	}

}
