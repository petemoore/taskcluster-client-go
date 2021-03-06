// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcgithubevents

import (
	"encoding/json"
)

type (
	// Message reporting that a GitHub pull request has occurred
	//
	// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#
	GitHubPullRequestMessage struct {

		// The GitHub `action` which triggered an event.
		//
		// Possible values:
		//   * "assigned"
		//   * "unassigned"
		//   * "labeled"
		//   * "unlabeled"
		//   * "opened"
		//   * "edited"
		//   * "closed"
		//   * "reopened"
		//   * "synchronize"
		//   * "review_requested"
		//   * "review_request_removed"
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/action
		Action string `json:"action"`

		// The raw body of github event (for version 1)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/body
		Body json.RawMessage `json:"body"`

		// The head ref of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/branch
		Branch string `json:"branch"`

		// Metadata describing the pull request (for version 0)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/details
		Details json.RawMessage `json:"details,omitempty"`

		// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
		//
		// Syntax:     ^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/eventId
		EventID string `json:"eventId"`

		// The installation which had an event.
		//
		// Mininum:    0
		// Maximum:    10000000000
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/installationId
		InstallationID int64 `json:"installationId"`

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/organization
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/repository
		Repository string `json:"repository"`

		// The type of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/tasks_for
		Tasks_For string `json:"tasks_for"`

		// Message version
		//
		// Possible values:
		//   * 1
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-pull-request-message.json#/properties/version
		Version float64 `json:"version"`
	}

	// Message reporting that a GitHub push has occurred
	//
	// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#
	GitHubPushMessage struct {

		// The raw body of github event (for version 1)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/body
		Body json.RawMessage `json:"body"`

		// The head ref of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/branch
		Branch string `json:"branch"`

		// Metadata describing the push (for version 0)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/details
		Details json.RawMessage `json:"details,omitempty"`

		// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
		//
		// Syntax:     ^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/eventId
		EventID string `json:"eventId"`

		// The installation which had an event.
		//
		// Min length: 0
		// Max length: 10000000000
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/installationId
		InstallationID int64 `json:"installationId"`

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/organization
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/repository
		Repository string `json:"repository"`

		// The type of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/tasks_for
		Tasks_For string `json:"tasks_for"`

		// Message version
		//
		// Possible values:
		//   * 1
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-push-message.json#/properties/version
		Version float64 `json:"version"`
	}

	// Message reporting that a GitHub release has occurred
	//
	// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#
	GitHubReleaseMessage struct {

		// The raw body of github event (for version 1)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/body
		Body json.RawMessage `json:"body"`

		// The head ref of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/branch
		Branch string `json:"branch"`

		// Metadata describing the release (for version 0)
		//
		// Additional properties allowed
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/details
		Details json.RawMessage `json:"details,omitempty"`

		// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
		//
		// Syntax:     ^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/eventId
		EventID string `json:"eventId"`

		// The installation which had an event.
		//
		// Mininum:    0
		// Maximum:    10000000000
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/installationId
		InstallationID int64 `json:"installationId"`

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/organization
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/repository
		Repository string `json:"repository"`

		// The type of the event (for version 1)
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/tasks_for
		Tasks_For string `json:"tasks_for"`

		// Message version
		//
		// Possible values:
		//   * 1
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/github-release-message.json#/properties/version
		Version float64 `json:"version"`
	}

	// Indicates that this service has created a new task group in response to a GitHub event.
	// This message is for internal use only and should not be relied on for other purposes.
	// Full specification on [GitHub docs](https://developer.github.com/v3/repos/statuses/#create-a-status)
	//
	// See https://community-tc.services.mozilla.com/schemas/github/v1/task-group-creation-requested.json#
	TaskGroupDefinedCreateStatus struct {

		// The GitHub `organization` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/task-group-creation-requested.json#/properties/organization
		Organization string `json:"organization"`

		// The GitHub `repository` which had an event.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/task-group-creation-requested.json#/properties/repository
		Repository string `json:"repository"`

		// The id of the taskGroup that had been created.
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/task-group-creation-requested.json#/properties/taskGroupId
		TaskGroupID string `json:"taskGroupId"`

		// Message version
		//
		// Possible values:
		//   * 1
		//
		// See https://community-tc.services.mozilla.com/schemas/github/v1/task-group-creation-requested.json#/properties/version
		Version float64 `json:"version"`
	}
)
