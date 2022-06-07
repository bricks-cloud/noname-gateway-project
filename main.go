package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	example "github.com/bricks-cloud/bricks/commons"
	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
)

// func main() {
// 	v, err := compiler.Build(context.Background(), "./cue/bricks", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	controller := flow.New(&flow.Config{
// 		FindHiddenTasks: true,
// 	}, v.Cue(), ioTaskFunc)

// 	if err := controller.Run(context.Background()); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func ioTaskFunc(flowVal cue.Value) (flow.Runner, error) {
// 	v := compiler.Wrap(flowVal)

// 	typ := v.LookupPath(cue.MakePath(
// 		cue.Str("spotifyClient"),
// 		cue.Str("$bricks"),
// 		cue.Str("type"),
// 		cue.Hid("_resourceName", "github.com/bricks-cloud/bricks/cue/bricks")),
// 	)

// 	if typ.Exists() {
// 		fmt.Println(typ.String())
// 	}

// 	return flow.RunnerFunc(func(t *flow.Task) error {
// 		return nil
// 	}), nil
// }

func main() {

	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pluginFolder := "plugin"
	pluginBuildPath := filepath.Join(path, pluginFolder, "plugin_tmp")
	excutable := "greeter"
	excutablePath := filepath.Join(pluginBuildPath, excutable)

	if err = os.Mkdir(pluginBuildPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "build", "-o", excutablePath)
	cmd.Dir = filepath.Join(path, pluginFolder)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(excutablePath),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(example.Greeter)
	fmt.Println(greeter.Greet())

	err = os.RemoveAll(pluginBuildPath)
	if err != nil {
		log.Fatal(err)
	}
}

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
