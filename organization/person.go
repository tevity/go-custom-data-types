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
	Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
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
