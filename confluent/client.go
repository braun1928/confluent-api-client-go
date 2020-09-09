package confluentclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const hostURL string = "https://confluent.cloud"

// NewClient -
func NewClient(host, email, password *string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		HostURL:    hostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (email != nil) && (password != nil) {

		// form request body
		rb, err := json.Marshal(sessionRequest{
			Email:    *email,
			Password: *password,
		})
		if err != nil {
			return nil, err
		}

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sessions", c.HostURL), strings.NewReader(string(rb)))
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(req)
		if err != nil {
			return nil, err
		}

		auth := sessionResponse{}
		err = json.Unmarshal(body, &auth)
		if err != nil {
			return nil, err
		}

		if auth.Token == "" {
			return nil, errors.New("Unable to properly authenticate. Failed to retrieve authentication token")
		}

		c.Token = auth.Token

		// Retrieve account information required for some requests
		me, err := c.Me()
		if err != nil {
			err := fmt.Errorf("Failed in retrieving account information: %s", err)
			return nil, err
		}

		c.AccountID = me.Account.ID
		c.UserID = me.User.ID
		c.OrganizationID = me.Organization.ID

	} else {
		err := errors.New("Either email or password was unset. Cannot authenticate")
		return nil, err
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("Cookie", fmt.Sprintf("auth_token=%s", c.Token))
	req.Header.Set("Content-Type", "application/json")

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
		return nil, fmt.Errorf("Confluent client - %s %s returned status: %d, body: %s", req.Method, req.URL, res.StatusCode, body)
	}

	return body, err
}
