package organization

type Identifiable interface {
	ID() string
}

type Person struct {
	FirstName string
	LastName string
}

func (person Person) ID() string {
	return "1234"
}
