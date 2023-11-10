## How to Compile and Run the Examples

From the `examples` directory, create a symbolic link to the snippet you want to use:

```bash
$ ln -s snippets/get_version.go main.go
```

Regenerate the `go.sum` file:

```bash
$ go mod tidy
```

Compile the example, it will go to `$GOPATH/bin`:

```bash
$ go install
```

Run the example:

```bash
$ example
```
