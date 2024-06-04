package main

import (
	"tryDesignPattern/v2/struct/composite/file"
	"tryDesignPattern/v2/struct/composite/folder"
)

//
/*

	组合：将对象组合成树状结构，并且能够像使用独立对象一样使用他们

	以下描述了一个文件系统的例子：
	1. 【文件系统】由文件和文件夹组成，某些情况下文件合文件夹可以被视为相同的对象。（组合模式发挥）
	2. 当在【文件系统】中搜索特殊特定的关键词时，这一操作会同时作用于文件合文件夹。
	3. 对于文件，其智慧查看文件的内容；对于文件夹，则会在内部的所有文件中查找关键词
*/
func main() {
	file1 := &file.File{Name: "File1"}
	file2 := &file.File{Name: "File2"}
	file3 := &file.File{Name: "File3"}

	folder1 := &folder.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &folder.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
