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
	Execute(*patient) error
	SetNext(IDepartment) IDepartment
	Do(*patient) error
}

// Next 结构体用于实现责任链模式中的下一个处理器
type Next struct {
	nextHandler IDepartment
}

func (n *Next) SetNext(handler IDepartment) IDepartment {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) (err error) {
	// 由于go无继承的概念, 只能用组合，组合跟继承不一样，这里如果Next 实现了 Do 方法，那么匿名组合它的具体处理类型，执行Execute的时候，调用的还是内部Next对象的Do方法
	// 调用不到外部类型的 Do 方法，所以 Next 不能实现 Do 方法
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(patient); err != nil {
			return
		}

		return n.nextHandler.Execute(patient)
	}

	return
}

func (n *Next) Do(c *patient) (err error) {
	return
}

// StartHandler 不做操作，作为第一个Handler向下转发请求
// Go 语法限制，抽象公共逻辑到通用Handler后，并不能跟继承一样让公共方法调用不通子类的实现
type StartHandler struct {
	Next
}

// Do 空Handler的Do
// func (h *StartHandler) Do(c *patient) (err error) {
// 	// 空Handler 这里什么也不做 只是载体 do nothing...
// 	return
// }
