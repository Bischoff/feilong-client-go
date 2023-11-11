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

	fileParams := feilong.ExportFileParams { SourceFile: "/etc/os-release" }
	result, err := client.ExportFile(&fileParams)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println(string(result.Contents))
}