package main

/*
go build -gcflags=-m
*/
func main() {
	NewPersonWithPointer(1, "12312", "123123")
}

type Person struct {
	Age   int
	Name  string
	Hobby string
}

func NewPersonWithPointer(age int, name, hobby string) *Person {
	person := new(Person)

	person.Age = age
	person.Name = name
	person.Hobby = hobby

	return person
}
