package main

import "fmt"

func main1() {
	person1 := new(Person)
	person1.Age = 12
	person1.Nation = "China"
	person1.Name = "test"

	//通过原型去创建
	person2 := person1.Clone()
	fmt.Println(person2.Name)
}

type Person struct {
	Name   string
	Age    uint8
	Nation string
}

func (person *Person) Clone() *Person {
	newPerson := new(Person)
	newPerson.Name = person.Name
	newPerson.Age = person.Age
	newPerson.Nation = person.Nation
	return newPerson
}
