// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcworkermanagerevents

type (
	// The message that is emitted when worker pools are created/changed/deleted.
	//
	// See https://community-tc.services.mozilla.com/schemas/worker-manager/v1/pulse-worker-pool-message.json#
	WorkerTypePulseMessage struct {

		// If this is defined, it was the provider that handled this worker pool in the
		// configuration before the current one. This will be used by providers to clean
		// up any resources created for this workerType when they are no longer responsible
		// for it.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		//
		// See https://community-tc.services.mozilla.com/schemas/worker-manager/v1/pulse-worker-pool-message.json#/properties/previousProviderId
		PreviousProviderID string `json:"previousProviderId,omitempty"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		//
		// See https://community-tc.services.mozilla.com/schemas/worker-manager/v1/worker-pool-full.json#/properties/providerId
		ProviderID string `json:"providerId"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		//
		// See https://community-tc.services.mozilla.com/schemas/worker-manager/v1/worker-pool-full.json#/properties/workerPoolId
		WorkerPoolID string `json:"workerPoolId"`
	}
)
