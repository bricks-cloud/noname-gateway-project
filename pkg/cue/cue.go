package cue

import (
	"path/filepath"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func MarshalJSON(file string) ([]byte, error) {
	ctx := cuecontext.New()
	s := load.Instances([]string{filepath.Base(file)}, &load.Config{Dir: filepath.Dir(file)})[0]
	v := ctx.BuildInstance(s)
	return v.MarshalJSON()
}
