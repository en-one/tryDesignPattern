package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := File{name: "file1"}
	file2 := File{name: "file2"}
	file3 := File{name: "file3"}

	folder1 := Floder{name: "folder1", children: []INode{&file1}}
	folder2 := Floder{name: "folder2", children: []INode{&file2, &file3, &folder1}}

	fmt.Println("打印文件夹2内容")
	folder2.printe("	")

	cloneFloder := folder2.clone()
	fmt.Println("打印克隆文件夹2内容")
	cloneFloder.printe("	")
}

/*
	输出：
		打印文件夹2内容
		folder2
			file2
			file3
			folder1
					file1
		打印克隆文件夹2内容
		folder2_clone
			file2_clone
			file3_clone
			folder1_clone
					file1_clone
*/
