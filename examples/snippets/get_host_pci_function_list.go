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

	result, err := client.GetHostPCIFunctionList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("PCIF:                %s\n", result.Output.PCIF)
	fmt.Printf("Status:              %s\n", result.Output.Status)
	fmt.Printf("UserId:              %s\n", result.Output.UserId)
	fmt.Printf("State:               %s\n", result.Output.State)
	fmt.Printf("Type:                %s\n", result.Output.Type)
}
