package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const (
	orderTime       = 1
	preparationTime = 3
	bakeTime        = 15

	cashiers    = 1
	pizzaMakers = 1
	ovens       = 1
)

type PizzaOrder struct {
	ID            string
	CustomerName  string
	Toppings      []string
	WaitStartTime time.Time
	OrderTime     time.Time
	BakedTime     time.Time
	ServeTime     time.Time
}

var (
	speed       = time.Duration(time.Millisecond) * 30
	numOfOrders = 100
)

func main() {
	log.SetOutput(ioutil.Discard)

	// queue up people
	orderChan := make(chan *PizzaOrder, numOfOrders)
	go func() {
		for _, order := range generateOrders(numOfOrders) {
			order.WaitStartTime = time.Now()
			orderChan <- order
		}
		close(orderChan)
	}()

	orderedPizzas := takeOrder(orderChan, cashiers, numOfOrders)
	preparedPizzas := prepare(orderedPizzas, pizzaMakers, numOfOrders)
	bakedPizzas := bake(preparedPizzas, ovens, numOfOrders)

	for pizzaToServe = range serve(bakedPizzas, 1, numOfOrders) {
		fmt.Printf("%s, come get your pizza.\n", pizzaToServe.CustomerName)
	}
}

func generateOrders(numOfOrders int) []*PizzaOrder {
	newOrders := make([]*PizzaOrder, 0, numOfOrders)
	for i := 0; i < numOfOrders; i++ {
		newOrders = append(newOrders, &PizzaOrder{
			CustomerName: fmt.Sprintf("Customer #%d", i),
			Toppings:     []string{"some", "toppings"},
		})
	}
	return newOrders
}
