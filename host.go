/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"encoding/json"
)

// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#get-host-info

type GetHostInfoCPUInfo struct {
	CECModel	string	`json:"cec_model"`
	Architecture	string	`json:"architecture"`
}

type GetHostInfoResult struct {
	OverallRC	int	`json:"overallRC"`
	ReturnCode	int	`json:"rc"`
	Reason		int	`json:"rs"`
	ErrorMsg	string	`json:"errmsg"`
	ModuleId	int	`json:"modID"`
	Output struct {
		ZVMHost		string	`json:"zvm_host"`
		ZCCUserID	string	`json:"zcc_userid"`
		IPLTime		string	`json:"ipl_time"`
		HypervisorHostname string `json:"hypervisor_hostname"`
		HypervisorType	string	`json:"hypervisor_type"`
		HypervisorVersion int	`json:"hypervisor_version"`
		DiskTotal	int	`json:"disk_total"`
		DiskUsed	int	`json:"disk_used"`
		DiskAvailable	int	`json:"disk_available"`
		VCPUs		int	`json:"vcpus"`
		VCPUsUsed	int	`json:"vcpus_used"`
		MemoryMB	float64	`json:"memory_mb"`
		MemoryMBUsed	float64	`json:"memory_mb_used"`
		CPUInfo		GetHostInfoCPUInfo `json:"cpu_info"`
	}
}

func (c *Client) GetHostInfo() (*GetHostInfoResult, error) {
	var result GetHostInfoResult

	body, err := c.doRequest("GET", "/host", nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
