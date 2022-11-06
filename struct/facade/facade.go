package facade

import "fmt"

// 提供一个邮寄信件功能： 写信，填写地址，检查，装信封，投递
type SampleLetters struct {
	letter string
}

func (l *SampleLetters) WriteLetters(content string) {
	l.letter = fmt.Sprintf("写信内容为:%s", content)
}

func (l *SampleLetters) AddAddress(address string) {
	l.letter = l.letter + "\n" + fmt.Sprintf("邮寄至: %s", address)
}

func (l *SampleLetters) PutInToEnvelope() {
	fmt.Println("信件存放到信封中...")
}

func (l *SampleLetters) SendLetters() {
	// fmt.Println("信件内容:", l.letter)
	fmt.Println("信件已发送...")
}

// 门面类
type SampleFacade struct {
	SampleLetters    // 基础发送信件
	PoliceInspection // 新增扩展，检查类
}

func NewSampleFacade() *SampleFacade {
	return &SampleFacade{}
}

// 如果此时需要新拓展功能，对信的内容进行检查
// 则继续添加SampleFacade的组合类
type PoliceInspection struct {
}

func (p *PoliceInspection) CheckLetterSecurity(content string) bool {
	// 如果检查内通通过，则返回true
	fmt.Printf("被检查内容: %s\n", content)
	fmt.Println("内容检查通过...")
	return true
}

// 将相关联性的子系统的集合进行综合，完成一系列的业务逻辑
func (f *SampleFacade) sendLetter(context, address string) {
	// 此处进行业务逻辑的实现
	// 1. 写信
	f.WriteLetters(context)
	// 2. 写地址
	f.AddAddress(address)
	// 此处新增检查信件安全性
	ok := f.CheckLetterSecurity(f.letter)
	if ok {
		fmt.Println("this is safe!")
	} else {
		fmt.Println("this is danger!")
		return
	}
	// 3. 封装
	f.PutInToEnvelope()
	// 4.发送
	f.SendLetters()
}
