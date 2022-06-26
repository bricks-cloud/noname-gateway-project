package main

import (
	"context"
	"fmt"

	example "github.com/bricks-cloud/bricks/commons"
	"github.com/bricks-cloud/bricks/compiler"
	plugin "github.com/hashicorp/go-plugin"
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

func main() {
	v, _ := compiler.Build(context.Background(), "./cue/bricks", nil)

	var mappings []T

	fields, _ := v.Fields()
	for _, e := range fields {
		var s = e.Selector.String()
		t, _ := v.Lookup(s + ".$bricks.type").String()
		mappings = append(mappings, T{Selector: s, Type: t})
	}

	for _, m := range mappings {
		if m.Type == "route" {
			t := &Route{}
			v.Lookup(m.Selector).Decode(t)
			fmt.Println("We got a Route:", t)
		} else if m.Type == "service" {
			t := &Service{}
			v.Lookup(m.Selector).Decode(t)
			fmt.Println("We got a Service:", t)
		} else {
			fmt.Println("Unkown type:", m)

		}
	}

}

// func ioTaskFunc(flowVal cue.Value) (flow.Runner, error) {
// 	v := compiler.Wrap(flowVal)

// 	typ := v.LookupPath(cue.MakePath(
// 		cue.Str("spotifyClient"),
// 		cue.Str("$bricks"),
// 		cue.Str("type"),
// 		cue.Hid("_name", "github.com/bricks-cloud/bricks/cue/bricks")),
// 	)

// 	if typ.Exists() {
// 		fmt.Println(typ.String())
// 	}

// 	return flow.RunnerFunc(func(t *flow.Task) error {
// 		return nil
// 	}), nil
// }

// func main() {

// 	// Create an hclog.Logger
// 	logger := hclog.New(&hclog.LoggerOptions{
// 		Name:   "plugin",
// 		Output: os.Stdout,
// 		Level:  hclog.Debug,
// 	})

// 	path, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pluginFolder := "plugin"
// 	pluginBuildPath := filepath.Join(path, pluginFolder, "plugin_tmp")
// 	excutable := "greeter"
// 	excutablePath := filepath.Join(pluginBuildPath, excutable)

// 	if err = os.Mkdir(pluginBuildPath, os.ModePerm); err != nil {
// 		log.Fatal(err)
// 	}

// 	cmd := exec.Command("go", "build", "-o", excutablePath)
// 	cmd.Dir = filepath.Join(path, pluginFolder)
// 	if err := cmd.Run(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// We're a host! Start by launching the plugin process.
// 	client := plugin.NewClient(&plugin.ClientConfig{
// 		HandshakeConfig: handshakeConfig,
// 		Plugins:         pluginMap,
// 		Cmd:             exec.Command(excutablePath),
// 		Logger:          logger,
// 	})
// 	defer client.Kill()

// 	// Connect via RPC
// 	rpcClient, err := client.Client()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Request the plugin
// 	raw, err := rpcClient.Dispense("greeter")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// We should have a Greeter now! This feels like a normal interface
// 	// implementation but is in fact over an RPC connection.
// 	greeter := raw.(example.Greeter)
// 	fmt.Println(greeter.Greet())

// 	err = os.RemoveAll(pluginBuildPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter": &example.GreeterPlugin{},
}
