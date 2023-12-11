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
		fmt.Println("Please specify guest userid")
		return
	}
	userid := os.Args[1]

	result, err := client.GetGuestAdaptersInfo(userid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, s := range result.Output.Adapters {
		fmt.Printf("LAN owner: %s\n", s.LANOwner)
		fmt.Printf("LAN name: %s\n", s.LANName)
		fmt.Printf("Adapter address: %s\n", s.AdapterAddress)
		fmt.Printf("Adapter status: %s\n", s.AdapterStatus)
		fmt.Printf("MAC address: %s\n", s.MACAddress)
		fmt.Printf("IP address: %s\n", s.IPAddress)
		fmt.Printf("IP version: %s\n", s.IPVersion)
		fmt.Printf("\n")
	}
}
