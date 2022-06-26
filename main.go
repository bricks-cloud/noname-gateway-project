package main

import (
	"fmt"

	example "github.com/bricks-cloud/bricks/commons"
	plugin "github.com/hashicorp/go-plugin"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

type Service struct {
	name        string `cue:"name"`
	description string `cue:"description"`
	url         string `cue:"url"`
	// Routes      []Route `json:"routes"`
}

type Route struct {
	name        string
	description string
	paths       []string
}

func main() {

	// We need a cue.Context, the New'd return is ready to use
	ctx := cuecontext.New()

	// The entrypoints are the same as the files you'd specify at the command line
	entrypoints := []string{"./cue/bricks/client.cue"}

	// Load Cue files into Cue build.Instances slice
	// the second arg is a configuration object, we'll see this later
	bis := load.Instances(entrypoints, nil)

	// Loop over the instances, checking for errors and printing
	for _, bi := range bis {
		// check for errors on the instance
		// these are typically parsing errors
		if bi.Err != nil {
			fmt.Println("Error during load:", bi.Err)
			continue
		}

		// Use cue.Context to turn build.Instance to cue.Instance
		value := ctx.BuildInstance(bi)
		if value.Err() != nil {
			fmt.Println("Error during build:", value.Err())
			continue
		}

		// print the error
		fmt.Println("root value:", value)

		myService := &Service{}
		_ = value.Decode(myService)
		fmt.Println("myService...", myService)

		// var myService *Service
		// myService = new(Service)
		// fmt.Println(value.Decode(&myService))
		// fmt.Println(myService)

		// Validate the value
		err := value.Validate()
		if err != nil {
			fmt.Println("Error during validate:", err)
			continue
		}
	}

	// Loads cue file into go
	// v, err := compiler.Build(context.Background(), "./cue/bricks", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// taken from cue docs...
	// 	const config = `
	// msg:   "Hello \(place)!"
	// place: string | *"world" // "world" is the default.
	// `
	// 	var r cue.Runtime

	// 	instance, _ := r.Compile("test", config)

	// 	str, _ := instance.Lookup("msg").String()

	// 	fmt.Println(str)

	// 	type ab struct{ A, B int }

	// 	var x ab

	// 	i, _ := r.Compile("test", `{A: 2, B: 4}`)
	// 	_ = i.Value().Decode(&x)
	// 	fmt.Println(x)

	// 	i, _ = r.Compile("test", `{B: "foo"}`)
	// 	_ = i.Value().Decode(&x)
	// 	fmt.Println(x)

	// way 1
	// var myService *Service
	// myService = new(Service)
	// fmt.Println(v.Decode(&myService))
	// fmt.Println(myService)

	// another way 2
	// var myService Service
	// _ = v.Decode(&myService)
	// fmt.Println(myService)

	// Use the loaded files and turn into go shemas ...
	// ...

	// controller := flow.New(&flow.Config{
	// 	FindHiddenTasks: true,
	// }, v.Cue(), ioTaskFunc)

	// if err := controller.Run(context.Background()); err != nil {
	// 	fmt.Println(err)
	// }
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
