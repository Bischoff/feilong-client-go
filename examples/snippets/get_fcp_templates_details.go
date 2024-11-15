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

	result, err := client.GetFCPTemplatesDetails(nil, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("FCP templates:\n")
	for _, template := range result.FCPTemplates {
		fmt.Printf("  Id: %s\n", template.Id)
		fmt.Printf("  Name: %s\n", template.Name)
		fmt.Printf("  Description: %s\n", template.Description)
		if (template.IsDefault) {
			fmt.Printf("  Is default: true\n")
		} else {
			fmt.Printf("  Is default: false\n")
		}
		fmt.Printf("  Storage providers:\n")
		for _, name := range template.SPName {
			fmt.Printf("   Name: %s\n", name)
		}
	}
}
