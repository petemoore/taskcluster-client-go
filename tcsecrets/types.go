// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcsecrets

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Message containing a Taskcluster Secret
	//
	// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret.json#
	Secret struct {

		// An expiration date for this secret.
		//
		// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret.json#/properties/expires
		Expires tcclient.Time `json:"expires"`

		// The secret value to be encrypted.
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret.json#/properties/secret
		Secret json.RawMessage `json:"secret"`
	}

	// Message containing a list of secret names
	//
	// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret-list.json#
	SecretsList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of provisioners.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called with `continuationToken` until you get a
		// result without a `continuationToken`.
		//
		// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret-list.json#/properties/continuationToken
		ContinuationToken string `json:"continuationToken,omitempty"`

		// Secret names
		//
		// Array items:
		// Secret name
		//
		// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret-list.json#/properties/secrets/items
		//
		// See https://community-tc.services.mozilla.com/schemas/secrets/v1/secret-list.json#/properties/secrets
		Secrets []string `json:"secrets"`
	}
)
