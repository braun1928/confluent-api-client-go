package confluentclient

import "net/http"

// Client -
type Client struct {
	HostURL        string
	HTTPClient     *http.Client
	Token          string
	AccountID      string
	OrganizationID int
	UserID         int
}

// SessionRequest -
type sessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SessionResponse -
type sessionResponse struct {
	Token string      `json:"token"`
	Error interface{} `json:"error"`
	User  User        `json:"user"`
}

// User -
type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	OrganizationID     int    `json:"organization_id"`
	Deactivated        bool   `json:"deactivated"`
	Verified           string `json:"verified"`
	Created            string `json:"created"`
	Modified           string `json:"modified"`
	PasswordChanged    string `json:"password_changed"`
	ServiceName        string `json:"service_name"`
	ServiceDescription string `json:"service_description"`
	ServiceAccount     bool   `json:"service_account"`
	Internal           bool   `json:"internal"`
	ResourceID         string `json:"resource_id"`
	Sso                struct {
		Enabled             bool   `json:"enabled"`
		Auth0ConnectionName string `json:"auth0_connection_name"`
		TenantID            string `json:"tenant_id"`
		MultiTenant         bool   `json:"multi_tenant"`
	} `json:"sso"`
	Preferences struct {
		GettingStartedStatus string `json:"gettingStartedStatus"`
	} `json:"preferences"`
}

// LogicalClusters -
type LogicalClusters struct {
	ID string `json:"id"`
}

// CreateKeyFields -
type CreateKeyFields struct {
	AccountID       string            `json:"account_id"`
	UserID          int               `json:"user_id"`
	Description     string            `json:"description"`
	LogicalClusters []LogicalClusters `json:"logical_clusters"`
}

// CreateKeyRequest -
type CreateKeyRequest struct {
	APIKey CreateKeyFields `json:"api_key"`
}

// UpdateKeyFields -
type UpdateKeyFields struct {
	AccountID   string `json:"account_id"`
	KeyID       int    `json:"id"`
	Description string `json:"description"`
}

// UpdateKeyRequest -
type UpdateKeyRequest struct {
	APIKey UpdateKeyFields `json:"api_key"`
}

// DeleteKeyFields -
type DeleteKeyFields struct {
	AccountID string `json:"account_id"`
	KeyID     int    `json:"id"`
}

// DeleteKeyRequest -
type DeleteKeyRequest struct {
	APIKey DeleteKeyFields `json:"api_key"`
}

// DeleteKeyResponse -
type DeleteKeyResponse struct {
	APIKey struct {
		Deactivated bool `json:"deactivated"`
	} `json:"api_key"`
}

// GetAPIKeyResponse -
type GetAPIKeyResponse struct {
	APIKeys  []APIKey    `json:"api_keys"`
	Error    interface{} `json:"error"`
	PageInfo interface{} `json:"page_info"`
}

// APIKey -
type APIKey struct {
	Key             string            `json:"key"`
	Description     string            `json:"description"`
	KeyID           int               `json:"id"`
	UserID          int               `json:"user_id"`
	Secret          string            `json:"secret"`
	LogicalClusters []LogicalClusters `json:"logical_clusters"`
}

// GetAPIKeyRequest -
type GetAPIKeyRequest struct {
	ClusterID string
	KeyID     int
}

// CreateKeyResponse -
type CreateKeyResponse struct {
	APIKey APIKey `json:"api_key"`
}

// GetServiceAccountResponse -
type GetServiceAccountResponse struct {
	Users    []User      `json:"users"`
	Error    interface{} `json:"error"`
	PageInfo interface{} `json:"page_info"`
}

// CreateServiceAccountRequest -
type CreateServiceAccountRequest struct {
	User serviceAccountRequest `json:"user"`
}

// CreateServiceAccountResponse -
type CreateServiceAccountResponse struct {
	User  User        `json:"user"`
	Error interface{} `json:"error"`
}

// UpdateServiceAccountRequest -
type UpdateServiceAccountRequest struct {
	User UpdateServiceAccountFields `json:"user"`
}

// UpdateServiceAccountFields -
type UpdateServiceAccountFields struct {
	ServiceAccountID int    `json:"id"`
	Description      string `json:"serviceDescription"`
}

// DeleteServiceAccountRequest -
type DeleteServiceAccountRequest struct {
	User DeleteServiceAccountFields `json:"user"`
}

// DeleteServiceAccountFields -
type DeleteServiceAccountFields struct {
	ServiceAccountID int `json:"id"`
}

// deleteServiceAccountResponse -
type deleteServiceAccountResponse struct {
	Error interface{} `json:"error"`
}

// NewCreateServiceAccountRequest -
func NewCreateServiceAccountRequest(name, description string) CreateServiceAccountRequest {
	return CreateServiceAccountRequest{
		User: serviceAccountRequest{
			ServiceAccount:     true,
			ServiceDescription: description,
			ServiceName:        name,
		},
	}
}

// NewDeleteServiceAccountRequest -
func NewDeleteServiceAccountRequest(id int) DeleteServiceAccountRequest {
	return DeleteServiceAccountRequest{
		User: DeleteServiceAccountFields{
			ServiceAccountID: id,
		},
	}
}

// NewUpdateServiceAccountRequest -
func NewUpdateServiceAccountRequest(id int, description string) UpdateServiceAccountRequest {
	return UpdateServiceAccountRequest{
		User: UpdateServiceAccountFields{
			ServiceAccountID: id,
			Description:      description,
		},
	}
}

type serviceAccountRequest struct {
	ServiceAccount     bool   `json:"serviceAccount"`
	ServiceDescription string `json:"serviceDescription"`
	ServiceName        string `json:"serviceName"`
}

type account struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	Deactivated    bool   `json:"deactivated"`
	Created        string `json:"created"`
	Modified       string `json:"modified"`
	Config         struct {
		MaxKafkaClusters int `json:"max_kafka_clusters"`
	} `json:"config"`
	Internal bool `json:"internal"`
}

// Me -
type Me struct {
	User         User         `json:"user"`
	Account      account      `json:"account"`
	Organization organization `json:"organization"`
	Error        interface{}  `json:"error"`
	Accounts     []account    `json:"accounts"`
}

type organization struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Deactivated      bool   `json:"deactivated"`
	StripeCustomerID string `json:"stripe_customer_id"`
	Created          string `json:"created"`
	Modified         string `json:"modified"`
	BillingEmail     string `json:"billing_email"`
	Plan             struct {
		TaxAddress struct {
			Street1 string `json:"street1"`
			Street2 string `json:"street2"`
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
			Zip     string `json:"zip"`
		} `json:"tax_address"`
		ProductLevel string      `json:"product_level"`
		TrialStart   interface{} `json:"trial_start"`
		TrialEnd     interface{} `json:"trial_end"`
		PlanStart    interface{} `json:"plan_start"`
		PlanEnd      interface{} `json:"plan_end"`
		Coupon       interface{} `json:"coupon"`
		Product      interface{} `json:"product"`
		Billing      struct {
			Method           string `json:"method"`
			Interval         string `json:"interval"`
			AccruedThisCycle string `json:"accrued_this_cycle"`
			StripeCustomerID string `json:"stripe_customer_id"`
			Email            string `json:"email"`
		} `json:"billing"`
		ReferralCode     string `json:"referral_code"`
		AcceptTos        bool   `json:"accept_tos"`
		AllowMultiTenant bool   `json:"allow_multi_tenant"`
	} `json:"plan"`
	Saml struct {
		Enabled     bool   `json:"enabled"`
		MetadataURL string `json:"metadata_url"`
	} `json:"saml"`
	Sso struct {
		Enabled             bool   `json:"enabled"`
		Auth0ConnectionName string `json:"auth0_connection_name"`
		TenantID            string `json:"tenant_id"`
		MultiTenant         bool   `json:"multi_tenant"`
	} `json:"sso"`
	Marketplace struct {
		Partner           string `json:"partner"`
		CustomerID        string `json:"customer_id"`
		CustomerState     string `json:"customer_state"`
		ConsoleIntegrated bool   `json:"console_integrated"`
	} `json:"marketplace"`
	ResourceID     string `json:"resource_id"`
	HasEntitlement bool   `json:"has_entitlement"`
	ShowBilling    bool   `json:"show_billing"`
	AuditLog       struct {
		ClusterID        string `json:"cluster_id"`
		AccountID        string `json:"account_id"`
		ServiceAccountID int    `json:"service_account_id"`
		TopicName        string `json:"topic_name"`
	} `json:"audit_log"`
	HasCommitment           bool   `json:"has_commitment"`
	MarketplaceSubscription string `json:"marketplace_subscription"`
}
