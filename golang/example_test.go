package golang_test

import (
	"bytes"
	"fmt"

	gogen "go.osspkg.com/gogen/golang"
)

func ExampleField() {
	file := gogen.Package("main").Join(
		gogen.Type().ID("User").Struct().Block(
			gogen.Field("ID", gogen.Uint64(), "json", "id,omitempty", "db", "user_id"),
			gogen.Field("Name", gogen.String(), "json", "name"),
		),
	)

	var source bytes.Buffer
	if err := gogen.Render(&source, file); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(source.String())

	// Output:
	// package main
	//
	// type User struct {
	// 	ID   uint64 `json:"id,omitempty" db:"user_id"`
	// 	Name string `json:"name"`
	// }
}
