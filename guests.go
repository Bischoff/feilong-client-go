/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"encoding/json"
)


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#create-guest

type CreateGuestDisk struct {
	Size		string	`json:"size"`
	IsBootDisk	bool	`json:"is_boot_disk"`
	Format		string	`json:"format,omitempty"`
	DiskPool	string	`json:"disk_pool,omitempty"`
}

type CreateGuestGuest struct {
	UserId		string	`json:"userid"`
	VCPUs		int	`json:"vcpus"`
	Memory		int	`json:"memory"`
	UserProfile	string	`json:"user_profile,omitempty"`
	DiskList	[]CreateGuestDisk `json:"disk_list"`
	MaxCPU		int	`json:"max_cpu,omitempty"`
	MaxMem		string	`json:"max_mem,omitempty"`
}

type CreateGuestParams struct {
	Guest		CreateGuestGuest `json:"guest"`
}

type CreateGuestResult struct {
	OverallRC	int	`json:"overallRC"`
	ReturnCode	int	`json:"rc"`
	Reason		int	`json:"rs"`
	ErrorMsg	string	`json:"errmsg"`
	ModuleId	int	`json:"modID"`
	Output		[]CreateGuestDisk `json:"output"`
}

func (c *Client) CreateGuest(params *CreateGuestParams) (*CreateGuestResult, error) {
	var result CreateGuestResult

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err = c.doRequest("POST", "/guests", body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#delete-guest

func (c *Client) DeleteGuest(userid string) (error) {
	_, err := c.doRequest("DELETE", "/guests/" + userid, nil)

	return err
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#start-guest

func (c *Client) StartGuest(userid string) (error) {
	params := simpleAction { Action: "start" }

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = c.doRequest("POST", "/guests/" + userid + "/action", body)

	return err
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#stop-guest

func (c *Client) StopGuest(userid string) (error) {
	params := simpleAction { Action: "stop" }

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = c.doRequest("POST", "/guests/" + userid + "/action", body)

	return err
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#deploy-guest

type DeployGuestParams struct {
	Action		string	`json:"action"`
	Image		string	`json:"image"`
	TransportFiles	string	`json:"transportfiles,omitempty"`
	RemoteHost	string	`json:"remotehost,omitempty"`
	VDev		string	`json:"vdev,omitempty"`
	Hostname	string	`json:"hostname,omitempty"`
	SkipDiskCopy	bool	`json:"skipdiskcopy,omitempty"`
}

func (c *Client) DeployGuest(userid string, params *DeployGuestParams) (error) {
	params.Action = "deploy"

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	_, err = c.doRequest("POST", "/guests/" + userid + "/action", body)

	return err
}

// For internal use

type simpleAction struct {
	Action		string	`json:"action"`
}
