package main

import "log"

// interface
type Observer interface {
	update(string)
	getId() int
}

type Subject interface {
	register(Observer)
	unRegister(Observer)
	notifyAll()
}

type Customer struct {
	Id int
}

func (c Customer) update(msg string) {
	log.Printf("user %v receive msg from shop: %v", c.getId(), msg)
}

func (c Customer) getId() int {
	return c.Id
}

type Shop struct {
	subscribers map[int]Observer
}

func (c *Shop) register(s Observer) {
	c.subscribers[s.getId()] = s
}

func (c *Shop) unRegister(s Observer) {
	delete(c.subscribers, s.getId())
}

func (c *Shop) notifyAll() {
	log.Println("shop prepares for incoming event")
	for _, s := range c.subscribers {
		s.update("invitation to HOT event on YYYY/MM/DD")
	}
}

func main() {
	shop := Shop{
		subscribers: make(map[int]Observer),
	}
	for i := 0; i < 10; i++ {
		shop.register(Customer{Id: i + 1})
	}
	shop.notifyAll()
}
