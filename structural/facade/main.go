package main

import "log"

type Shop struct {
	CartService
	InventoryService
	OrderService
	PaymentService
	ShipService
	NotificationService
}

func (s Shop) AddItemToCart(item string) {
	log.Printf("User add item %v", item)
	if s.InventoryService.CheckInventory(item) {
		s.CartService.Add(item)
		return
	}
}

func (s Shop) Checkout() {
	if len(s.CartService.items) == 0 {
		log.Println("- there is not any items to checkout")
		return
	}
	log.Println("User do checkout")
	order := s.OrderService.PlaceOrder(s.CartService.items)
	s.PaymentService.DoPayment(order)
	s.OrderService.ConfirmOrder(order)
	s.ShipService.Ship(order)
	s.NotificationService.Notify(order)
}

type CartService struct {
	items []string
}

func (c *CartService) Add(item string) {
	log.Printf("- save item on cart")
	c.items = append(c.items, item)
}

type InventoryService struct {
}

func (i InventoryService) CheckInventory(item string) bool {
	log.Printf("- check inventory to ensure that it has enough item %v", item)
	return true
}

type OrderService struct {
}

func (o OrderService) PlaceOrder(items []string) string {
	order := "#001"
	log.Printf("- place order %v with state PENDING with %d items: %v", order, len(items), items)
	return order
}

func (o OrderService) ConfirmOrder(order string) {
	log.Printf("- confirm order %v", order)
}

type ShipService struct {
}

func (s ShipService) Ship(order string) {
	log.Printf("- create shipment label for order %v and delivery to vendor", order)
}

type PaymentService struct {
}

func (p PaymentService) DoPayment(order string) {
	log.Printf("- pay for order %v", order)
}

type NotificationService struct {
}

func (n NotificationService) Notify(order string) {
	log.Printf("- send confirmation email of order %s to user", order)
}

func main() {
	s := Shop{
		CartService: CartService{
			items: make([]string, 0),
		},
		InventoryService:    InventoryService{},
		OrderService:        OrderService{},
		PaymentService:      PaymentService{},
		ShipService:         ShipService{},
		NotificationService: NotificationService{},
	}
	s.AddItemToCart("IPhone 14 Pro Max")
	s.AddItemToCart("Airpods Pro 2nd Generation")
	s.Checkout()
}
