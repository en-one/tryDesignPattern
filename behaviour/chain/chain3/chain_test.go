package chain

import (
	"fmt"
	"testing"
)

func Test_chan(t *testing.T) {
	patientHealthHandler := StartHandler{}
	//
	patient := &patient{Name: "abc"}
	// 设置病人看病的链路
	patientHealthHandler.SetNext(&Reception{}). // 挂号
							SetNext(&Clinic{}).  // 诊室看病
							SetNext(&Cashier{}). // 收费处交钱
							SetNext(&Pharmacy{}) // 药房拿药
	// 还可以无限扩展，比如中间加入化验科化验，图像科拍片等等

	// 执行上面设置好的业务流程
	if err := patientHealthHandler.Execute(patient); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")

}
