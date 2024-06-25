package chain

import (
	"testing"
)

// 病人看病的例子。前台->医生->药房->收银.
func Test_chan(t *testing.T) {
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
