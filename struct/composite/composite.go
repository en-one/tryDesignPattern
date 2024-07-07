package composite

import (
	"fmt"
)

/*
组合：将对象组合成树状结构，并且能够像使用独立对象一样使用他们

以下描述了一个文件系统的例子：
1. 【文件系统】由文件和文件夹组成，某些情况下文件合文件夹可以被视为相同的对象。（组合模式发挥）
2. 当在【文件系统】中搜索特殊特定的关键词时，这一操作会同时作用于文件合文件夹。
3. 对于文件，其智慧查看文件的内容；对于文件夹，则会在内部的所有文件中查找关键词
*/
type INode interface {
	Search(string)
}

// 文件
type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf(" Searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) getName() string {
	return f.Name
}

// 文件夹
type Folder struct {
	children []INode
	Name     string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("遍历找单词 %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.children {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c INode) {
	f.children = append(f.children, c)
}
