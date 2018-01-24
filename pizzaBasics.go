package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"

	"github.com/satori/go.uuid"
)

const (
	orderTime       = 1
	preparationTime = 5
	bakeTime        = 10
)

type PizzaOrder struct {
	ID           string
	CustomerName string
	Toppings     []string
}

func main() {
	// extra tracing stuff
	traceFile, err := os.Create("./pizzaBasic.trace")
	if err != nil {
		panic(err)
	}
	trace.Start(traceFile)
	defer func() {
		trace.Stop()
		err := traceFile.Close()
		if err != nil {
			log.Printf("error closing trace file. %s", err)
		}
	}()
	orders := []*PizzaOrder{
		{CustomerName: "Asaf", Toppings: []string{"double cheese", "pepperoni"}},
		{CustomerName: "Asaf", Toppings: []string{"double cheese", "pepperoni"}},
		{CustomerName: "Asaf", Toppings: []string{"double cheese", "pepperoni"}},
		{CustomerName: "Asaf", Toppings: []string{"double cheese", "pepperoni"}},
	}

	// queue up people
	orderChan := make(chan *PizzaOrder, len(orders))
	for _, order := range orders {
		orderChan <- order
	}
	// since ther's no further input, we can close the channel
	close(orderChan)

	// create the output channel for the order step
	preperationChan := make(chan *PizzaOrder, 10)
	go func() {
		// close our channel once steopped
		defer close(preperationChan)
		for customerOrder := range orderChan {
			preperationChan <- takeOrder(customerOrder)
		}
	}()

	// prepare pizzas
	bakingChan := make(chan *PizzaOrder, 10)
	go func() {
		defer close(bakingChan)
		for preparedPizza := range preperationChan {
			bakingChan <- prepare(preparedPizza)
		}
	}()

	// bake pizzas
	serveChan := make(chan *PizzaOrder, 10)
	go func() {
		defer close(serveChan)
		for pizzaIDToBake := range bakingChan {
			serveChan <- bake(pizzaIDToBake)
		}
	}()

	for readyPizza := range serveChan {
		fmt.Printf("%s, Pizza Pizza Pie!!!!", readyPizza.CustomerName)
	}
}

func prepare(order *PizzaOrder) *PizzaOrder {
	log.Printf("Preparing pizza with %+v", order.Toppings)
	time.Sleep(time.Second * preparationTime)
	log.Printf("Pizza %s is ready for baking!", order.ID)
	return order
}

func bake(order *PizzaOrder) *PizzaOrder {
	log.Printf("Baking pizza: %s", order.ID)
	time.Sleep(time.Second * preparationTime)
	log.Printf("Pizza %s is ready!", order.ID)
	return order
}

func takeOrder(order *PizzaOrder) *PizzaOrder {
	log.Printf("Taking an order from %s", order.CustomerName)
	time.Sleep(time.Second * orderTime)
	order.ID = uuid.Must(uuid.NewV4()).String()
	log.Printf("Placing order with %v", order.Toppings)
	return order
}
