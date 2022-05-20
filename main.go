package main

import (
	"fmt"

	"cuelang.org/go/cue/cuecontext"
)

func main() {
	ctx := cuecontext.New()

	v := ctx.CompileString(`
		a: 2
		b: 3
		"a+b": a + b
	`)
	js, err := v.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
