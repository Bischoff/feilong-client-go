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
	result, err := client.GetFCPTemplatesDetails(nil, false, true, false)
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
		fmt.Printf("    Storage providers:\n")
		for _, name := range template.StorageProviders {
			fmt.Printf("     Name: %s\n", name)
		}
		fmt.Printf("    Min FCP paths count: %d\n", template.MinFCPPathsCount)
		if template.Raw != nil {
			fmt.Printf("    Raw:\n")
			fmt.Printf("           fcp_id template_id assigner_id connections reserved wwpn_npiv wwpn_phy chpid state owner tmpl_id\n")
			for num, raw_data := range template.Raw {
				fmt.Printf("      %s: %v\n", num, raw_data)
			}
		}
		if template.Statistics != nil {
			fmt.Printf("    Statistics:\n")
			for num, statistics := range template.Statistics {
				fmt.Printf("      %s:\n", num)
				fmt.Printf("        Total: %s\n", statistics.Total)
				fmt.Printf("        Total count: ")
				for name, count := range statistics.TotalCount {
					fmt.Printf("%s: %d  ", name, count)
				}
				fmt.Printf("\n        Single FCP: %s\n", statistics.SingleFCP)
				fmt.Printf("        Range FCP: %s\n", statistics.RangeFCP)
				fmt.Printf("        Available: %s\n", statistics.Available)
				fmt.Printf("        Available count: ")
				for name, count := range statistics.AvailableCount {
					fmt.Printf("%s: %d  ", name, count)
				}
				fmt.Printf("\n        Allocated: %s\n", statistics.Allocated)
				fmt.Printf("        Reserve only: %s\n", statistics.ReserveOnly)
				fmt.Printf("        Connection only: %s\n", statistics.ConnectionOnly)
				fmt.Printf("        Unallocated but active: ")
				for name, value := range statistics.UnallocatedButActive {
					fmt.Printf("%s: %s  ", name, value)
				}
				fmt.Printf("\n        Allocated but free: %s\n", statistics.AllocatedButFree)
				fmt.Printf("        Not found: %s\n", statistics.NotFound)
				fmt.Printf("        Offline: %s\n", statistics.Offline)
				fmt.Printf("        Channel identifiers: ")
				for name, value := range statistics.ChannelIds {
					fmt.Printf("%s: %s  ", name, value)
				}
				fmt.Printf("\n        Physical channel identifiers: ")
				for name, value := range statistics.PhysicalChannelIds {
					fmt.Printf("%s: %s  ", name, value)
				}
				fmt.Printf("\n")
			}
		}
		fmt.Printf("    Physical channel identifiers: ")
		for _, name := range template.PhysicalChannelIds {
			fmt.Printf("%s ", name)
		}
		fmt.Printf("\n")
		fmt.Printf("    CPC serial number: %s\n", template.CPCSerialNumber)
		fmt.Printf("    CPC name: %s\n", template.CPCName)
		fmt.Printf("    Logical partition: %s\n", template.LogicalPartition)
		fmt.Printf("    Hypervisor hostname: %s\n", template.HypervisorHostname)
	}
	fmt.Printf("\n")
}
