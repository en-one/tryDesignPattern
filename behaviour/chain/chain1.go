package chain

import "fmt"

/**
责任链模式（Chain of Responsibility Pattern）：行为设计模式，它允许对象将请求沿处理程序链进行传递。
	程序链既可以处理请求，也可以继续将请求传递给链中的下一个处理程序。请求沿着处理程序链传递，直到请求被处理或到达处理程序链的末端。

以下是个例子
	责任链将负责验证请求、打印错误日志和处理请求本身,处理流程logger->validator->basic
*/

// 定义一个基础接口
type IService interface {
	HelloWorldChain(name string) (string, error)
}

// 定义一个基础打印服务
type basic struct{}

func (b basic) HelloWorldChain(name string) (string, error) {
	fmt.Println("HelloWorld method from basic")
	return fmt.Sprintf("Hello World from %s", name), nil
}

// 定义一个验证器，
type validator struct{ next IService }

func (v validator) HelloWorldChain(name string) (string, error) {
	if len(name) <= 3 {
		return "", fmt.Errorf("name length must be greater than 3")
	}
	fmt.Println("HelloWorld method from validator")
	return v.next.HelloWorldChain(name)
}

// 定义一个日志服务 logger->validator->basic
type logger struct{ next IService }

func (l logger) HelloWorldChain(name string) (string, error) {
	res, err := l.next.HelloWorldChain(name)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	fmt.Println("HelloWorld method from logger")
	return res, err
}

func NewIService() IService {
	return logger{
		next: validator{
			next: basic{},
		},
	}
}
