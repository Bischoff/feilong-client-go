# Go Client Library for Feilong

This library enables to deploy s390 virtual machines on z/VM via [Feilong](https://openmainframeproject.org/projects/feilong/) from a Go program.


## Requirements

- [Go](https://golang.org/doc/install) >= 1.21


## Using the Library

Here is a small sample program using the library:

```go
package main

import (
        "fmt"
        "os"

        "github.com/Bischoff/feilong-client-go"
)

func main() {
        connector := os.Getenv("ZVM_CONNECTOR")

        client := feilong.NewClient(&connector, nil)

        result, err := client.GetFeilongVersion()
        if err != nil {
                fmt.Println(err.Error())
                return
        }

        fmt.Printf("API version: %s\n", result.Output.APIVersion)
}
```

For more examples, look in directory `examples/snippets`.


## Completeness

The library implements the [Feilong API](https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#) version 1.0.

The following are not implemented yet:

 * authentication
 * documentation
 * acceptance tests
 * many API functions (see below).


### Implemented Functions

The numbers below refer to the section numbers in the Feilong documentation.

 * 7.2 - Version
   * 7.2.1 - `GetVersion()`
 * 7.5 - Guests
   * 7.5.1 - `ListGuests()`
   * 7.5.2 - `CreateGuest()`
   * 7.5.15 - `ShowGuestDefinition()`
   * 7.5.16 - `DeleteGuest()`
   * 7.5.18 - `GetGuestInfo()`
   * 7.5.20 - `GetGuestAdaptersInfo()`
   * 7.5.21 - `CreateGuestNIC()`
   * 7.5.24 - `StartGuest()`
   * 7.5.25 - `StopGuest()`
   * 7.5.39 - `DeployGuest()`
   * 7.5.43 - `UpdateGuestNIC()`
 * 7.6 - Host
   * 7.6.2 - `GetHostInfo()`
 * 7.7 - Images
   * 7.7.1 - `ListImages()`
 * 7.9 - Files
   * 7.9.1 - `ImportFile()`
   * 7.9.2 - `ExportFile()`


## License

Apache 2.0, See LICENSE file
