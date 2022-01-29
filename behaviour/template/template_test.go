package template

import "testing"

func TestTemplate(t *testing.T) {
	paigu := PaiGu{}
	DoCook(&paigu)

	ymc := YouMaiCai{}
	DoCook(&ymc)

}
