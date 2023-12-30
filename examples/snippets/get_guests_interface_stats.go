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

	result, err := client.GetGuestsInterfaceStats(userid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Interface statistics for %s:\n\n", userid)
	for _, entry := range result.Output[userid] {
		fmt.Printf("VSwitch: %s\n", entry.VSwitch)
		fmt.Printf("VDev: %s\n", entry.VDev)
		fmt.Printf("Frames received: %d\n", entry.FramesRec)
		fmt.Printf("Frames sent: %d\n", entry.FramesSent)
		fmt.Printf("Frames received discarded: %d\n", entry.FramesRecDisc)
		fmt.Printf("Frames sent discarded: %d\n", entry.FramesSentDisc)
		fmt.Printf("Frames received error: %d\n", entry.FramesRecErr)
		fmt.Printf("Frames sent error: %d\n", entry.FramesSentErr)
		fmt.Printf("Bytes received: %d\n", entry.BytesRec)
		fmt.Printf("Bytes sent: %d\n\n", entry.BytesSent)
	}
}
