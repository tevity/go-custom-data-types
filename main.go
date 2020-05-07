package main

import "github.com/tevity/go-custom-data-types/organization"

func main() {
	p := organization.NewPerson("Kevin", "Gurton")
	println(p.ID())
}
