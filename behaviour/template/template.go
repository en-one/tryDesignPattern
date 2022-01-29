package template

import "fmt"

/*
模板方法 template

   意图：父类定义骨架，子类实现某些细节。
   关键：通用步骤在抽象类中实现，变化的步骤在具体子类中实现
*/

type Cook interface {
	open()
	fire()
	cooke()
	outfire()
	close()
}

// CookStep 抽象类
type CookStep struct{}

// 通用步骤
func (CookStep) open()    { fmt.Println("打开开关") }
func (CookStep) fire()    { fmt.Println("开火") }
func (CookStep) cooke()   {}
func (CookStep) outfire() { fmt.Println("关火") }
func (CookStep) close()   { fmt.Println("关闭开关") }

// 做菜，交给具体的子类实现

type PaiGu struct {
	CookStep
}

func (x *PaiGu) cooke() {
	fmt.Println("做排骨")
}

type YouMaiCai struct {
	CookStep
}

func (j *YouMaiCai) cooke() {
	fmt.Println("做鸡蛋")
}

// 封装所有步骤
func DoCook(cook Cook) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outfire()
	cook.close()
}
