// The following code is AUTO-GENERATED. Please DO NOT edit.
// To update this generated code, run the following command:
// in the client subdirectory:
//
// go generate && go fmt

package client

type (
	// Response to a request for an Shared-Access-Signature to access and Azure
	// Table Storage table.
	AzureSharedAccessSignatureResponse struct {
		// Date and time of when the Shared-Access-Signature expires.
		Expiry string
		// Shared-Access-Signature string. This is the querystring parameters to
		// be appened after `?` or `&` depending on whether or not a querystring is
		// already present in the URL.
		Sas string
	}

	// Credentials, scopes and expiration date for a client
	GetClientCredentialsResponse struct {
		// AccessToken used for authenticating requests
		AccessToken string
		// ClientId of the client scopes is requested about
		ClientId string
		// Date and time where the clients credentials are set to expire
		Expires string
		// List of scopes the client is authorized to access
		Scopes []string
	}

	// Scopes and expiration date for a client
	GetClientScopesResponse struct {
		// ClientId of the client scopes is requested about
		ClientId string
		// Date and time where the clients credentials are set to expire
		Expires string
		// List of scopes the client is authorized to access
		Scopes []string
	}

	// Credentials, scopes and expiration date for a client
	GetClientCredentialsResponse1 struct {
		// Description of what these credentials are used for in markdown.
		// Please write a few details here, including who is the owner, point of
		// contact. Why it is scoped as is, think of this as documentation.
		Description string
		// Date and time where the clients credentials are set to expire
		Expires string
		// Human readable name of this set of credentials, typical
		// component/server-name or IRC nickname of the user.
		Name string
		// List of scopes the client is authorized to access
		Scopes []string
	}

	// Get all detaisl about a client, useful for tools modifying a client
	GetClientResponse struct {
		// AccessToken used for authenticating requests
		AccessToken string
		// ClientId of the client scopes is requested about
		ClientId string
		// Description of what these credentials are used for in markdown.
		// Should include who is the owner, point of contact.
		// Why it is scoped as is, think of this as documentation.
		Description string
		// Date and time where the clients credentials are set to expire
		Expires string
		// Human readable name of this set of credentials, typical
		// component/server-name or IRC nickname of the user.
		Name string
		// List of scopes the client is authorized to access
		Scopes []string
	}

	// Get a list of all clients, including basic information, but not credentials.
	ListClientsResponse struct {
	}

	// Representation of an indexed task.
	IndexedTaskResponse struct {
		// Data that was reported with the task. This is an arbitrary JSON object.
		Data struct {
		}
		// Date at which this entry expires from the task index.
		Expires string
		// Namespace of the indexed task, used to find the indexed task in the index.
		Namespace string
		// If multiple tasks are indexed with the same `namespace` the task with the
		// highest `rank` will be stored and returned in later requests. If two tasks
		// has the same `rank` the latest task will be stored.
		Rank int
		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		TaskId string
	}

	// Representation of an a task to be indexed.
	InsertTaskRequest struct {
		// This is an arbitrary JSON object. Feel free to put whatever data you want
		// here, but do limit it, you'll get errors if you store more than 32KB.
		// So stay well, below that limit.
		Data struct {
		}
		// Date at which this entry expires from the task index.
		Expires string
		// If multiple tasks are indexed with the same `namespace` the task with the
		// highest `rank` will be stored and returned in later requests. If two tasks
		// has the same `rank` the latest task will be stored.
		Rank int
		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		TaskId string
	}

	// Request to list namespaces within a given namespace.
	ListNamespacesRequest struct {
		// A continuation token previously returned in a response to this list
		// request. This property is optional and should not be provided for first
		// requests.
		ContinuationToken string
		// Maximum number of results per page. If there are more results than this
		// a continuation token will be return.
		Limit int
	}

	// Response from a request to list namespaces within a given namespace.
	ListNamespacesResponse struct {
		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string
		// List of namespaces.
		Namespaces []struct {
			// Date at which this entry, and by implication all entries below it,
			// expires from the task index.
			Expires string
			// Name of namespace within it's parent namespace.
			Name interface{}
			// Fully qualified name of the namespace, you can use this to list
			// namespaces or tasks under this namespace.
			Namespace string
		}
	}

	// Request to list tasks within a given namespace.
	ListTasksRequest struct {
		// A continuation token previously returned in a response to this list
		// request. This property is optional and should not be provided for first
		// requests.
		ContinuationToken string
		// Maximum number of results per page. If there are more results than this
		// a continuation token will be return.
		Limit int
	}

	// Representation of an indexed task.
	ListTasksResponse struct {
		// A continuation token is returned if there are more results than listed
		// here. You can optionally provide the token in the request payload to
		// load the additional results.
		ContinuationToken string
		// List of tasks.
		Tasks []struct {
			// Data that was reported with the task. This is an arbitrary JSON
			// object.
			Data struct {
			}
			// Date at which this entry expires from the task index.
			Expires string
			// Namespace of the indexed task, used to find the indexed task in the
			// index.
			Namespace string
			// If multiple tasks are indexed with the same `namespace` the task
			// with the highest `rank` will be stored and returned in later
			// requests. If two tasks has the same `rank` the latest task will be
			// stored.
			Rank int
			// Unique task identifier, this is UUID encoded as
			// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
			// stripped of `=` padding.
			TaskId string
		}
	}

	// Message reporting a new artifact has been created for a given task.
	ArtifactCreatedMessage struct {
		// Information about the artifact that was created
		Artifact interface{}
		// Id of the run on which artifact was created.
		RunId  int
		Status interface{}
		// Message version
		Version interface{}
		// Identifier for the worker-group within which the run with the created
		// artifacted is running.
		WorkerGroup string
		// Identifier for the worker within which the run with the created artifact
		// is running.
		WorkerId string
	}

	// Request to claim work
	WorkClaimRequest struct {
		// Identifier for group that worker claiming the task is a part of.
		WorkerGroup string
		// Identifier for worker within the given workerGroup
		WorkerId string
	}

	// Definition of a task that can be scheduled
	TaskDefinition struct {
		// Creation time of task
		Created string
		// Deadline of the task, `pending` and `running` runs are resolved as **failed** if not resolved by other means before the deadline
		Deadline string
		// Object with properties that can hold any kind of extra data that should be
		// associated with the task. This can be data for the task which doesn't
		// fit into `payload`, or it can supplementary data for use in services
		// listening for events from this task. For example this could be details to
		// display on _treeherder_, or information for indexing the task. Please, try
		// to put all related information under one property, so `extra` data keys
		// for treeherder reporting and task indexing don't conflict, hence, we have
		// reusable services. **Warning**, do not stuff large data-sets in here,
		// task definitions should not take-up multiple MiBs.
		Extra struct {
		}
		// Required task metadata
		Metadata struct {
			// Human readable description of the task, please **explain** what the
			// task does. A few lines of documentation is not going to hurt you.
			Description string
			// Human readable name of task, used to very briefly given an idea about
			// what the task does.
			Name string
			// E-mail of person who caused this task, e.g. the person who did
			// `hg push`. The person we should contact to ask why this task is here.
			Owner string
			// Link to source of this task, should specify a file, revision and
			// repository. This should be place someone can go an do a git/hg blame
			// to who came up with recipe for this task.
			Source string
		}
		// Task-specific payload following worker-specific format. For example the
		// `docker-worker` requires keys like: `image`, `commands` and
		// `features`. Refer to the documentation of `docker-worker` for details.
		Payload struct {
		}
		// Unique identifier for a provisioner, that can supply specified
		// `workerType`
		ProvisionerId string
		// Number of times to retry the task in case of infrastructure issues.
		// An _infrastructure issue_ is a worker node that crashes or is shutdown,
		// these events are to be expected.
		Retries int
		// List of task specific routes, AMQP messages will be CC'ed to these routes.
		Routes []string
		// Identifier for the scheduler that _defined_ this task, this can be an
		// identifier for a user or a service like the `"task-graph-scheduler"`.
		// Along with the `taskGroupId` this is used to form the permission scope
		// `queue:assume:scheduler-id:<schedulerId>/<taskGroupId>`,
		// this scope is necessary to _schedule_ a defined task, or _rerun_ a task.
		SchedulerId string
		// List of scopes (or scope-patterns) that the task is
		// authorized to use.
		Scopes []string
		// Arbitrary key-value tags (only strings limited to 4k). These can be used
		// to attach informal meta-data to a task. Use this for informal tags that
		// tasks can be classified by. You can also think of strings here as
		// candidates for formal meta-data. Something like
		// `purpose: 'build' || 'test'` is a good example.
		Tags struct {
		}
		// Identifier for a group of tasks scheduled together with this task, by
		// scheduler identified by `schedulerId`. For tasks scheduled by the
		// task-graph scheduler, this is the `taskGraphId`.  Defaults to `taskId` if
		// property isn't specified.
		TaskGroupId string
		// Unique identifier for a worker-type within a specific provisioner
		WorkerType string
	}

	// List of artifacts for a given `taskId` and `runId`.
	ListArtifactsResponse struct {
		// List of artifacts for given `taskId` and `runId`.
		Artifacts []struct {
			// Mimetype for the artifact that was created.
			ContentType string
			// Date and time after which the artifact created will be automatically
			// deleted by the queue.
			Expires string
			// Name of the artifact that was created, this is useful if you want to
			// attempt to fetch the artifact.
			Name string
			// This is the `storageType` for the request that was used to create
			// the artifact.
			StorageType interface{}
		}
	}

	// Response to request for poll task urls.
	PollTaskUrlsResponse struct {
		// Date and time after which the signed URLs provided in this response
		// expires and not longer works for authentication.
		Expires string
		// List of signed URLs to poll tasks from, they must be called in the order
		// they are given. As the first entry in this array **may** have higher
		// priority.
		SignedPollTaskUrls []string
	}

	// Request a authorization to put and artifact or posting of a URL as an artifact. Note that the `storageType` property is referenced in the response as well.
	PostArtifactRequest struct {
	}

	// Response to a request for posting an artifact. Note that the `storageType` property is referenced in the request as well.
	PostArtifactResponse struct {
	}

	// Request to claim (or reclaim) a task
	TaskClaimRequest struct {
		// MessageId from Azure Queue message
		MessageId string
		// PopReceipt from Azure Queue message
		Receipt string
		// Opaque token from the JSON parsed and base64 decoded MessageText in the Azure Queue message
		Token string
		// Identifier for group that worker claiming the task is a part of.
		WorkerGroup string
		// Identifier for worker within the given workerGroup
		WorkerId string
	}

	// Response to a successful task claim
	TaskClaimResponse struct {
		// `run-id` assigned to this run of the task
		RunId  int
		Status interface{}
		// Time at which the run expires and is resolved as `failed`, if the run isn't reclaimed.
		TakenUntil string
		// Identifier for the worker-group within which this run started.
		WorkerGroup string
		// Identifier for the worker executing this run.
		WorkerId string
	}

	// Message reporting that a task has complete successfully.
	TaskCompletedMessage struct {
		// Id of the run that completed the task
		RunId  int
		Status interface{}
		// Message version
		Version interface{}
		// Identifier for the worker-group within which this run ran.
		WorkerGroup string
		// Identifier for the worker that executed this run.
		WorkerId string
	}

	// Request for a task to be declared completed
	TaskCompletedRequest struct {
		// True, if task is completed, and false if task is failed. This property
		// is optional and only present for backwards compatibility. It will be
		// removed in the future.
		Success bool
	}

	// Message reporting that a task has been defined. The task may or may not be
	// _scheduled_ too.
	TaskDefinedMessage struct {
		Status interface{}
		// Message version
		Version interface{}
	}

	// Message reporting that TaskCluster have failed to run a task.
	TaskExceptionMessage struct {
		// Id of the last run for the task, not provided if `deadline`
		// was exceeded before a run was started.
		RunId  int
		Status interface{}
		// Message version
		Version interface{}
		// Identifier for the worker-group within which the last attempt of the task
		// ran. Not provided, if `deadline` was exceeded before a run was started.
		WorkerGroup string
		// Identifier for the last worker that failed to report, causing the task
		// to fail. Not provided, if `deadline` was exceeded before a run
		// was started.
		WorkerId string
	}

	// Request for a run of a task to be resolved with an exception
	TaskExceptionRequest struct {
		// Reason that the task is resolved with an exception. This is a subset
		// of the values for `resolvedReason` given in the task status structure.
		// Please, report `worker-shutdown` if the run failed because the worker
		// had to shutdown (spot node disappearing).
		// And report `malformed-payload` if the `task.payload` doesn't match the
		// schema for the worker payload, or referenced dependencies doesn't exists.
		// In either case, you should still log the error to a log file under the
		// specific run.
		Reason interface{}
	}

	// Message reporting that a task failed to complete successfully.
	TaskFailedMessage struct {
		// Id of the run that failed.
		RunId  int
		Status interface{}
		// Message version
		Version interface{}
		// Identifier for the worker-group within which this run ran.
		WorkerGroup string
		// Identifier for the worker that executed this run.
		WorkerId string
	}

	// Message reporting that a task is now pending
	TaskPendingMessage struct {
		// Id of run that became pending, `run-id`s always starts from 0
		RunId  int
		Status interface{}
		// Message version
		Version interface{}
	}

	// Message reporting that a given run of a task have started
	TaskRunningMessage struct {
		// Id of the run that just started, always starts from 0
		RunId  int
		Status interface{}
		// Time at which the run expires and is resolved as `failed`, if the run
		// isn't reclaimed.
		TakenUntil string
		// Message version
		Version interface{}
		// Identifier for the worker-group within which this run started.
		WorkerGroup string
		// Identifier for the worker executing this run.
		WorkerId string
	}

	// Response to a task status request
	TaskStatusResponse struct {
		Status interface{}
	}

	// A representation of **task status** as known by the queue
	TaskStatusStructure struct {
		// Deadline of the task, `pending` and `running` runs are resolved as **failed** if not resolved by other means before the deadline
		Deadline string
		// Unique identifier for the provisioner that this task must be scheduled on
		ProvisionerId string
		// Number of retries left for the task in case of infrastructure issues
		RetriesLeft int
		// List of runs, ordered so that index `i` has `runId == i`
		Runs []struct {
			// Reason for the creation of this run,
			// **more reasons may be added in the future**."
			ReasonCreated interface{}
			// Reason that run was resolved, this is mainly
			// useful for runs resolved as `exception`.
			// Note, **more reasons may be added in the future**, also this
			// property is only available after the run is resolved.
			ReasonResolved interface{}
			// Date-time at which this run was resolved, ie. when the run changed
			// state from `running` to either `completed`, `failed` or `exception`.
			// This property is only present after the run as been resolved.
			Resolved string
			// Id of this task run, `run-id`s always starts from `0`
			RunId int
			// Date-time at which this run was scheduled, ie. when the run was
			// created in state `pending`.
			Scheduled string
			// Date-time at which this run was claimed, ie. when the run changed
			// state from `pending` to `running`. This property is only present
			// after the run has been claimed.
			Started string
			// State of this run
			State interface{}
			// Time at which the run expires and is resolved as `failed`, if the
			// run isn't reclaimed. Note, only present after the run has been
			// claimed.
			TakenUntil string
			// Identifier for group that worker who executes this run is a part of,
			// this identifier is mainly used for efficient routing.
			// Note, this property is only present after the run is claimed.
			WorkerGroup string
			// Identifier for worker evaluating this run within given
			// `workerGroup`. Note, this property is only available after the run
			// has been claimed.
			WorkerId string
		}
		// Identifier for the scheduler that _defined_ this task.
		SchedulerId string
		// State of this task. This is just an auxiliary property derived from state
		// of latests run, or `unscheduled` if none.
		State interface{}
		// Identifier for a group of tasks scheduled together with this task, by
		// scheduler identified by `schedulerId`. For tasks scheduled by the
		// task-graph scheduler, this is the `taskGraphId`.
		TaskGroupId string
		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		TaskId string
		// Identifier for worker type within the specified provisioner
		WorkerType string
	}

	// Definition of a task that can be scheduled
	TaskDefinition1 struct {
		// Creation time of task
		Created string
		// Deadline of the task, `pending` and `running` runs are resolved as **failed** if not resolved by other means before the deadline
		Deadline string
		// Object with properties that can hold any kind of extra data that should be
		// associated with the task. This can be data for the task which doesn't
		// fit into `payload`, or it can supplementary data for use in services
		// listening for events from this task. For example this could be details to
		// display on _treeherder_, or information for indexing the task. Please, try
		// to put all related information under one property, so `extra` data keys
		// for treeherder reporting and task indexing don't conflict, hence, we have
		// reusable services. **Warning**, do not stuff large data-sets in here,
		// task definitions should not take-up multiple MiBs.
		Extra struct {
		}
		// Required task metadata
		Metadata struct {
			// Human readable description of the task, please **explain** what the
			// task does. A few lines of documentation is not going to hurt you.
			Description string
			// Human readable name of task, used to very briefly given an idea about
			// what the task does.
			Name string
			// E-mail of person who caused this task, e.g. the person who did
			// `hg push`. The person we should contact to ask why this task is here.
			Owner string
			// Link to source of this task, should specify a file, revision and
			// repository. This should be place someone can go an do a git/hg blame
			// to who came up with recipe for this task.
			Source string
		}
		// Task-specific payload following worker-specific format. For example the
		// `docker-worker` requires keys like: `image`, `commands` and
		// `features`. Refer to the documentation of `docker-worker` for details.
		Payload struct {
		}
		// Unique identifier for a provisioner, that can supply specified
		// `workerType`
		ProvisionerId string
		// Number of times to retry the task in case of infrastructure issues.
		// An _infrastructure issue_ is a worker node that crashes or is shutdown,
		// these events are to be expected.
		Retries int
		// List of task specific routes, AMQP messages will be CC'ed to these routes.
		Routes []string
		// Identifier for the scheduler that _defined_ this task, this can be an
		// identifier for a user or a service like the `"task-graph-scheduler"`.
		// Along with the `taskGroupId` this is used to form the permission scope
		// `queue:assume:scheduler-id:<schedulerId>/<taskGroupId>`,
		// this scope is necessary to _schedule_ a defined task, or _rerun_ a task.
		SchedulerId string
		// List of scopes (or scope-patterns) that the task is
		// authorized to use.
		Scopes []string
		// Arbitrary key-value tags (only strings limited to 4k). These can be used
		// to attach informal meta-data to a task. Use this for informal tags that
		// tasks can be classified by. You can also think of strings here as
		// candidates for formal meta-data. Something like
		// `purpose: 'build' || 'test'` is a good example.
		Tags struct {
		}
		// Identifier for a group of tasks scheduled together with this task, by
		// scheduler identified by `schedulerId`. For tasks scheduled by the
		// task-graph scheduler, this is the `taskGraphId`.  Defaults to `taskId` if
		// property isn't specified.
		TaskGroupId string
		// Unique identifier for a worker-type within a specific provisioner
		WorkerType string
	}

	// Definition of a task-graph that can be scheduled
	TaskGraphDefinition struct {
		// List of nodes in the task-graph, each featuring a task definition and scheduling preferences, such as number of _reruns_ to attempt.
		Tasks []struct {
			// List of required `taskId`s
			Requires []string
			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			Reruns int
			Task   interface{}
			// Task identifier (`taskId`) for the task when submitted to the queue, also used in `requires` below. This must be formatted as a **slugid** that is a uuid encoded in url-safe base64 following [RFC 4648 sec. 5](http://tools.ietf.org/html/rfc4648#section-5)), but without `==` padding.
			TaskId string
		}
	}

	// Information about a **task-graph** as known by the scheduler, with all the state of all individual tasks.
	InspectTaskGraphResponse struct {
		// Required task metadata
		Metadata struct {
			// Human readable description of task-graph, **explain** what it does!
			Description string
			// Human readable name of task-graph
			Name string
			// E-mail of person who caused this task-graph, e.g. the person who did `hg push`
			Owner string
			// Link to source of this task-graph, should specify file, revision and repository
			Source string
		}
		Status interface{}
		// Arbitrary key-value tags (only strings limited to 4k)
		Tags struct {
		}
		// Mapping from task-labels to task information and state.
		Tasks []struct {
			// List of `taskId`s that requires this task to be _complete successfully_ before they can be scheduled.
			Dependents []string
			// Human readable name from the task definition
			Name string
			// List of required `taskId`s
			Requires []string
			// List of `taskId`s that have yet to complete successfully, before this task can be scheduled.
			RequiresLeft []string
			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			Reruns int
			// Number of reruns that haven't been used yet.
			RerunsLeft int
			// true, if the scheduler considers the task node as satisfied and hence no-longer prevents dependent tasks from running.
			Satisfied bool
			// State of the task as considered by the scheduler
			State interface{}
			// Unique task identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
			TaskId string
		}
	}

	// Information about a **task** in a task-graph as known by the scheduler.
	InspectTaskGraphTaskResponse struct {
		// List of `taskId`s that requires this task to be _complete successfully_ before they can be scheduled.
		Dependents []string
		// Human readable name from the task definition
		Name string
		// List of required `taskId`s
		Requires []string
		// List of `taskId`s that have yet to complete successfully, before this task can be scheduled.
		RequiresLeft []string
		// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
		Reruns int
		// Number of reruns that haven't been used yet.
		RerunsLeft int
		// true, if the scheduler considers the task node as satisfied and hence no-longer prevents dependent tasks from running.
		Satisfied bool
		// State of the task as considered by the scheduler
		State interface{}
		// Unique task identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
		TaskId string
	}

	// Message that all reruns of a task has failed it is now blocking the task-graph from finishing.
	BlockedTaskGraphMessage struct {
		Status interface{}
		// Unique `taskId` that is blocking this task-graph from completion.
		TaskId string
		// Message version
		Version interface{}
	}

	// Messages as posted to `scheduler/v1/task-graph-extended` informing the world that a task-graph have been extended.
	TaskGraphExtendedMessage struct {
		Status interface{}
		// Message version
		Version interface{}
	}

	// Message that all tasks in a task-graph have now completed successfully and the graph is _finished_.
	TaskGraphFinishedMessage struct {
		Status interface{}
		// Message version
		Version interface{}
	}

	// Response for a request for task-graph information
	TaskGraphInfoResponse struct {
		// Required task metadata
		Metadata struct {
			// Human readable description of task-graph, **explain** what it does!
			Description string
			// Human readable name of task-graph
			Name string
			// E-mail of person who caused this task-graph, e.g. the person who did `hg push`
			Owner string
			// Link to source of this task-graph, should specify file, revision and repository
			Source string
		}
		Status interface{}
		// Arbitrary key-value tags (only strings limited to 4k)
		Tags struct {
		}
	}

	// Messages as posted to `scheduler/v1/task-graph-running` informing the world that a new task-graph have been submitted.
	NewTaskGraphMessage struct {
		Status interface{}
		// Message version
		Version interface{}
	}

	// Response containing the status structure for a task-graph
	TaskGraphStatusResponse struct {
		Status interface{}
	}

	// A representation of **task-graph status** as known by the scheduler, without the state of all individual tasks.
	TaskGraphStatusStructure struct {
		// Unique identifier for task-graph scheduler managing the given task-graph
		SchedulerId string
		// Task-graph state, this enum is **frozen** new values will **not** be added.
		State interface{}
		// Unique task-graph identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
		TaskGraphId string
	}

	// Definition of a task-graph that can be scheduled
	TaskGraphDefinition1 struct {
		// Required task metadata"
		Metadata struct {
			// Human readable description of task-graph, **explain** what it does!
			Description string
			// Human readable name of task-graph, give people finding this an idea
			// what this graph is about.
			Name string
			// E-mail of person who caused this task-graph, e.g. the person who did
			// `hg push` or whatever triggered it.
			Owner string
			// Link to source of this task-graph, should specify file, revision and
			// repository
			Source string
		}
		// List of task-graph specific routes, AMQP messages will be CC'ed to these
		// routes prefixed by `'route.'`.
		Routes []string
		// List of scopes (or scope-patterns) that tasks of the task-graph is
		// authorized to use.
		Scopes []string
		// Arbitrary key-value tags (only strings limited to 4k)
		Tags struct {
		}
		// List of nodes in the task-graph, each featuring a task definition and scheduling preferences, such as number of _reruns_ to attempt.
		Tasks []struct {
			// List of required `taskId`s
			Requires []string
			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			Reruns int
			Task   interface{}
			// Task identifier (`taskId`) for the task when submitted to the queue, also used in `requires` below. This must be formatted as a **slugid** that is a uuid encoded in url-safe base64 following [RFC 4648 sec. 5](http://tools.ietf.org/html/rfc4648#section-5)), but without `==` padding.
			TaskId string
		}
	}
)
type AuthAPI Auth

// Returns the scopes the client is authorized to access and the date-time
// where the clients authorization is set to expire.
//
// This API end-point allows you inspect clients without getting access to
// credentials, as provide by the `getCredentials` request below.
func (a AuthAPI) Scopes(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Returns the clients `accessToken` as needed for verifying signatures.
// This API end-point also returns the list of scopes the client is
// authorized for and the date-time where the client authorization expires
//
// Remark, **if you don't need** the `accessToken` but only want to see what
// scopes a client is authorized for, you should use the `getScopes`
// function described above.
func (a AuthAPI) GetCredentials(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Returns all information about a given client. This end-point is mostly
// building tools to administrate clients. Do not use if you only want to
// authenticate a request, see `getCredentials` for this purpose.
func (a AuthAPI) Client(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Create client with given `clientId`, `name`, `expires`, `scopes` and
// `description`. The `accessToken` will always be generated server-side,
// and will be returned from this request.
//
// **Required scopes**, in addition the scopes listed
// above, the caller must also posses the all the scopes that is given to
// the client that is created.
func (a AuthAPI) CreateClient(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Modify client `name`, `expires`, `scopes` and
// `description`.
//
// **Required scopes**, in addition the scopes listed
// above, the caller must also posses the all the scopes that is given to
// the client that is updated.
func (a AuthAPI) ModifyClient(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Delete a client with given `clientId`.
func (a AuthAPI) RemoveClient(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Reset credentials for a client. This will generate a new `accessToken`.
// as always the `accessToken` will be generated server-side and returned.
func (a AuthAPI) ResetCredentials(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Return list with all clients
func (a AuthAPI) ListClients(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get an SAS string for use with a specific Azure Table Storage table.
// Note, this will create the table, if it doesn't already exists.
func (a AuthAPI) AzureTableSAS(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning** this api end-point is **not stable**.
func (a AuthAPI) Ping(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

type IndexAPI Auth

// Find task by namespace, if no task existing for the given namespace, this
// API end-point respond `404`.
func (a IndexAPI) FindTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// List the namespaces immediately under a given namespace. This end-point
// list up to 1000 namespaces. If more namespaces are present a
// `continuationToken` will be returned, which can be given in the next
// request. For the initial request, the payload should be an empty JSON
// object.
//
// **Remark**, this end-point is designed for humans browsing for tasks, not
// services, as that makes little sense.
func (a IndexAPI) ListNamespaces(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// List the tasks immediately under a given namespace. This end-point
// list up to 1000 tasks. If more tasks are present a
// `continuationToken` will be returned, which can be given in the next
// request. For the initial request, the payload should be an empty JSON
// object.
//
// **Remark**, this end-point is designed for humans browsing for tasks, not
// services, as that makes little sense.
func (a IndexAPI) ListTasks(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Insert a task into the index. Please see the introduction above, for how
// to index successfully completed tasks automatically, using custom routes.
func (a IndexAPI) InsertTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning** this api end-point is **not stable**.
func (a IndexAPI) Ping(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

type QueueAPI Auth

// Create a new task, this is an **idempotent** operation, so repeat it if
// you get an internal server error or network connection is dropped.
//
// **Task `deadline´**, the deadline property can be no more than 7 days
// into the future. This is to limit the amount of pending tasks not being
// taken care of. Ideally, you should use a much shorter deadline.
//
// **Task specific routing-keys**, using the `task.routes` property you may
// define task specific routing-keys. If a task has a task specific
// routing-key: `<route>`, then the poster will be required to posses the
// scope `queue:route:<route>`. And when the an AMQP message about the task
// is published the message will be CC'ed with the routing-key:
// `route.<route>`. This is useful if you want another component to listen
// for completed tasks you have posted.
func (a QueueAPI) CreateTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get task definition from queue.
func (a QueueAPI) GetTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Define a task without scheduling it. This API end-point allows you to
// upload a task definition without having scheduled. The task won't be
// reported as pending until it is scheduled, see the scheduleTask API
// end-point.
//
// The purpose of this API end-point is allow schedulers to upload task
// definitions without the tasks becoming _pending_ immediately. This useful
// if you have a set of dependent tasks. Then you can upload all the tasks
// and when the dependencies of a tasks have been resolved, you can schedule
// the task by calling `/task/:taskId/schedule`. This eliminates the need to
// store tasks somewhere else while waiting for dependencies to resolve.
//
// **Note** this operation is **idempotent**, as long as you upload the same
// task definition as previously defined this operation is safe to retry.
func (a QueueAPI) DefineTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// If you have define a task using `defineTask` API end-point, then you
// can schedule the task to be scheduled using this method.
// This will announce the task as pending and workers will be allowed, to
// claim it and resolved the task.
//
// **Note** this operation is **idempotent** and will not fail or complain
// if called with `taskId` that is already scheduled, or even resolved.
// To reschedule a task previously resolved, use `rerunTask`.
func (a QueueAPI) ScheduleTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get task status structure from `taskId`
func (a QueueAPI) Status(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get a signed url to get a message from azure queue.
// Once messages are polled from here, you can claim the referenced task
// with `claimTask`.
func (a QueueAPI) PollTaskUrls(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// claim a task, more to be added later...
//
// **Warning,** in the future this API end-point will require the presents
// of `receipt`, `messageId` and `token` in the body.
func (a QueueAPI) ClaimTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// reclaim a task more to be added later...
func (a QueueAPI) ReclaimTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Claim work for a worker, returns information about an appropriate task
// claimed for the worker. Similar to `claimTask`, which can be
// used to claim a specific task, or reclaim a specific task extending the
// `takenUntil` timeout for the run.
//
// **Note**, that if no tasks are _pending_ this method will not assign a
// task to you. Instead it will return `204` and you should wait a while
// before polling the queue again.
//
// **WARNING, this API end-point is deprecated and will be removed**.
func (a QueueAPI) ClaimWork(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Report a task completed, resolving the run as `completed`.
//
// For legacy, reasons the `success` parameter is accepted. This will be
// removed in the future.
func (a QueueAPI) ReportCompleted(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Report a run failed, resolving the run as `failed`. Use this to resolve
// a run that failed because the task specific code behaved unexpectedly.
// For example the task exited non-zero, or didn't produce expected output.
//
// Don't use this if the task couldn't be run because if malformed payload,
// or other unexpected condition. In these cases we have a task exception,
// which should be reported with `reportException`.
func (a QueueAPI) ReportFailed(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Resolve a run as _exception_. Generally, you will want to report tasks as
// failed instead of exception. But if the payload is malformed, or
// dependencies referenced does not exists you should also report exception.
// However, do not report exception if an external resources is unavailable
// because of network failure, etc. Only if you can validate that the
// resource does not exist.
func (a QueueAPI) ReportException(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// This method _reruns_ a previously resolved task, even if it was
// _completed_. This is useful if your task completes unsuccessfully, and
// you just want to run it from scratch again. This will also reset the
// number of `retries` allowed.
//
// Remember that `retries` in the task status counts the number of runs that
// the queue have started because the worker stopped responding, for example
// because a spot node died.
//
// **Remark** this operation is idempotent, if you try to rerun a task that
// isn't either `failed` or `completed`, this operation will just return the
// current task status.
func (a QueueAPI) RerunTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// This API end-point creates an artifact for a specific run of a task. This
// should **only** be used by a worker currently operating on this task, or
// from a process running within the task (ie. on the worker).
//
// All artifacts must specify when they `expires`, the queue will
// automatically take care of deleting artifacts past their
// expiration point. This features makes it feasible to upload large
// intermediate artifacts from data processing applications, as the
// artifacts can be set to expire a few days later.
//
// We currently support 4 different `storageType`s, each storage type have
// slightly different features and in some cases difference semantics.
//
// **S3 artifacts**, is useful for static files which will be stored on S3.
// When creating an S3 artifact is create the queue will return a pre-signed
// URL to which you can do a `PUT` request to upload your artifact. Note
// that `PUT` request **must** specify the `content-length` header and
// **must** give the `content-type` header the same value as in the request
// to `createArtifact`.
//
// **Azure artifacts**, are stored in _Azure Blob Storage_ service, which
// given the consistency guarantees and API interface offered by Azure is
// more suitable for artifacts that will be modified during the execution
// of the task. For example docker-worker has a feature that persists the
// task log to Azure Blob Storage every few seconds creating a somewhat
// live log. A request to create an Azure artifact will return a URL
// featuring a [Shared-Access-Signature](http://msdn.microsoft.com/en-us/library/azure/dn140256.aspx),
// refer to MSDN for further information on how to use these.
//
// **Reference artifacts**, only consists of meta-data which the queue will
// store for you. These artifacts really only have a `url` property and
// when the artifact is requested the client will be redirect the URL
// provided with a `303` (See Other) redirect. Please note that we cannot
// delete artifacts you upload to other service, we can only delete the
// reference to the artifact, when it expires.
//
// **Error artifacts**, only consists of meta-data which the queue will
// store for you. These artifacts are only meant to indicate that you the
// worker or the task failed to generate a specific artifact, that you
// would otherwise have uploaded. For example docker-worker will upload an
// error artifact, if the file it was supposed to upload doesn't exists or
// turns out to be a directory. Clients requesting an error artifact will
// get a `403` (Forbidden) response. This is mainly designed to ensure that
// dependent tasks can distinguish between artifacts that were suppose to
// be generated and artifacts for which the name is misspelled.
//
// **Artifact immutability**, generally speaking you cannot overwrite an
// artifact when created. But if you repeat the request with the same
// properties the request will succeed as the operation is idempotent.
// This is useful if you need to refresh a signed URL while uploading.
// Do not abuse this to overwrite artifacts created by another entity!
// Such as worker-host overwriting artifact created by worker-code.
//
// As a special case the `url` property on _reference artifacts_ can be
// updated. You should only use this to update the `url` property for
// reference artifacts your process has created.
func (a QueueAPI) CreateArtifact(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get artifact by `<name>` from a specific run.
//
// **Public Artifacts**, in-order to get an artifact you need the scope
// `queue:get-artifact:<name>`, where `<name>` is the name of the artifact.
// But if the artifact `name` starts with `public/`, authentication and
// authorization is not necessary to fetch the artifact.
//
// **API Clients**, this method will redirect you to the artifact, if it is
// stored externally. Either way, the response may not be JSON. So API
// client users might want to generate a signed URL for this end-point and
// use that URL with a normal HTTP client.
func (a QueueAPI) GetArtifact(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get artifact by `<name>` from the last run of a task.
//
// **Public Artifacts**, in-order to get an artifact you need the scope
// `queue:get-artifact:<name>`, where `<name>` is the name of the artifact.
// But if the artifact `name` starts with `public/`, authentication and
// authorization is not necessary to fetch the artifact.
//
// **API Clients**, this method will redirect you to the artifact, if it is
// stored externally. Either way, the response may not be JSON. So API
// client users might want to generate a signed URL for this end-point and
// use that URL with a normal HTTP client.
//
// **Remark**, this end-point is slightly slower than
// `queue.getArtifact`, so consider that if you already know the `runId` of
// the latest run. Otherwise, just us the most convenient API end-point.
func (a QueueAPI) GetLatestArtifact(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Returns a list of artifacts and associated meta-data for a given run.
func (a QueueAPI) ListArtifacts(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Returns a list of artifacts and associated meta-data for the latest run
// from the given task.
func (a QueueAPI) ListLatestArtifacts(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning** this api end-point is **not stable**.
//
// **This end-point is deprecated!**
func (a QueueAPI) GetPendingTasks(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning: This is an experimental end-point!**
func (a QueueAPI) PendingTaskCount(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
// This probably the end-point that will remain after rewriting to azure
// queue storage...
//
// **Warning: This is an experimental end-point!**
func (a QueueAPI) PendingTasks(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning** this api end-point is **not stable**.
func (a QueueAPI) Ping(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

type SchedulerAPI Auth

// Create a new task-graph, the `status` of the resulting JSON is a
// task-graph status structure, you can find the `taskGraphId` in this
// structure.
//
// **Referencing required tasks**, it is possible to reference other tasks
// in the task-graph that must be completed successfully before a task is
// scheduled. You just specify the `taskId` in the list of `required` tasks.
// See the example below, where the second task requires the first task.
// ```js
// {
//   ...
//   tasks: [
//     {
//       taskId:     "XgvL0qtSR92cIWpcwdGKCA",
//       requires:   [],
//       ...
//     },
//     {
//       taskId:     "73GsfK62QNKAk2Hg1EEZTQ",
//       requires:   ["XgvL0qtSR92cIWpcwdGKCA"],
//       task: {
//         payload: {
//           env: {
//             DEPENDS_ON:  "XgvL0qtSR92cIWpcwdGKCA"
//           }
//           ...
//         }
//         ...
//       },
//       ...
//     }
//   ]
// }
// ```
//
// **The `schedulerId` property**, defaults to the `schedulerId` of this
// scheduler in production that is `"task-graph-scheduler"`. This
// property must be either undefined or set to `"task-graph-scheduler"`,
// otherwise the task-graph will be rejected.
//
// **The `taskGroupId` property**, defaults to the `taskGraphId` of the
// task-graph submitted, and if provided much be the `taskGraphId` of
// the task-graph. Otherwise the task-graph will be rejected.
//
// **Task-graph scopes**, a task-graph is assigned a set of scopes, just
// like tasks. Tasks within a task-graph cannot have scopes beyond those
// the task-graph has. The task-graph scheduler will execute all requests
// on behalf of a task-graph using the set of scopes assigned to the
// task-graph. Thus, if you are submitting tasks to `my-worker-type` under
// `my-provisioner` it's important that your task-graph has the scope
// required to define tasks for this `provisionerId` and `workerType`.
// See the queue for details on permissions required. Note, the task-graph
// does not require permissions to schedule the tasks. This is done with
// scopes provided by the task-graph scheduler.
//
// **Task-graph specific routing-keys**, using the `taskGraph.routes`
// property you may define task-graph specific routing-keys. If a task-graph
// has a task-graph specific routing-key: `<route>`, then the poster will
// be required to posses the scope `scheduler:route:<route>`. And when the
// an AMQP message about the task-graph is published the message will be
// CC'ed with the routing-key: `route.<route>`. This is useful if you want
// another component to listen for completed tasks you have posted.
func (a SchedulerAPI) CreateTaskGraph(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Add a set of tasks to an existing task-graph. The request format is very
// similar to the request format for creating task-graphs. But `routes`
// key, `scopes`, `metadata` and `tags` cannot be modified.
//
// **Referencing required tasks**, just as when task-graphs are created,
// each task has a list of required tasks. It is possible to reference
// all `taskId`s within the task-graph.
//
// **Safety,** it is only _safe_ to call this API end-point while the
// task-graph being modified is still running. If the task-graph is
// _finished_ or _blocked_, this method will leave the task-graph in this
// state. Hence, it is only truly _safe_ to call this API end-point from
// within a task in the task-graph being modified.
func (a SchedulerAPI) ExtendTaskGraph(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get task-graph status, this will return the _task-graph status
// structure_. which can be used to check if a task-graph is `running`,
// `blocked` or `finished`.
//
// **Note**, that `finished` implies successfully completion.
func (a SchedulerAPI) Status(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Get task-graph information, this includes the _task-graph status
// structure_, along with `metadata` and `tags`, but not information
// about all tasks.
//
// If you want more detailed information use the `inspectTaskGraph`
// end-point instead.
func (a SchedulerAPI) Info(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Inspect a task-graph, this returns all the information the task-graph
// scheduler knows about the task-graph and the state of its tasks.
//
// **Warning**, some of these fields are borderline internal to the
// task-graph scheduler and we may choose to change or make them internal
// later. Also note that note all of the information is formalized yet.
// The JSON schema will be updated to reflect formalized values, we think
// it's safe to consider the values stable.
//
// Take these considerations into account when using the API end-point,
// as we do not promise it will remain fully backward compatible in
// the future.
func (a SchedulerAPI) Inspect(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Inspect a task from a task-graph, this returns all the information the
// task-graph scheduler knows about the specific task.
//
// **Warning**, some of these fields are borderline internal to the
// task-graph scheduler and we may choose to change or make them internal
// later. Also note that note all of the information is formalized yet.
// The JSON schema will be updated to reflect formalized values, we think
// it's safe to consider the values stable.
//
// Take these considerations into account when using the API end-point,
// as we do not promise it will remain fully backward compatible in
// the future.
func (a SchedulerAPI) InspectTask(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}

// Documented later...
//
// **Warning** this api end-point is **not stable**.
func (a SchedulerAPI) Ping(clientId string) GetClientScopesResponse {
	return apiCall().(GetClientScopesResponse)
}