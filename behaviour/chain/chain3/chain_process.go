package chain

import "fmt"

// Reception 挂号处处理器
type Reception struct{ Next }

func (r *Reception) Do(p *patient) (err error) {
	if p.RegistrationDone {
		fmt.Println("病人已经挂号了 去看医生吧")
		return
	}
	fmt.Println("病人来了，请先挂号")
	p.RegistrationDone = true
	return
}

// Clinic 诊室处理器--用于医生给病人看病
type Clinic struct{ Next }

func (d *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("医生已经看过了，去取药吧")
		return
	}
	fmt.Println("医生给病人看病")
	p.DoctorCheckUpDone = true
	return
}

// Pharmacy 药房处理器
type Pharmacy struct{ Next }

func (m *Pharmacy) Do(p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("已经拿过药了")
		return
	}
	fmt.Println("病房给病人药")
	p.MedicineDone = true
	return
}

// Cashier 收费处处理器
type Cashier struct{ Next }

func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("已经结账了")
		return
	}
	fmt.Println("收银给病人结账")
	p.PaymentDone = true
	return
}
