package main

import (
	"log"
	"sync"
	"time"
)

func prepare(input chan *PizzaOrder, numOfWorkers, bufferSize int) chan *PizzaOrder {
	// create output channel
	outputChan := make(chan *PizzaOrder, bufferSize)

	// use a wait group to know when to close the output channel
	wg := &sync.WaitGroup{}
	wg.Add(numOfWorkers)

	go func() {
		// wait for all producers to drain and stop producing data downstream
		wg.Wait()
		// close output channel
		close(outputChan)
	}()

	for i := 0; i < numOfWorkers; i++ {
		// spawn a worker
		go func() {
			// read from input while it's open or not empty
			for order := range input {
				log.Printf("Preparing pizza with %+v", order.Toppings)
				time.Sleep(speed * preparationTime)
				log.Printf("Pizza %s is ready for baking!", order.ID)
				// push results to output channel
				outputChan <- order
			}
			// repott done to wait group
			wg.Done()
		}()
	}
	// return the output channel to caller
	return outputChan
}
