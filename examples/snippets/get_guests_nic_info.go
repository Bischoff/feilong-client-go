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

	result, err := client.GetGuestsNICInfo(nil, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Information about NICs:\n\n")
	for _, entry := range result.Output {
		fmt.Printf("UserId: %s\n", entry.UserId)
		fmt.Printf("Interface: %s\n", entry.Interface)
		fmt.Printf("VSwitch: %s\n", entry.VSwitch)
		fmt.Printf("Port: %s\n", entry.Port)
		fmt.Printf("Comments: %s\n\n", entry.Comments)
	}
}
