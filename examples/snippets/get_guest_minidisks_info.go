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

	result, err := client.GetGuestMinidisksInfo(userid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, s := range result.Output.Minidisks {
		fmt.Printf("VDev: %s\n", s.VDev)
		fmt.Printf("RDev: %s\n", s.RDev)
		fmt.Printf("Access type: %s\n", s.AccessType)
		fmt.Printf("Device type: %s\n", s.DeviceType)
		fmt.Printf("Device size: %d %s\n", s.DeviceSize, s.DeviceUnits)
		fmt.Printf("Volume label: %s\n\n", s.VolumeLabel)
	}
}
