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

	var pool *string = nil
	if len(os.Args) == 2 {
		// e.g. "ECKD:VMPOOL"
		pool = &os.Args[1]
	}

	result, err := client.GetHostDiskPoolVolumeNames(pool)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Volumes: %s\n", result.Output.DiskPoolVolumes)
}
