// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	err := entc.Generate("../schema", &gen.Config{
		Header: `
			// Code generated (@generated) by entc, DO NOT EDIT.
		`,
		IDType:  &field.TypeInfo{Type: field.TypeInt},
		Target:  "../gen/models",
		Package: "gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models",
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
