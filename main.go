package main

import "github.com/tevity/go-custom-data-types/organization"

func main() {
	p := organization.Person{ FirstName: "Kevin", LastName: "Gurton"}
	
	println(p.ID())
}
