package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status int
	mu     sync.Mutex
}

const (
	Pending = iota
	Processed
	Shipped
	Delivered
	Completed
)

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: Pending}
	}

	return orders
}

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		order.Status = Processed
		outChan <- order
	}

}

func updateOrder(order *Order) {
	order.mu.Lock()
	defer order.mu.Unlock()
	time.Sleep(
		time.Duration(rand.Intn(300)) *
			time.Millisecond,
	)
	status := rand.Intn(4)
	order.Status = status
	fmt.Printf(
		"Updated order %d status: %d\n",
		order.ID, status,
	)
}

func reportOrderStatus(orders []*Order) {
	time.Sleep(1 * time.Second)
	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %d\n",
			order.ID, order.Status,
		)
	}
	fmt.Println("---------------------------")
	fmt.Println()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	orderCount := 20
	orderChan := make(chan *Order, orderCount)
	processedChan := make(chan *Order, orderCount)

	go func() {
		defer wg.Done()
		defer close(orderChan)
		for _, order := range generateOrders(orderCount) {
			orderChan <- order // blocking until receiver due to make(chan *Order)
		}

		fmt.Println("Done with generating orderes")
	}()

	go processOrders(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()

		for {
			select {
			case processOrder, ok := <-processedChan:
				if !ok {
					fmt.Println("Processing channel closed")
					return
				}
				fmt.Printf("Processed order %d with status: %d\n", processOrder.ID, processOrder.Status)

			case <-time.After(300 * time.Millisecond):
				// if we are stuck waiting for maximum wait
				fmt.Println("Timeout waiting for operations.")
				return

			}
		}
	}()

	wg.Wait()

	fmt.Println("Complete")
}
