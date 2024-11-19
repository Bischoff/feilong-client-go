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
		fmt.Println("Please specify FCP id")
		return
	}
	fcpid := os.Args[1]

	result, err := client.GetVolumeFCPUsage(fcpid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("FCP usage:\n")
	fmt.Printf("  Userid: %s\n", result.Output[0])
	fmt.Printf("  Reserved: %d\n", int(result.Output[1].(float64)))
	fmt.Printf("  Connections: %d\n", int(result.Output[2].(float64)))
	fmt.Printf("  FCP Template: %s\n", result.Output[3])
	fmt.Printf("\n")
}
