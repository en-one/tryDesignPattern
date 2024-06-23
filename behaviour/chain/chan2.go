package chain

import "fmt"

/*
	以下是一个医院的例子,医院应用的责任链模式例子。 医院中会有多个部门， 如：

		前台
		医生
		药房
		收银
	病人来访时，会先去前台；然后看医生；取药；结账；病人需要通过一条部门链.

	此模式适用于多个候选选项处理相同请求时，
	1、不希望客户端选择接收者
	2、客户端同接收者解耦，客户端只需要链中的首个元素即可
*/

// 病人
type Patient struct {
	name              string // 病人名字
	registrationDone  bool   // 是否挂号
	doctorCheckUpDone bool   // 是否看医生
	medicineDone      bool   // 是否取药
	paymentDone       bool   // 是否结账
}

// 部门接口，每个部门处理自己的任务，同时可以将任务向下一个链路传递
type IDepartment interface {
	execute(*Patient)    // 执行任务
	setNext(IDepartment) // 设置下一个部门
}

// 前台
type Reception struct{ Next IDepartment }

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println(" 病人已经挂好了，去看医生吧")
		r.Next.execute(p)
		return
	}
	fmt.Println(" 病人来了，请先挂号")
	p.registrationDone = true
	r.Next.execute(p)
}

func (r *Reception) setNext(next IDepartment) {
	r.Next = next
}

// 医生
type Doctor struct{ next IDepartment }

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println(" 医生已经看过了，去取药吧")
		d.next.execute(p)
		return
	}
	fmt.Println(" 医生给病人看病")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next IDepartment) {
	d.next = next
}

// 药房
type Medical struct{ next IDepartment }

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("已经拿过药了")
		m.next.execute(p)
		return
	}
	fmt.Println("药房给病人药")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next IDepartment) {
	m.next = next
}

// 收银
type Cashier struct {
	next IDepartment
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("已经结账了")
	}
	fmt.Println("收银给病人结账")
}

func (c *Cashier) setNext(next IDepartment) {
	c.next = next
}
