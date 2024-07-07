package prototype

/*

	原型：使你能够复制对象，甚至是复杂对象。而又无需使代码依赖他们所属的累。
		 - 所有原型类都必须有一个通用的接口。使得即使对象所属类未知的情况下也能复制对象
		 - 原型类可以生成自身的完整副本，（包括私有对象）

	以下描述了一个文件系统的例子（组合模式也举例了文件系统）：
	1. 操作系统中的文件系统是递归的，文件夹中包含文件和文件夹。
	2.
*/

type INode interface {
	printe(string)
	clone() INode // 通用的复制接口
}

// 文件
type File struct {
	name string
}

func (f *File) printe(prefix string) {
	println(prefix + f.name)
}

func (f *File) clone() INode {
	return &File{name: f.name + "_clone"}
}

// 文件夹
type Floder struct {
	children []INode
	name     string
}

func (f *Floder) printe(prefix string) {
	println(prefix + f.name)
	for _, child := range f.children {
		child.printe(prefix + prefix)
	}
}

func (f *Floder) clone() INode {
	cloneFloder := &Floder{name: f.name + "_clone"}
	var tempChildren []INode
	for _, child := range f.children {
		tempChildren = append(tempChildren, child.clone())
	}
	cloneFloder.children = tempChildren
	return cloneFloder
}
