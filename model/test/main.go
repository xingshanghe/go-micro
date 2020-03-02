package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/model"
	"github.com/micro/go-micro/v2/store"
)

// users, err := model.FromFile("./models/user.yaml")

// Users is the user model
var Users = model.NewModel(
	model.Name("user"),
	model.Store(store.DefaultStore),
	model.WithUUIDAttribute("uuid",
		model.AttributePrimaryKey(),
	),
	model.WithStringAttribute("username",
		model.AttributeValidations(
			model.ValidateUnique(),
			model.ValidateURLSafe(),
			model.ValidateRequired(),
			model.ValidateMinimumLength(3),
			model.ValidateMaximumLength(15),
		),
	),
	model.WithStringAttribute("first_name"),
	model.WithStringAttribute("last_name"),
	model.WithStringAttribute("email",
		model.AttributeValidations(
			model.ValidateUnique(),
			model.ValidateEmail(),
		),
	),
)

func main() {
	fmt.Println("It compiled")
}

// r, err := Users.Get("foobar") // Looks up with primary key
// r.Set("first_name", "bar")
// r.Get("first_name")
// r.Save()
