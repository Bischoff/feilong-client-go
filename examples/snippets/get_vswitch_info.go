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
		fmt.Println("Please specify switch port id")
		return
	}
	port := os.Args[1]

	getVSwitchInfoParams := feilong.GetVSwitchInfoParams { PortId: port }
	result, err := client.GetVSwitchInfo(&getVSwitchInfoParams)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, nic := range result.Output {
		fmt.Printf("Userid: %s\n", nic.UserId)
		fmt.Printf("Interface: %s\n", nic.Interface)
		fmt.Printf("Switch: %s\n", nic.Switch)
		fmt.Printf("Port: %s\n", nic.Port)
		fmt.Printf("Comments: %s\n\n", nic.Comments)
	}
}
