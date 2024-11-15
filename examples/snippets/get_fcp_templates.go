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

	// To select explicitely some FCP templates, use a list of identifiers as first argument
	// Example:
	//       []string {"45c1d1be-a437-11ef-be87-e1b40fa110c4", "4982739c-a457-11ef-be87-e1b40fa110c4"}
	result, err := client.GetFCPTemplates(nil, nil, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("FCP templates:\n")
	for _, template := range result.Output.FCPTemplates {
		fmt.Printf("\n  Id: %s\n", template.Id)
		fmt.Printf("    Name: %s\n", template.Name)
		fmt.Printf("    Description: %s\n", template.Description)
		if (template.HostDefault) {
			fmt.Printf("    Host default: true\n")
		} else {
			fmt.Printf("    Host default: false\n")
		}
		fmt.Printf("    Storage provider default:\n")
		for _, name := range template.StorageProviderDefault {
			fmt.Printf("     Name: %s\n", name)
		}
		fmt.Printf("    Min FCP paths count: %d\n", template.MinFCPPathsCount)
		fmt.Printf("    CPC serial number: %s\n", template.CPCSerialNumber)
		fmt.Printf("    CPC name: %s\n", template.CPCName)
		fmt.Printf("    Logical partition: %s\n", template.LogicalPartition)
		fmt.Printf("    Hypervisor hostname: %s\n", template.HypervisorHostname)
		fmt.Printf("    Physical channel identifiers: ")
		for _, name := range template.PhysicalChannelIds {
			fmt.Printf("%s ", name)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
