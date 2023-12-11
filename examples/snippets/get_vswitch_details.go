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
		fmt.Println("Please specify vswitch name")
		return
	}
	name := os.Args[1]

	result, err := client.GetVSwitchDetails(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Information for %s:\n", name)
	fmt.Printf("  Switch name: %s\n", result.Output.SwitchName)
	fmt.Printf("  Switch type: %s\n", result.Output.SwitchType)
	fmt.Printf("  Switch status: %s\n", result.Output.SwitchStatus)
	fmt.Printf("  Port type: %s\n", result.Output.PortType)
	fmt.Printf("  Transport type: %s\n", result.Output.TransportType)
	fmt.Printf("  Isolation status: %s\n", result.Output.IsolationStatus)
	fmt.Printf("  User port based: %s\n", result.Output.UserPortBased)
	fmt.Printf("  VLAN awareness: %s\n", result.Output.VLANAwareness)
	fmt.Printf("  VLAN id: %s\n", result.Output.VLANId)
	fmt.Printf("  Native VLAN id: %s\n", result.Output.NativeVLANId)
	fmt.Printf("  Queue memory limit: %s\n", result.Output.QueueMemoryLimit)
	fmt.Printf("  MAC address: %s\n", result.Output.MACAddress)
	fmt.Printf("  MAC protect: %s\n", result.Output.MACProtect)
	fmt.Printf("  GVRP request attribute: %s\n", result.Output.GVRPRequestAttribute)
	fmt.Printf("  GVRP enabled attribute: %s\n", result.Output.GVRPEnabledAttribute)
	fmt.Printf("  Link ag: %s\n", result.Output.LinkAg)
	fmt.Printf("  Link ag group: %s\n", result.Output.LAgGroup)
	fmt.Printf("  Link ag interval: %s\n", result.Output.LAgInterval)
	fmt.Printf("  IP timeout: %s\n", result.Output.IPTimeout)
	fmt.Printf("  Routing value: %s\n", result.Output.RoutingValue)
	fmt.Printf("  VLAN counters: %s\n", result.Output.VLANCounters)
	fmt.Printf("  Real devices:     %s %11s %10s %19s  %s\n", "VDev", "controller", "port name", "dev status", "dev error")
	for name, details := range result.Output.RealDevices {
		fmt.Printf("    %14s: %s %11s %10s %19s  %s\n", name, details.VDev, details.Controller, details.PortName, details.DevStatus, details.DevErr)
	}
	fmt.Printf("  Authorized users: %9s %10s %10s %11s  %s\n", "port num", "prom mode", "OSD sim", "VLAN count", "VLAN ids")
	for name, details := range result.Output.AuthorizedUsers {
		fmt.Printf("    %14s: %9s %10s %10s %11s  %s\n", name, details.PortNum, details.PromMode, details.OSDSim, details.VLANCount, details.VLANIds)
	}
	fmt.Printf("  Adapters:         %18s %6s\n", "MAC", "type")
	for name, details := range result.Output.Adapters {
		fmt.Printf("    %14s: %18s %6s\n", name, details.MAC, details.Type)
	}
}
