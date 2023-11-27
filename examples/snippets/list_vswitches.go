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

	result, err := client.ListVSwitches()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, s := range result.Output {
		fmt.Println(s)
	}
	Println()
}
