package work

import (
	"fmt"
	"testing"
)

func Test_chan(t *testing.T) {
	// 创建过滤器实例
	filter1 := &Filter1{}
	filter2 := &Filter2{}
	filter3 := &Filter3{}

	// 设置过滤器链
	filter1.SetNext(filter2)
	filter2.SetNext(filter3)

	// 创建请求和响应
	request := &Request{Data: "Request"}
	response := &Response{Data: "Response"}

	// 执行过滤器链
	filter1.Execute(request, response)

	// 输出最终的请求和响应数据
	fmt.Println("Final Request Data:", request.Data)
	fmt.Println("Final Response Data:", response.Data)

}
