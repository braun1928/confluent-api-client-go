package confluentclient_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	client "github.com/watkinsmike/terraform-provider/confluent-api-go-client"
)

func TestGetServiceAccount(t *testing.T) {

	const cRespData = `{"user": {"id": 123, "service_name": "service_account_A", "service_description": "special account", "service_account": true}}`
	newReq := client.NewCreateServiceAccountRequest("service_account_A", "special account")
	cReq := client.CreateServiceAccountRequest{}
	cResp := client.CreateServiceAccountResponse{}

	// form request body
	rb, err := json.Marshal(newReq)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(cRespData))

		err = json.Unmarshal(rb, &cReq)
		if err != nil {
			t.Fatal(err)
		}

		if cReq.User.ServiceAccount != true {
			t.Fatalf("expect service account %t. Got %t", true, cReq.User.ServiceAccount)
		}

		if cReq.User.ServiceName != "service_account_A" {
			t.Fatalf("expect service user name %s. Got %s", "service_account_A", cReq.User.ServiceName)
		}

		if cReq.User.ServiceDescription != "special account" {
			t.Fatalf("expect service user description %s. Got %s", "special account", cReq.User.ServiceDescription)

		}

	}))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader(string(rb)))
	c := &http.Client{}
	resp, err := c.Do(req)

	err = json.NewDecoder(resp.Body).Decode(&cResp)

	if cReq.User.ServiceAccount != cResp.User.ServiceAccount {
		t.Fatalf("expect service user %t. Got %t", cReq.User.ServiceAccount, cResp.User.ServiceAccount)
	}

	if cReq.User.ServiceName != cResp.User.ServiceName {
		t.Fatalf("expect service user name %s. Got %s", cReq.User.ServiceName, cResp.User.ServiceName)
	}

	if cReq.User.ServiceDescription != cResp.User.ServiceDescription {
		t.Fatalf("expect service user description %s. Got %s", cReq.User.ServiceDescription, cResp.User.ServiceDescription)

	}
}
