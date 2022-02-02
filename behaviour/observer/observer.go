package observer

import "fmt"

type Customer interface {
	update()
}

type CustomerA struct {
}

func (c *CustomerA) update() {
	fmt.Println("i am A, i receive")
}

type CustomerB struct {
}

func (c *CustomerB) update() {
	fmt.Println("i am B, i receive")
}

type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newsPaperCome() {
	n.NotifyAllCustomer()
}

func (n *NewsOffice) NotifyAllCustomer() {
	for _, cus := range n.customers {
		cus.update()
	}
}
