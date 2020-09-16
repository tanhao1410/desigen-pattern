package main

import "fmt"

func main() {

	//读者
	read1 := &Reader{Name: "read1"}
	read2 := &Reader{Name: "read2"}

	newspaper := new(MorningNewspaper)
	newspaper.AddReader(read1)
	newspaper.AddReader(read2)

	newspaper.PublishArticle(&Article{
		Title:   "这是标题",
		Content: "这是内容",
	})

}

type Article struct {
	Title   string
	Content string
}

type IReader interface {
	Read(article *Article)
}

type Reader struct {
	Name string
}

func (reader *Reader) Read(article *Article) {
	fmt.Println(reader.Name, "读了：", article.Title, "内容是：", article.Content)
}

type Newspaper interface {
	AddReader(read IReader)
	PublishArticle(article *Article)
}

type MorningNewspaper struct {
	readers map[IReader]bool
}

func (self *MorningNewspaper) AddReader(read IReader) {
	if self.readers == nil {
		self.readers = make(map[IReader]bool)
	}
	if _, ok := self.readers[read]; !ok {
		self.readers[read] = true
	}
}

func (self *MorningNewspaper) PublishArticle(article *Article) {
	for reader, _ := range self.readers {
		reader.Read(article)
	}
}
