/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"encoding/json"
)


// https://cloudlib4zvm.readthedocs.io/en/latest/restapi.html#list-images

type ListImagesImage struct {
	ImageName		string	`json:"imagename"`
	ImageOSDistro		string	`json:"imageosdistro"`
	MD5Sum			string	`json:"md5sum"`
	DiskSizeUnits		string	`json:"disk_size_units"`
	ImageSizeInBytes	string	`json:"image_size_in_bytes"`
	Type			string	`json:"type"`
	Comments		string	`json:"comments"`
	LastAccessTime		float64	`json:"last_access_time"`
}

type ListImagesResult struct {
	OverallRC	int	`json:"overallRC"`
	ReturnCode	int	`json:"rc"`
	Reason		int	`json:"rs"`
	ErrorMsg	string	`json:"errmsg"`
	ModuleId	int	`json:"modID"`
	Output		[]ListImagesImage `json:"output"`
}

func (c *Client) ListImages(imageName *string) (*ListImagesResult, error) {
	var result ListImagesResult

	path := "/images"
	if imageName != nil {
		path = path + "/" + *imageName
	}
	body, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
