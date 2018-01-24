package main

import (
	"log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

func takeOrder(input chan *PizzaOrder, numOfWorkers, bufferSize int) chan *PizzaOrder {
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
				log.Printf("Taking an order from %s", order.CustomerName)
				time.Sleep(speed * orderTime)
				order.ID = uuid.Must(uuid.NewV4()).String()
				log.Printf("Placing order with %v", order.Toppings)
				order.OrderTime = time.Now()
				outputChan <- order
			}
			wg.Done()
		}()
	}
	return outputChan
}
