package chain

import (
	"fmt"
)

type patient struct {
	Name string
}

type PatientHandler interface {
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

type Next struct {
	nextHandler PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(patient); err != nil {
			return err
		}
		return n.nextHandler.Execute(patient)
	}
	return nil
}

// 错误的地方：Next 结构体实现了 Do 方法
func (n *Next) Do(p *patient) error {
	fmt.Println("Next: Default Do method")
	return nil
}

// Reception 挂号处处理器
type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) error {
	fmt.Println("Reception: Handling patient")
	return nil
}

func check() {
	// 创建处理器
	reception := &Reception{}

	// 设置责任链
	reception.SetNext(nil)

	// 创建病人
	patient := &patient{Name: "张三"}

	// 处理病人
	if err := reception.Execute(patient); err != nil {
		fmt.Println("Error:", err)
	}
}
