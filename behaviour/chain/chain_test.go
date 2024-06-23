package chain

import (
	"fmt"
	"testing"
)

func Test_chan1(t *testing.T) {
	s := NewIService()

	res, err := s.HelloWorldChain("煎鱼13231")
	fmt.Println(res, err)

	// 会返回错误输出
	res, err = s.HelloWorldChain("12")
	fmt.Println(res, err)
}

// 病人看病的例子。前台->医生->药房->收银.
func Test_chan2(t *testing.T) {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "josiah"}
	reception.execute(patient)

}
