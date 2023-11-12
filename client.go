/**
  Copyright Contributors to the Feilong Project.

  SPDX-License-Identifier: Apache-2.0
**/

package feilong

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultConnector string = "localhost:35000"
const defaultTimeout time.Duration = 300 * time.Second

type Client struct {
	Host		string
	HTTPClient	*http.Client
}

func NewClient(connector *string, timeout *time.Duration) (*Client) {
	h := defaultConnector
	if connector != nil {
		h = *connector
	}

	t := defaultTimeout
	if timeout != nil {
		t = *timeout
	}

	c := Client{
		HTTPClient:	&http.Client{Timeout: t},
		Host:		h,
	}

	return &c
}

// For internal use
func (c *Client) doRequest(method string, path string, params []byte) ([]byte, error) {
	url := "http://" + c.Host + path
	reader := bytes.NewReader(params)
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
