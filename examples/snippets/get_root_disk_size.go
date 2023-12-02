/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package main

import (
	"fmt"
	"os"

	"github.com/Bischoff/feilong-client-go"
)

func main() {
	connector := os.Getenv("ZVM_CONNECTOR")
	client := feilong.NewClient(&connector, nil)

	if len(os.Args) != 2 {
		fmt.Println("Please specify image name")
		return
	}
	name := os.Args[1]

	result, err := client.GetRootDiskSize(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Information for %s:\n", name)
	fmt.Printf("  Root disk size: %s\n", result.Output)
}
