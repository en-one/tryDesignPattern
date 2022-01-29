package observer

import "testing"

func TestObserver(t *testing.T) {
	customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}
	office.addCustomer(customerA)
	office.addCustomer(customerB)

	office.NotifyAllCustomer()
}
