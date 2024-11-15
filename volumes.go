/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"strings"
	"encoding/json"
)


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#refresh-volume-bootmap-info

type RefreshVolumeBootmapInfoParams struct {
	FCPChannel	[]string	`json:"fcpchannel"`
	WWPN		[]string	`json:"wwpn"`
	LUN		string		`json:"lun"`
	TransportFiles	string		`json:"transportfiles,omitempty"`
	GuestNetworks	[]GuestNetwork	`json:"guest_networks,omitempty"`
}

func (c *Client) RefreshVolumeBootmapInfo(params *RefreshVolumeBootmapInfoParams) error {
	wrapper := refreshVolumeBootmapInfoWrapper { Info: *params }
	body, err := json.Marshal(&wrapper)
	if err != nil {
		return err
	}
	_, err = c.doRequest("PUT", "/volumes/volume_refresh_bootmap", body)

	return err
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#get-volume-connector

type GetVolumeConnectorParams struct {
	Reserve		bool		`json:"reserve"`
	FCPTemplateId	string		`json:"fcp_template_id,omitempty"`
	StorageProvider	string		`json:"storage_provider,omitempty"`
	PCHIdInfo	string		`json:"pchid_info,omitempy"`
}

type GetVolumeConnectorOutput struct {
	FCP		[]string	`json:"zvm_fcp"`
	WWPNs		[]string	`json:"wwpns"`
	Host		string		`json:"host"`
}

type GetVolumeConnectorResult struct {
	OverallRC	int		`json:"overallRC"`
	ReturnCode	int		`json:"rc"`
	Reason		int		`json:"rs"`
	ErrorMsg	string		`json:"errmsg"`
	ModuleId	int		`json:"modID"`
	Output		GetVolumeConnectorOutput `json:"output"`
}

func (c *Client) GetVolumeConnector(userid string, params *GetVolumeConnectorParams) (*GetVolumeConnectorResult, error) {
	var result GetVolumeConnectorResult

	wrapper := getVolumeConnectorWrapper { Info: *params }
	body, err := json.Marshal(&wrapper)
	if err != nil {
		return nil, err
	}

	body, err = c.doRequest("GET", "/volumes/conn/" + userid, body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#get-fcp-usage

type GetVolumeFCPUsageResult struct {
	OverallRC	int		`json:"overallRC"`
	ReturnCode	int		`json:"rc"`
	Reason		int		`json:"rs"`
	ErrorMsg	string		`json:"errmsg"`
	ModuleId	int		`json:"modID"`
	Output		[]interface{}	`json:"output"`
}

func (c *Client) GetVolumeFCPUsage(fcpid string) (*GetVolumeFCPUsageResult, error) {
	var result GetVolumeFCPUsageResult

	body, err := c.doRequest("GET", "/volumes/fcp/" + fcpid, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#set-fcp-usage

type SetVolumeFCPUsageParams struct {
	UserId		string		`json:"userid"`
	Reserved	bool		`json:"reserved"`
	Connections	int		`json:"connections"`
}

func (c *Client) SetVolumeFCPUsage(fcpid string, params *SetVolumeFCPUsageParams) error {
	wrapper := setVolumeFCPUsageWrapper { Info: *params }
	body, err := json.Marshal(&wrapper)
	if err != nil {
		return err
	}
	_, err = c.doRequest("PUT", "/volumes/fcp/" + fcpid, body)

	return err
}


// Undocumented - GetFCPTemplatesDetails()

// TODO: create variants of this method with raw information and/or statistics
//       see GetHostDiskPoolInfo and GetHostDiskPoolDetails as model

type FCPTemplate struct {
	Id		string		`json:"id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	IsDefault	bool		`json:"is_default"`
	SPName		[]string	`json:"sp_name"`
}

type GetFCPTemplatesDetailsResult struct {
	FCPTemplates []FCPTemplate `json:"fcp_templates,omitempty"`
}

func (c *Client) GetFCPTemplatesDetails(templateIdList []string, syncWithZVM bool) (*GetFCPTemplatesDetailsResult, error) {
	var result GetFCPTemplatesDetailsResult

	req := "/volumes/fcptemplates/detail?"
	if templateIdList != nil {
		req += "template_id_list=" + strings.Join(templateIdList, " ") + "&"
	}
	req += "raw=false&"
	req += "statistics=false&"
	if (syncWithZVM) {
		req += "sync_with_zvm=true"
	} else {
		req += "sync_with_zvm=false"
	}

	body, err := c.doRequest("GET", req, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


// For internal use

type refreshVolumeBootmapInfoWrapper struct {
	Info		RefreshVolumeBootmapInfoParams `json:"info"`
}

type getVolumeConnectorWrapper struct {
	Info		GetVolumeConnectorParams `json:"info"`
}

type setVolumeFCPUsageWrapper struct {
	Info		SetVolumeFCPUsageParams `json:"info"`
}
