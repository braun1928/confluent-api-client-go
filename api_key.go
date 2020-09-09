package confluentclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Might be a way to list Cloud API Keys
// https://confluent.cloud/api/api_keys?account_id=t385&cloud=true

// GetAPIKey -
func (c *Client) GetAPIKey(keyID int, clusterID *string) (*APIKey, error) {

	// Confluence API does not have a specific key query from ID, only of key name from
	// what I have found. List all the keys in the account or on the cluster and filter on them
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/api_keys", c.HostURL), nil)

	// Form query parameters
	q := req.URL.Query()
	q.Add("account_id", c.AccountID)
	if clusterID != nil {
		q.Add("cluster_id", *clusterID)
	}
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := GetAPIKeyResponse{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	// No pagination?
	for _, k := range resp.APIKeys {
		if k.KeyID == keyID {
			return &k, nil
		}
	}

	return nil, nil

}

// CreateAPIKey -
func (c *Client) CreateAPIKey(createRequest CreateKeyRequest) (*CreateKeyResponse, error) {

	// form request body
	rb, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/api_keys", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := CreateKeyResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// UpdateAPIKey -
func (c *Client) UpdateAPIKey(updateRequest UpdateKeyRequest) (*CreateKeyResponse, error) {

	// form request body
	rb, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/api_keys/%d", c.HostURL, updateRequest.APIKey.KeyID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := CreateKeyResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// DeleteAPIKey -
func (c *Client) DeleteAPIKey(deleteRequest DeleteKeyRequest) error {

	// form request body
	rb, err := json.Marshal(deleteRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/api_keys/%d", c.HostURL, deleteRequest.APIKey.KeyID), strings.NewReader(string(rb)))

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	resp := DeleteKeyResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if !resp.APIKey.Deactivated {
		return errors.New(string("Key response did not specify deactivate. Unable to verify delete."))
	}

	return nil

}
