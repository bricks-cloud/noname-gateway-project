package cue

import (
	"fmt"
	"os"

	"cuelang.org/go/cue/cuecontext"
)

func MarshalJSON(filepath string) ([]byte, error) {
	ctx := cuecontext.New()

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read cue file: %v", err)
	}

	v := ctx.CompileBytes(data)

	return v.MarshalJSON()
}
