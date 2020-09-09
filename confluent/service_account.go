package confluentclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetServiceAccount -
func (c *Client) GetServiceAccount(serviceAccountID int) (*User, error) {

	// Confluent API does not have a specific service account query from what I have found.
	// Have to list all the service accounts and filter on them
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/service_accounts", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := GetServiceAccountResponse{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	for _, sa := range resp.Users {
		if sa.ID == serviceAccountID {
			return &sa, nil
		}
	}

	return nil, nil

}

// CreateServiceAccount -
func (c *Client) CreateServiceAccount(createRequest CreateServiceAccountRequest) (*CreateServiceAccountResponse, error) {

	// form request body
	rb, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/service_accounts", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := CreateServiceAccountResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// UpdateServiceAccount -
func (c *Client) UpdateServiceAccount(updateRequest UpdateServiceAccountRequest) (*CreateServiceAccountResponse, error) {

	// form request body
	rb, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/service_accounts", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := CreateServiceAccountResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// DeleteServiceAccount -
func (c *Client) DeleteServiceAccount(deleteRequest DeleteServiceAccountRequest) error {

	// form request body
	rb, err := json.Marshal(deleteRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/service_accounts", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	resp := deleteServiceAccountResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return errors.New("service account not properly deleted. Unable to verify delete")
	}

	return nil

}
