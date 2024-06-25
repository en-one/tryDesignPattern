package chain

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

type patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

// IDepartment 接口定义了处理病人请求的方法,每个部门处理自己的任务，同时可以将任务向下一个链路传递
type IDepartment interface {
	Execute(*patient) error          // 执行任务
	SetNext(IDepartment) IDepartment // 链路传递
	Do(*patient) error               // 具体任务逻辑
}

// Next 结构体用于实现责任链模式中的下一个处理器
type Next struct {
	nextHandler IDepartment
}

func (n *Next) SetNext(handler IDepartment) IDepartment {
	n.nextHandler = handler
	return handler
}

func (n *Next) Do(patient *patient) error { return nil }

func (n *Next) Execute(patient *patient) (err error) {
	// 需要哨兵节点
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(patient); err != nil {
			return
		}

		return n.nextHandler.Execute(patient)
	}

	return
}

// StartHandler 不做操作，作为第一个Handler向下转发请求。如果没有这个哨兵节点，会导致最后一个对象的next无法执行
type StartHandler struct {
	Next
}
