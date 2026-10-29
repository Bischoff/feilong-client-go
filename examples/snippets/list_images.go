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

const lineFormat = "%-27s%-11s%-34s%-10s%-12s%-9s %s\n"

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

	result, err := client.ListImages(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf(
		lineFormat,
		"Image",
		"Distro",
		"Checksum",
		"Size",
		"Bytes",
		"Type",
		"Comments",
	)
	for _, s := range result.Output {
		fmt.Printf(
			lineFormat,
			s.ImageName,
			s.ImageOSDistro,
			s.Checksum,
			s.DiskSizeUnits,
			s.ImageSizeInBytes,
			s.Type,
			s.Comments,
		)
	}
	fmt.Printf("\n")
}
