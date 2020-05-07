package main

import "github.com/tevity/go-custom-data-types/organization"

func main() {
	var p organization.Identifiable = organization.Person{}
	println(p.ID())
}
