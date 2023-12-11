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

	result, err := client.GetHostInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("ZVM Host:            %s\n", result.Output.ZVMHost)
	fmt.Printf("ZCC User ID:         %s\n", result.Output.ZCCUserID)
	fmt.Printf("IPL time:            %s\n\n", result.Output.IPLTime)
	fmt.Printf("Hypervisor hostname: %s\n", result.Output.HypervisorHostname)
	fmt.Printf("Hypervisor type:     %s\n", result.Output.HypervisorType)
	fmt.Printf("Hypervisor version:  %d\n\n", result.Output.HypervisorVersion)
	fmt.Printf("Disk total:          %d\n", result.Output.DiskTotal)
	fmt.Printf("Disk used:           %d\n", result.Output.DiskUsed)
	fmt.Printf("Disk available:      %d\n\n", result.Output.DiskAvailable)
	fmt.Printf("VCPUs:               %d\n", result.Output.VCPUs)
	fmt.Printf("VCPUs used:          %d\n\n", result.Output.VCPUsUsed)
	fmt.Printf("Memory (MB):         %f\n", result.Output.MemoryMB)
	fmt.Printf("Memory used (MB):    %f\n\n", result.Output.MemoryMBUsed)
	fmt.Printf("CPU CEC model:       %s\n", result.Output.CPUInfo.CECModel)
	fmt.Printf("CPU architecture:    %s\n\n", result.Output.CPUInfo.Architecture)
}
