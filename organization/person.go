package organization

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
}

func NewPerson(firstName, lastName string) Person {
	return Person {
		firstName: firstName,
		lastName: lastName,
	}
}

func (person Person) ID() string {
	return "1234"
}
