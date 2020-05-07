package main

import (
	"fmt"

	"github.com/tevity/go-custom-data-types/organization"
)

func main() {
	p := organization.NewPerson("Kevin", "Gurton")
	println(p.FullName())

	err := p.SetTwitterHandle("@whatever")
	if err != nil {
		fmt.Printf("An error occurred setting Twitter handle: %s\n", err.Error())
	}

	println(p.TwitterHandle())
	println(p.TwitterHandle().RedirectUrl())
}
