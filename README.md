# Local test for terraform

```
 go run main.go --file ./cue/timetest.cue
```

You should be about to find terraform artifacts under `.workdir`.

# Build CLI

To build the CLI, run:
``` bash
go build -o bricks cli.go
```

then run the CLI:

``` bash
./bricks
```
