package confluentclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Me -
// Returns information about the caller and the account they are associated with
func (c *Client) Me() (*Me, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/me", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var resp Me
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil

}
