package main

import (
	"log"
	"math/rand"
	"time"
)

type OrderCommand interface {
	Type() string
	Name() string
}

type BaseCommand struct {
	name string
}

func (c BaseCommand) Type() string {
	return ""
}
func (c BaseCommand) Name() string {
	return c.name
}

type AppetizerCommand struct {
	BaseCommand
}

func (a AppetizerCommand) Type() string {
	return "appetizer"
}

type MainCommand struct {
	BaseCommand
}

func (a MainCommand) Type() string {
	return "main"
}

type DesertCommand struct {
	BaseCommand
}

func (a DesertCommand) Type() string {
	return "desert"
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
			log.Printf("Chief prepares %s: %s", order.Type(), order.Name())
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
	keys := []string{"appetizers", "main", "desert"}
	key := keys[r.Intn(len(keys))]
	arr := meals[key]
	chief.orders <- arr[r.Intn(len(arr))]
}

var (
	r     = rand.New(rand.NewSource(time.Now().UnixMilli()))
	meals = map[string][]OrderCommand{
		"appetizers": NewAppertizerCommand("Chicken Wign", "Crab Rangoon", "Fresh Rolls", "Egg Rolls"),
		"main":       NewMainCommand("Chicken Noodle Soup", "Beef Noodle Soup", "Grilled Pork"),
		"desert":     NewDesertCommand("Vanilla Ice Cream", "Fresh Berries", "Chocolate Cake"),
	}
)

func NewAppertizerCommand(names ...string) []OrderCommand {
	rs := make([]OrderCommand, 0)
	for _, name := range names {
		rs = append(rs, AppetizerCommand{
			BaseCommand: BaseCommand{
				name: name,
			},
		})
	}
	return rs
}
func NewMainCommand(names ...string) []OrderCommand {
	rs := make([]OrderCommand, 0)
	for _, name := range names {
		rs = append(rs, MainCommand{
			BaseCommand: BaseCommand{
				name: name,
			},
		})
	}
	return rs
}
func NewDesertCommand(names ...string) []OrderCommand {
	rs := make([]OrderCommand, 0)
	for _, name := range names {
		rs = append(rs, DesertCommand{
			BaseCommand: BaseCommand{
				name: name,
			},
		})
	}
	return rs
}

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
