package main

import (
	"context"
	"fmt"

	"github.com/bricks-cloud/noname-gateway-project/compiler"
)

type Service struct {
	Name        string
	Description string
	Url         string
	Routes      []Route
}

type Route struct {
	Name        string
	Description string
	Paths       []string
	Id          int
}

type T struct {
	Selector string
	Type     string
}

// TODO: Move to constants file
const (
	SERVICE_TYPE = "service"
)

func main() {
	v, error := compiler.Build(context.Background(), "./cue", nil)

	if error != nil {
		panic(error)
	}

	schemas := genSchemas(v)
	fmt.Println("Schemas: ", schemas)
}

func genSchemas(v *compiler.Value) []Service {
	// TODO: Once we have more types, schemas implement common interface to make list generic
	var schemas = []Service{}
	fields, error := v.Fields()

	if error != nil {
		panic(error)
	}

	for _, field := range fields {
		var selector = field.Selector.String()
		schemaType, _ := v.Lookup(selector + ".$bricks.type").String()
		switch schemaType {
		case SERVICE_TYPE:
			service := &Service{}
			v.Lookup(selector).Decode(service)
			fmt.Println("We got a Service:", service)
			schemas = append(schemas, *service)
		default:
			fmt.Println("Unknown type: ", selector)
		}
	}

	return schemas
}
