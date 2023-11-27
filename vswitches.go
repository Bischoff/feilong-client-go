/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"encoding/json"
)


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#list-vswitches

type ListVSwitchesResult struct {
	OverallRC	int	`json:"overallRC"`
	ReturnCode	int	`json:"rc"`
	Reason		int	`json:"rs"`
	ErrorMsg	string	`json:"errmsg"`
	ModuleId	int	`json:"modID"`
	Output		[]string `json:"output"`
}

func (c *Client) ListVSwitches() (*ListVSwitchesResult, error) {
	var result ListVSwitchesResult

	body, err := c.doRequest("GET", "/vswitches", nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
