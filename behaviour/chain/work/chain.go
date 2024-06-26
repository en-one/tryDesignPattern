package work

import "fmt"

/*

	1、我有过滤器1，2，3
	2、我有两个类，request，reponse
	3、要求request 执行过滤器1，2，3；response执行过滤器3，2，1
*/

type Request struct {
	Data string
}

type Response struct {
	Data string
}

// IFilter 过滤器接口 (核心)
type IFilter interface {
	Execute(request *Request, response *Response)
	SetNext(filter IFilter)
}

// 基类，用来组合实现过滤器链
type BaseFilter struct {
	next IFilter
}

func (b *BaseFilter) SetNext(filter IFilter) {
	b.next = filter
}

func (b *BaseFilter) Execute(request *Request, response *Response) {
	if b.next != nil {
		b.next.Execute(request, response)
	}
}

// 第一个过滤器
type Filter1 struct {
	BaseFilter
}

func (f *Filter1) Execute(request *Request, response *Response) {
	fmt.Println("Filter1 processing request")
	request.Data += " -> Filter1"
	if f.next != nil {
		f.next.Execute(request, response)
	}
	fmt.Println("Filter1 processing response")
	response.Data += " -> Filter1"
}

// 第二个过滤器
type Filter2 struct {
	BaseFilter
}

func (f *Filter2) Execute(request *Request, response *Response) {
	fmt.Println("Filter2 processing request")
	request.Data += " -> Filter2"
	if f.next != nil {
		f.next.Execute(request, response)
	}
	fmt.Println("Filter2 processing response")
	response.Data += " -> Filter2"
}

// 第三个过滤器
type Filter3 struct {
	BaseFilter
}

func (f *Filter3) Execute(request *Request, response *Response) {
	fmt.Println("Filter3 processing request")
	request.Data += " -> Filter3"
	if f.next != nil {
		f.next.Execute(request, response)
	}
	fmt.Println("Filter3 processing response")
	response.Data += " -> Filter3"
}
