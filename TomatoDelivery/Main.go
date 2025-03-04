package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FoodOrder struct {
	ID       int
	FoodItem string
	Quantity int
	Address  string
}

type DeliveryProcessor struct {
	orders    chan FoodOrder
	orderData map[int]string
	mu        sync.Mutex
	wg        sync.WaitGroup
}

func (dp *DeliveryProcessor) ProcessOrder(order FoodOrder) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) // Simulate delivery time

	// Lock the map while updating the status of the order
	dp.mu.Lock()
	dp.orderData[order.ID] = fmt.Sprintf("Order %d (%s) - Delivered to: %s ✅", order.ID, order.FoodItem, order.Address)
	dp.mu.Unlock()

	// Print Order ID and details when processed
	fmt.Printf("Processing Order ID: %d - %s (Qty: %d) for delivery to: %s... Done!\n", order.ID, order.FoodItem, order.Quantity, order.Address)
	dp.wg.Done()
}

func (dp *DeliveryProcessor) Worker() {
	for order := range dp.orders {
		dp.ProcessOrder(order)
	}
}

func (dp *DeliveryProcessor) AddOrder(order FoodOrder) {
	dp.wg.Add(1)
	dp.orders <- order
	fmt.Printf("Order %d: %s (Qty: %d) added to delivery queue for %s...\n", order.ID, order.FoodItem, order.Quantity, order.Address)
}

func (dp *DeliveryProcessor) Start(numWorkers int) {
	dp.orderData = make(map[int]string)  // Initialize order data map
	dp.orders = make(chan FoodOrder, 10) // Create a buffered channel

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go dp.Worker()
	}
}

func (dp *DeliveryProcessor) Stop() {
	close(dp.orders)
	dp.wg.Wait()
	fmt.Println("\nAll food orders delivered successfully!")
}

func (dp *DeliveryProcessor) DisplayOrderStatus() {
	fmt.Println("\n🍕 Food Delivery Status:")
	for id, status := range dp.orderData {
		// Use order ID here to display the final status
		fmt.Printf("Order ID %d: %s\n", id, status)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	processor := &DeliveryProcessor{}
	processor.Start(3) // Start with 3 worker goroutines for delivery

	// Sample food orders
	orders := []FoodOrder{
		{ID: 1, FoodItem: "Pizza", Quantity: 2, Address: "123 Main St"},
		{ID: 2, FoodItem: "Burger", Quantity: 1, Address: "456 Oak Ave"},
		{ID: 3, FoodItem: "Pasta", Quantity: 3, Address: "789 Pine Blvd"},
		{ID: 4, FoodItem: "Sushi", Quantity: 2, Address: "101 Maple Dr"},
		{ID: 5, FoodItem: "Salad", Quantity: 1, Address: "202 Elm St"},
	}

	// Add orders to the delivery queue
	for _, order := range orders {
		processor.AddOrder(order)
	}

	// Stop processing after all orders are sent
	processor.Stop()

	// Display final food delivery statuses
	processor.DisplayOrderStatus()
}
