package folder

import (
	"fmt"
	"tryDesignPattern/v2/struct/composite/component"
)

type Folder struct {
	components []component.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("遍历找单词 %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c component.Component) {
	f.components = append(f.components, c)
}
