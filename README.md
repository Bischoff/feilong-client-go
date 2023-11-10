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
                fmt.Errorf("%s", err)
                return
        }

        fmt.Printf("API version: %s\n", result.Output.APIVersion)
}
```


## Completeness

The library implements [Feilong API](https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#) version 1.0.

The following are not implemented yet:

 * documentation
 * acceptance tests
 * many API functions (see below).


### Implemented Functions

The numbers below refer to the section numbers in the Feilong documentation. "partial" means that only the most common attributes are supported.

 * 7.2 - Version
   * 7.2.1 - `GetVersion`
 * 7.5 - Guests
   * 7.5.2 - `CreateGuest` - partial
   * 7.5.16 - `DeleteGuest`
   * 7.5.39 - `DeployGuest`
 * 7.7 - Images
   * 7.7.1 - `ListImages`
 * 7.9 - Files
   * 7.9.2 - `ExportFile`


## License

Apache 2.0, See LICENSE file
