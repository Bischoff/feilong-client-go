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
	if connector == "" {
		fmt.Println("Please define variable $ZVM_CONNECTOR")
		return
	}
	client := feilong.NewClient(&connector, nil)

	adminToken := os.Getenv("ZVM_ADMIN_TOKEN")
	if adminToken != "" {
		err := client.CreateToken(adminToken)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if len(os.Args) != 2 {
		fmt.Println("Please specify image name")
		return
	}
	name := os.Args[1]

	result, err := client.GetImageRootDiskSize(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Information for %s:\n", name)
	fmt.Printf("  Root disk size: %s\n", result.Output)
}
