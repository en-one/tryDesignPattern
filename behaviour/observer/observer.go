package observer

import "fmt"

/*
观察者模式 template

   意图：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所哟依赖于他的对象都得到通知并被自动更新
   关键：被观察者 有一个集合 存放 观察者
   示例：海绵宝宝 和 派大星 观察 蟹老板
		顾客A 和 顾客B 关注报社， 报社有一个集合，存了报社A和报社B，报社有新消息，则通知A和B
*/

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
