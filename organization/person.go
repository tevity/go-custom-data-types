package organization

import (
	"errors"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type Person struct {
	Name
	twitterHandle TwitterHandle
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
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
