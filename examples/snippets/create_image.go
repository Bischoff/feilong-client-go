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

	if len(os.Args) != 3 {
		fmt.Println("Please specify image name and upload URL")
		return
	}
	name := os.Args[1]
	url := os.Args[2]

	imageMeta := feilong.CreateImageMeta {
		OSVersion: "sles15.5",
	}
	imageParams := feilong.CreateImageParams {
		ImageName: name,
		URL: url,
		ImageMeta: imageMeta,
	}

	err := client.CreateImage(&imageParams)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
