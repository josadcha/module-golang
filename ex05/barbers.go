package main

import (
	"fmt"
	"time"
)

type Barber struct {
	shop  chan *Customer
	seats chan *Customer
	sleep chan struct{}
}

type Customer struct {
	id int
}

func new_barber(maxwait int) *Barber {
	bar := &Barber{make(chan *Customer, 1), make(chan *Customer, maxwait), make(chan struct{}, 1)}
	go bar.barber()
	return bar
}

func new_cust(id int) *Customer {
	return &Customer{id}
}

func (b *Barber) barber() {
	b.sleep <- struct{}{}
	fmt.Println("Barber is sleeping")
	for {
		c := <-b.shop
		b.CutHair(c)
		select {
		case c := <-b.seats:
			b.shop <- c
			fmt.Println("Barber takes the customer", c.id)
		default:
			b.sleep <- struct{}{}
			fmt.Println("Barber is sleeping")
		}
	}
}

func (b *Barber) CutHair(c *Customer) {
	fmt.Println("Barber cuts the customer", c.id)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Cuts end", c.id)
}

func (c *Customer) customer(b *Barber) {
	fmt.Println("Customer", c.id, "enters the barbershop.")
	select {
	case <-b.sleep:
		b.shop <- c
		fmt.Println("Customer", c.id, "wakes up the barber")
	default:
		select {
		case b.seats <- c:
			fmt.Println("Customer", c.id, "sits in the lobby.")
		default:
			fmt.Println("Customer", c.id, "leaves the barbershop.")
		}
	}
}

func main() {
	ncust := 5
	b := new_barber(3)

	for i := 0; i < ncust; i++ {
		c := new_cust(i)
		time.Sleep(500 * time.Millisecond)
		go c.customer(b)
	}
	time.Sleep(2 * time.Second)
}
