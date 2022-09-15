package main

import (
	"log"
	"math/rand"
	"time"
)

type OrderCommand struct {
	Type string
	Name string
}

type Chief struct {
	orders    chan OrderCommand
	completed chan bool
}

func (c Chief) StartWork() {
	log.Printf("Chief start working")
	go func() {
		for {
			order, ok := <-c.orders
			if !ok {
				break
			}
			//process order
			log.Printf("Chief prepares %s: %s", order.Type, order.Name)
			<-time.After(time.Duration(r.Intn(50)+10) * time.Millisecond)
		}
		c.completed <- true
	}()
}

func (c Chief) StopWork() {
	close(c.orders)
	<-c.completed
	log.Printf("Chief stop working")
}

type Customer struct{}

func (c Customer) Order(chief Chief) {
	meals := map[string][]string{
		"appetizers": {"Chicken Wign", "Crab Rangoon ", "Fresh Rolls", "Egg Rolls"},
		"main":       {"Chicken Noodle Soup", "Beef Noodle Soup", "Grilled Pork"},
		"desert":     {"Vanilla Ice Cream", "Fresh Berries", "Chocolate Cake"},
	}
	keys := []string{"appetizers", "main", "desert"}
	key := keys[r.Intn(len(keys))]
	arr := meals[key]
	chief.orders <- OrderCommand{
		Type: key,
		Name: arr[r.Intn(len(arr))],
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixMilli()))

func main() {
	chief := Chief{
		orders:    make(chan OrderCommand, 100),
		completed: make(chan bool),
	}
	chief.StartWork()
	for i := 0; i < 100; i++ {
		c := Customer{}
		c.Order(chief)
	}
	chief.StopWork()
}
