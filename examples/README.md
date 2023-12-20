## Feilong Go Client Library Examples


### How to Compile the Examples

From the `examples` directory, create a symbolic link to the snippet you want to use:

```bash
$ ln -s snippets/get_version.go main.go
```

Generate the `go.sum` file:

```bash
$ go mod tidy
```

Compile the example, it will go to `$GOPATH/bin`:

```bash
$ go install
```


### How to Run the Examples

To run the example, you need to export the URL of the Feilong connector:

```bash
$ export ZVM_CONNECTOR="http://1.2.3.4"
$ example
```

Some snippets need a parameter on the command line. For example, if you compiled `get_guest_info.go`:

```bash
$ export ZVM_CONNECTOR="http://1.2.3.4"
$ example MYGUEST
```

If you implemented token authentication on the server side, you need to export one more environment variable:

```bash
$ export ZVM_CONNECTOR="http://1.2.3.4"
$ export ZVM_ADMIN_TOKEN="zvX2mFxuj8HcrYkAacLReV0RTQ0K5IIEighOR9F8AG"
$ example
```

If you mandated SSL on the server side, you might need to add your Certification Authority to the client. On a SUSE system:

```bash
$ sudo cp feilongCA.crt /etc/pki/trust/anchors/
$ sudo update-ca-certificates

$ export ZVM_CONNECTOR="https://1.2.3.4"
$ example
```
