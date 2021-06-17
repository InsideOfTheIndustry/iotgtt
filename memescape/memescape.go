package memescape

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

func NewPersonWithValue(age int, name, hobby string) Person {
	return Person{
		Age:   age,
		Name:  name,
		Hobby: hobby,
	}
}
