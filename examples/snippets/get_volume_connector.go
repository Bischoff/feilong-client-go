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

	volumeConnectorParams := feilong.GetVolumeConnectorParams { Reserve: false }
	result, err := client.GetVolumeConnector(userid, &volumeConnectorParams)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Connection info for %s:\n", userid)
	for _, fcp := range result.Output.FCP {
		fmt.Printf("  FCP: %s\n", fcp)
	}
	for _, wwpn := range result.Output.WWPNs {
		fmt.Printf("  WWPN: %s\n", wwpn)
	}
	fmt.Printf("  Host: %s\n", result.Output.Host)
}
