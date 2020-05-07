package organization

import "fmt"

type Name struct {
	first string
	last  string
}

func (name *Name) FullName() string {
	return fmt.Sprintf("%s %s", name.first, name.last)
}
