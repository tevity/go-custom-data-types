package organization

type socialSecurityNumber string

func NewSocialSecurityNumber(ssn string) Citizen {
	return socialSecurityNumber(ssn)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}
