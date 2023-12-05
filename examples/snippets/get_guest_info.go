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
	client := feilong.NewClient(&connector, nil)

	if len(os.Args) != 2 {
		fmt.Println("Please specify guest userid")
		return
	}
	userid := os.Args[1]

	result, err := client.GetGuestInfo(userid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Information for %s:\n", userid)
	fmt.Printf("  Maximum memory: %d KiB\n", result.Output.MaxMemKB)
	fmt.Printf("  Number of CPUs: %d\n", result.Output.NumCPUs)
	fmt.Printf("  CPU time: %d µs\n", result.Output.CPUTimeMuSec)
	fmt.Printf("  Power state: %s\n", result.Output.PowerState)
	fmt.Printf("  Current memory: %d KiB\n", result.Output.MemKB)
	fmt.Printf("  Number of online CPUs: %d\n", result.Output.OnlineCPUNum)
	fmt.Printf("  Operating system and distribution: %s\n", result.Output.OSDistro)
	fmt.Printf("  Kernel information: %s\n", result.Output.KernelInfo)
}
