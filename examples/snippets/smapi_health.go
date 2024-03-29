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

	result, err := client.SMAPIHealth()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Total success:       %d\n", result.Output.TotalSuccess)
	fmt.Printf("Total fail:          %d\n", result.Output.TotalFail)
	fmt.Printf("Last success:        %s\n", result.Output.LastSuccess)
	fmt.Printf("Last fail:           %s\n", result.Output.LastFail)
	fmt.Printf("Continuous fail:     %d\n", result.Output.ContinuousFail)
	if result.Output.Healthy {
		fmt.Printf("Healthy:             yes\n\n")
	} else {
		fmt.Printf("Healthy:             no\n\n")
	}
}
