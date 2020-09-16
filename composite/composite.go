package main

import "fmt"

//安全模式 树叶组件和树枝组件不同
func main() {

	root := &Directory{Name: "根目录"}
	first := &Directory{Name: "第一层目录"}
	second := &Directory{Name: "第二层目录"}
	index1 := &Index{Name: "指标1"}
	//构建整体
	root.Add(first)
	first.Add(second)
	second.Add(index1)
	//访问
	first.Operation()
	index1.Operation()
}

type IDirectory interface {
	Operation()
}

type Index struct {
	Name string
}

func (index *Index) Operation() {
	fmt.Println("指标：", index.Name)
}

type Directory struct {
	Name     string
	Children []IDirectory
}

func (d *Directory) Operation() {
	fmt.Println("目录：", d.Name)
}

func (d *Directory) Add(child IDirectory) {
	d.Children = append(d.Children, child)
}

func (d *Directory) GetChildren() []IDirectory {
	return d.Children
}
