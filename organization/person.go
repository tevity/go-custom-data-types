package organization

type Identifiable interface {
	ID() string
}

type Person struct {
}

func (person Person) ID() string {
	return "1234"
}
