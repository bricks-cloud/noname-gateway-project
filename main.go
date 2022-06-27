package main

import (
	"context"
	"fmt"

	"github.com/bricks-cloud/bricks/compiler"
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

// Constants to map with CUE bricks.type selectors
const ROUTE_TYPE = "route"
const SERVICE_TYPE = "service"

func main() {
	v, _ := compiler.Build(context.Background(), "./cue", nil)

	var mappings []T

	fields, _ := v.Fields()
	for _, e := range fields {
		var s = e.Selector.String()
		t, _ := v.Lookup(s + ".$bricks.type").String()
		mappings = append(mappings, T{Selector: s, Type: t})
	}

	for _, m := range mappings {
		if m.Type == ROUTE_TYPE {
			t := &Route{}
			v.Lookup(m.Selector).Decode(t)
			fmt.Println("We got a Route:", t)
		} else if m.Type == SERVICE_TYPE {
			t := &Service{}
			v.Lookup(m.Selector).Decode(t)
			fmt.Println("We got a Service:", t)
		} else {
			fmt.Println("Unkown type:", m)

		}
	}

}
