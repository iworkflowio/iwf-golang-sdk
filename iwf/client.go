package iwf

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

// Client is a full-featured client
type Client interface {
	clientCommon
	// StartWorkflow starts a workflow execution
	// startStateId is the first stateId to start
	// workflowId is the required identifier for the workflow execution(see Cadence/Temporal for more details about WorkflowId uniqueness)
	// timeoutSecs is required as the workflow execution timeout in seconds
	// input can be optional, it's the input for the startState
	// options is optional includes like IdReusePolicy, RetryPolicy, CronSchedule and also WorkflowStateOptions for the startState. Empty by default(when nil).
	// return the workflowRunId
	StartWorkflow(ctx context.Context, workflow Workflow, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error)
	// SignalWorkflow signals a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// signalChannelName is required, signalValue is optional(for case of empty value)
	SignalWorkflow(ctx context.Context, workflow Workflow, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error
	// GetWorkflowDataObjects returns the data objects of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowDataObjects API instead
	GetWorkflowDataObjects(ctx context.Context, workflow Workflow, workflowId, workflowRunId string, keys []string) (map[string]Object, error)
	// GetWorkflowSearchAttributes returns search attributes of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowSearchAttributes API instead
	GetWorkflowSearchAttributes(ctx context.Context, workflow Workflow, workflowId, workflowRunId string, keys []string) (map[string]interface{}, error)
	// GetAllWorkflowSearchAttributes returns all search attributes of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	GetAllWorkflowSearchAttributes(ctx context.Context, workflow Workflow, workflowId, workflowRunId string) (map[string]interface{}, error)
}

// clientCommon is the common APIs between Client and UnregisteredClient
type clientCommon interface {
	// StopWorkflow stops a workflow execution.
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// options is optional, default (when nil)to use Cancel as stopType
	StopWorkflow(ctx context.Context, workflowId, workflowRunId string, options *WorkflowStopOptions) error
	// GetSimpleWorkflowResult returns the result of a workflow execution, for simple case that only one WorkflowState completes with result
	// If there are more than one WorkflowStates complete with result, GetComplexWorkflowResults must be used instead
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// resultPtr is the pointer to retrieve the result
	GetSimpleWorkflowResult(ctx context.Context, workflowId, workflowRunId string, resultPtr interface{}) error
	// GetComplexWorkflowResults returns the results of a workflow execution
	// It returns a list of iwfidl.StateCompletionOutput and user code will have to use ObjectEncoder to deserialize
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	GetComplexWorkflowResults(ctx context.Context, workflowId, workflowRunId string) ([]iwfidl.StateCompletionOutput, error)
	// ResetWorkflow resets a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// resetWorkflowTypeAndOptions is optional, it provides combination parameter for reset. Default (when nil) will reset to iwfidl.BEGINNING resetType
	// return the workflowRunId
	ResetWorkflow(ctx context.Context, workflowId, workflowRunId string, options *ResetWorkflowTypeAndOptions) (string, error)
	// DescribeWorkflow describes the basic info of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error)

	// SearchWorkflow searches for workflow executions given a query (see SearchAttribute query in Cadence/Temporal)
	//  https://cadenceworkflow.io/docs/concepts/search-workflows/
	//  https://docs.temporal.io/concepts/what-is-a-search-attribute/
	SearchWorkflow(ctx context.Context, request iwfidl.WorkflowSearchRequest) (*iwfidl.WorkflowSearchResponse, error)
	// GetAllWorkflowDataObjects returns all the data objects of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	GetAllWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string) (map[string]Object, error)
}

// UnregisteredClient is a client without workflow registry
type UnregisteredClient interface {
	clientCommon
	// StartWorkflow starts a workflow execution
	// startStateId is the first stateId to start
	// workflowId is the required identifier for the workflow execution(see Cadence/Temporal for more details about WorkflowId uniqueness)
	// timeoutSecs is required as the workflow execution timeout in seconds
	// input can be optional, it's the input for the startState
	// options is optional includes like IdReusePolicy, RetryPolicy, CronSchedule and also WorkflowStateOptions for the startState. Empty by default(when nil).
	// return the workflowRunId
	StartWorkflow(ctx context.Context, workflowType string, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error)
	// SignalWorkflow signals a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// signalChannelName is required, signalValue is optional(for case of empty value)
	SignalWorkflow(ctx context.Context, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error
	// GetWorkflowDataObjects returns the data objects of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowDataObjects API instead
	GetWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string, keys []string) (map[string]Object, error)
	// GetWorkflowSearchAttributes returns search attributes of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowSearchAttributes API instead
	GetWorkflowSearchAttributes(ctx context.Context, workflowId, workflowRunId string, keys []iwfidl.SearchAttributeKeyAndType) (map[string]iwfidl.SearchAttribute, error)
}

// NewUnregisteredClient returns a UnregisteredClient
// It will let you invoke the APIs to iWF server without much type validation checks(workflow type, channel names, etc).
// It's useful for calling Client APIs without workflow registry(which may require to have all the workflow dependencies)
func NewUnregisteredClient(options *ClientOptions) UnregisteredClient {
	if options == nil {
		options = GetLocalDefaultClientOptions()
	}

	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: options.ServerUrl,
			},
		},
	})

	return &unregisteredClientImpl{
		options:   options,
		apiClient: apiClient,
	}
}

// NewClient returns a Client
// It requires a registry in order to perform validation checks (workflow type, channel names, etc)
// Use NewUnregisteredClient if you don't have a registry in your application
func NewClient(registry Registry, options *ClientOptions) Client {
	if registry == nil {
		panic("cannot have nil registry")
	}
	return &clientImpl{
		UnregisteredClient: NewUnregisteredClient(options),
		registry:           registry,
		options:            options,
	}
}
