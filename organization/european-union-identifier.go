package organization

import "strconv"

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return europeanUnionIdentifier{
			id:      v.ID(),
			country: country,
		}
	default:
		panic("Using an invalid type to initialise EU Identifier")
	}

}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}
