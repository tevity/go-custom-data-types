package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandle string

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName     string
	lastName      string
	twitterHandle TwitterHandle
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (person *Person) FullName() string {
	return fmt.Sprintf("%s %s", person.firstName, person.lastName)
}

func (person *Person) ID() string {
	return "1234"
}

func (person *Person) SetTwitterHandle(handle TwitterHandle) error {
	if len(handle) > 0 && !strings.HasPrefix(string(handle), "@") {
		return errors.New("Twitter handle must start with @")
	}

	person.twitterHandle = handle
	return nil
}

func (person *Person) TwitterHandle() TwitterHandle {
	return person.twitterHandle
}
