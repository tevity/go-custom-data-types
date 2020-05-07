package organization

type socialSecurityNumber string

func NewSocialSecurityNumber(ssn string) Identifiable {
	return socialSecurityNumber(ssn)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}
