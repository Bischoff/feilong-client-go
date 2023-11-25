/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"encoding/json"
)

// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#smapi-healthy

type SMAPIHealthySMAPI struct {
	TotalSuccess	int	`json:"totalSuccess"`
	TotalFail	int	`json:"totalFail"`
	LastSuccess	string	`json:"lastSuccess"`
	LastFail	string	`json:"lastFail"`
	ContinuousFail	int	`json:"continueousFail"`
	Healthy		bool	`json:"healthy"`
}

type SMAPIHealthyResult struct {
	SMAPI		SMAPIHealthySMAPI `json:"SMAPI"`
}

func (c *Client) SMAPIHealthy() (*SMAPIHealthyResult, error) {
	var result SMAPIHealthyResult

	body, err := c.doRequest("GET", "/smapi-healthy", nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
