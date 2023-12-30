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

	result, err := client.GetGuestsStats(userid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Guest statistics for %s:\n\n", userid)
	for _, entry := range result.Output {
		fmt.Printf("Guest CPUs: %d\n", entry.GuestCPUs)
		fmt.Printf("Used CPU time: %d µs\n", entry.UsedCPUTime)
		fmt.Printf("Elapsed CPU time: %d µs\n", entry.ElapsedCPUTime)
		fmt.Printf("Min CPU count: %d\n", entry.MinCPUCount)
		fmt.Printf("Max CPU limit: %d\n", entry.MaxCPULimit)
		fmt.Printf("Samples CPU in use: %d\n", entry.SamplesCPUInUse)
		fmt.Printf("Samples CPU delay: %d\n", entry.SamplesCPUDelay)
		fmt.Printf("Used memory: %d KB\n", entry.UsedMemoryKB)
		fmt.Printf("Max memory: %d KB\n", entry.MaxMemoryKB)
		fmt.Printf("Min memory: %d KB\n", entry.MinMemoryKB)
		fmt.Printf("Shared memory: %d KB\n", entry.SharedMemoryKB)
		fmt.Printf("Total memory: %d\n", entry.TotalMemory)
		fmt.Printf("Available memory: %d\n", entry.AvailableMemory)
		fmt.Printf("Free memory: %d\n", entry.FreeMemory)
	}
}
