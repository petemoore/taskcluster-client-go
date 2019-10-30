// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tclogin

import (
	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// A response containing credentials corresponding to a supplied OIDC `access_token`.
	//
	// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#
	CredentialsResponse struct {

		// Taskcluster credentials. Note that the credentials may not contain a certificate!
		//
		// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/credentials
		Credentials TaskclusterCredentials `json:"credentials"`

		// Time after which the credentials are no longer valid.  Callers should
		// call `oidcCredentials` again to get fresh credentials before this time.
		//
		// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/expires
		Expires tcclient.Time `json:"expires"`
	}

	// Taskcluster credentials. Note that the credentials may not contain a certificate!
	//
	// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/credentials
	TaskclusterCredentials struct {

		// Syntax:     ^[a-zA-Z0-9_-]{22,66}$
		//
		// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/credentials/properties/accessToken
		AccessToken string `json:"accessToken"`

		// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/credentials/properties/certificate
		Certificate string `json:"certificate,omitempty"`

		// Syntax:     ^[A-Za-z0-9!@/:.+|_-]+$
		//
		// See https://community-tc.services.mozilla.com/schemas/login/v1/oidc-credentials-response.json#/properties/credentials/properties/clientId
		ClientID string `json:"clientId"`
	}
)
