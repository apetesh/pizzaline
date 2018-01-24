package main

import (
	"log"
	"sync"
	"time"
)

func bake(input chan *PizzaOrder, numOfWorkers, bufferSize int) chan *PizzaOrder {
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
				log.Printf("Baking pizza: %s", order.ID)
				time.Sleep(speed * bakeTime)
				log.Printf("Pizza %s is ready!", order.ID)
				outputChan <- order
			}
			wg.Done()
		}()
	}
	return outputChan
}
