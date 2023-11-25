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

	result, err := client.SMAPIHealthy()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Total success:       %d\n", result.SMAPI.TotalSuccess)
	fmt.Printf("Total fail:          %d\n\n", result.SMAPI.TotalFail)
	fmt.Printf("Last success:        %s\n", result.SMAPI.LastSuccess)
	fmt.Printf("Last fail:           %s\n\n", result.SMAPI.LastFail)
	fmt.Printf("Continuous fail:     %d\n", result.SMAPI.ContinuousFail)
	if result.SMAPI.Healthy {
		fmt.Printf("Healthy:             yes\n\n")
	} else {
		fmt.Printf("Healthy:             no\n\n")
	}
}
