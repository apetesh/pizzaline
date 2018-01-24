package main

import (
	"log"
	"sync"
	"time"
)

func serve(input chan *PizzaOrder, numOfWorkers, bufferSize int) chan *PizzaOrder {
	outputChan := make(chan *PizzaOrder, bufferSize)
	wg := &sync.WaitGroup{}
	wg.Add(numOfWorkers)
	go func() {
		wg.Wait()
		close(outputChan)
	}()
	for i := 0; i < numOfWorkers; i++ {
		go func() {
			for order := range input {
				log.Printf("Serving pizza: %s", order.ID)
				outputChan <- order
				order.ServeTime = time.Now()
			}
			wg.Done()
		}()
	}
	return outputChan
}
