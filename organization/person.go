package organization

import (
	"errors"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	Name
	twitterHandle TwitterHandle
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (person *Person) ID() string {
	return "1234"
}

func (person *Person) SetTwitterHandle(handle TwitterHandle) error {
	if len(handle) > 0 && !handle.Valid() {
		return errors.New("Twitter handle must start with @")
	}

	person.twitterHandle = handle
	return nil
}

func (person *Person) TwitterHandle() TwitterHandle {
	return person.twitterHandle
}
