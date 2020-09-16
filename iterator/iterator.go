package main

import (
	"errors"
	"fmt"
)

func main() {
	book1 := &Book{Name: "book1"}
	book2 := &Book{Name: "book2"}

	shelf := new(MyBookShlef)
	shelf.Add(book1)
	shelf.Add(book2)

	iterator := shelf.GetIterator()
	for iterator.HasNext() {
		ele := iterator.Next().Ele
		book := ele.(*Book)
		fmt.Println(book.Name)
	}
}

//一个统一的接口，用Element包装了下，这样，该迭代器接口可以给给多的具体迭代器使用。
//因为golang中无泛型，因此相比java麻烦了些
type MyIterator interface {
	Next() *Element
	HasNext() bool
}
type Element struct {
	Ele interface{}
}

type BookShelf interface {
	GetIterator() MyIterator
	Add(book *Book)
}

type MyBookShlef struct {
	list []*Book // 私有
}

func (bookShelf *MyBookShlef) GetIterator() MyIterator {

	bookIterator := new(BookIterator)
	bookIterator.list = bookShelf.list
	bookIterator.index = 0
	return bookIterator
}

func (bookShelf *MyBookShlef) Add(book *Book) {
	bookShelf.list = append(bookShelf.list, book)
}

type BookIterator struct {
	//bookShelf BookShelf // 持有一个聚合类的引用，以访问它里面的数据,因为被私有化了，所以不能通过这种方式了
	list  []*Book // 持有聚合类里面数据的引用也是可以的
	index int
}

func (iterator *BookIterator) HasNext() bool {
	size := len(iterator.list)
	if size > 0 && iterator.index < size {
		return true
	}
	return false
}

func (iterator *BookIterator) Next() *Element {
	if !iterator.HasNext() {
		panic(errors.New("下标越界"))
	}
	book := iterator.list[iterator.index]
	iterator.index++
	return &Element{Ele: book}
}

type Book struct {
	Name string
}
