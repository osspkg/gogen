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

func ExampleImportBlock() {
	file := gogen.Package("main").
		ImportBlock(
			gogen.Text("fmt"),
			gogen.ID("json").Text("encoding/json"),
		).
		TypeBlock(
			gogen.ID("Name").String(),
			gogen.ID("Count").Int(),
		).
		Join(
			gogen.Func().ID("main").Bracket().Block(
				gogen.Pkg("fmt").ID("Println").Call(gogen.Text("generated")),
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
	// import (
	// 	json "encoding/json"
	// 	"fmt"
	// )
	//
	// type (
	// 	Name  string
	// 	Count int
	// )
	//
	// func main() {
	// 	fmt.Println("generated")
	// }
}

func Example() {
	expression := gogen.ID("lookup").TypeArgs(gogen.String()).Call(
		gogen.ID("items").Index(gogen.Raw("0")),
	)

	var source bytes.Buffer
	if err := expression.Render(&source); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(source.String())

	// Output:
	// lookup[string](items[0])
}
