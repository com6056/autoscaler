// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package internetmonitoriface provides an interface to enable mocking the Amazon CloudWatch Internet Monitor service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package internetmonitoriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/internetmonitor"
)

// InternetMonitorAPI provides an interface to enable mocking the
// internetmonitor.InternetMonitor service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon CloudWatch Internet Monitor.
//	func myFunc(svc internetmonitoriface.InternetMonitorAPI) bool {
//	    // Make svc.CreateMonitor request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := internetmonitor.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockInternetMonitorClient struct {
//	    internetmonitoriface.InternetMonitorAPI
//	}
//	func (m *mockInternetMonitorClient) CreateMonitor(input *internetmonitor.CreateMonitorInput) (*internetmonitor.CreateMonitorOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockInternetMonitorClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type InternetMonitorAPI interface {
	CreateMonitor(*internetmonitor.CreateMonitorInput) (*internetmonitor.CreateMonitorOutput, error)
	CreateMonitorWithContext(aws.Context, *internetmonitor.CreateMonitorInput, ...request.Option) (*internetmonitor.CreateMonitorOutput, error)
	CreateMonitorRequest(*internetmonitor.CreateMonitorInput) (*request.Request, *internetmonitor.CreateMonitorOutput)

	DeleteMonitor(*internetmonitor.DeleteMonitorInput) (*internetmonitor.DeleteMonitorOutput, error)
	DeleteMonitorWithContext(aws.Context, *internetmonitor.DeleteMonitorInput, ...request.Option) (*internetmonitor.DeleteMonitorOutput, error)
	DeleteMonitorRequest(*internetmonitor.DeleteMonitorInput) (*request.Request, *internetmonitor.DeleteMonitorOutput)

	GetHealthEvent(*internetmonitor.GetHealthEventInput) (*internetmonitor.GetHealthEventOutput, error)
	GetHealthEventWithContext(aws.Context, *internetmonitor.GetHealthEventInput, ...request.Option) (*internetmonitor.GetHealthEventOutput, error)
	GetHealthEventRequest(*internetmonitor.GetHealthEventInput) (*request.Request, *internetmonitor.GetHealthEventOutput)

	GetMonitor(*internetmonitor.GetMonitorInput) (*internetmonitor.GetMonitorOutput, error)
	GetMonitorWithContext(aws.Context, *internetmonitor.GetMonitorInput, ...request.Option) (*internetmonitor.GetMonitorOutput, error)
	GetMonitorRequest(*internetmonitor.GetMonitorInput) (*request.Request, *internetmonitor.GetMonitorOutput)

	GetQueryResults(*internetmonitor.GetQueryResultsInput) (*internetmonitor.GetQueryResultsOutput, error)
	GetQueryResultsWithContext(aws.Context, *internetmonitor.GetQueryResultsInput, ...request.Option) (*internetmonitor.GetQueryResultsOutput, error)
	GetQueryResultsRequest(*internetmonitor.GetQueryResultsInput) (*request.Request, *internetmonitor.GetQueryResultsOutput)

	GetQueryResultsPages(*internetmonitor.GetQueryResultsInput, func(*internetmonitor.GetQueryResultsOutput, bool) bool) error
	GetQueryResultsPagesWithContext(aws.Context, *internetmonitor.GetQueryResultsInput, func(*internetmonitor.GetQueryResultsOutput, bool) bool, ...request.Option) error

	GetQueryStatus(*internetmonitor.GetQueryStatusInput) (*internetmonitor.GetQueryStatusOutput, error)
	GetQueryStatusWithContext(aws.Context, *internetmonitor.GetQueryStatusInput, ...request.Option) (*internetmonitor.GetQueryStatusOutput, error)
	GetQueryStatusRequest(*internetmonitor.GetQueryStatusInput) (*request.Request, *internetmonitor.GetQueryStatusOutput)

	ListHealthEvents(*internetmonitor.ListHealthEventsInput) (*internetmonitor.ListHealthEventsOutput, error)
	ListHealthEventsWithContext(aws.Context, *internetmonitor.ListHealthEventsInput, ...request.Option) (*internetmonitor.ListHealthEventsOutput, error)
	ListHealthEventsRequest(*internetmonitor.ListHealthEventsInput) (*request.Request, *internetmonitor.ListHealthEventsOutput)

	ListHealthEventsPages(*internetmonitor.ListHealthEventsInput, func(*internetmonitor.ListHealthEventsOutput, bool) bool) error
	ListHealthEventsPagesWithContext(aws.Context, *internetmonitor.ListHealthEventsInput, func(*internetmonitor.ListHealthEventsOutput, bool) bool, ...request.Option) error

	ListMonitors(*internetmonitor.ListMonitorsInput) (*internetmonitor.ListMonitorsOutput, error)
	ListMonitorsWithContext(aws.Context, *internetmonitor.ListMonitorsInput, ...request.Option) (*internetmonitor.ListMonitorsOutput, error)
	ListMonitorsRequest(*internetmonitor.ListMonitorsInput) (*request.Request, *internetmonitor.ListMonitorsOutput)

	ListMonitorsPages(*internetmonitor.ListMonitorsInput, func(*internetmonitor.ListMonitorsOutput, bool) bool) error
	ListMonitorsPagesWithContext(aws.Context, *internetmonitor.ListMonitorsInput, func(*internetmonitor.ListMonitorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*internetmonitor.ListTagsForResourceInput) (*internetmonitor.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *internetmonitor.ListTagsForResourceInput, ...request.Option) (*internetmonitor.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*internetmonitor.ListTagsForResourceInput) (*request.Request, *internetmonitor.ListTagsForResourceOutput)

	StartQuery(*internetmonitor.StartQueryInput) (*internetmonitor.StartQueryOutput, error)
	StartQueryWithContext(aws.Context, *internetmonitor.StartQueryInput, ...request.Option) (*internetmonitor.StartQueryOutput, error)
	StartQueryRequest(*internetmonitor.StartQueryInput) (*request.Request, *internetmonitor.StartQueryOutput)

	StopQuery(*internetmonitor.StopQueryInput) (*internetmonitor.StopQueryOutput, error)
	StopQueryWithContext(aws.Context, *internetmonitor.StopQueryInput, ...request.Option) (*internetmonitor.StopQueryOutput, error)
	StopQueryRequest(*internetmonitor.StopQueryInput) (*request.Request, *internetmonitor.StopQueryOutput)

	TagResource(*internetmonitor.TagResourceInput) (*internetmonitor.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *internetmonitor.TagResourceInput, ...request.Option) (*internetmonitor.TagResourceOutput, error)
	TagResourceRequest(*internetmonitor.TagResourceInput) (*request.Request, *internetmonitor.TagResourceOutput)

	UntagResource(*internetmonitor.UntagResourceInput) (*internetmonitor.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *internetmonitor.UntagResourceInput, ...request.Option) (*internetmonitor.UntagResourceOutput, error)
	UntagResourceRequest(*internetmonitor.UntagResourceInput) (*request.Request, *internetmonitor.UntagResourceOutput)

	UpdateMonitor(*internetmonitor.UpdateMonitorInput) (*internetmonitor.UpdateMonitorOutput, error)
	UpdateMonitorWithContext(aws.Context, *internetmonitor.UpdateMonitorInput, ...request.Option) (*internetmonitor.UpdateMonitorOutput, error)
	UpdateMonitorRequest(*internetmonitor.UpdateMonitorInput) (*request.Request, *internetmonitor.UpdateMonitorOutput)
}

var _ InternetMonitorAPI = (*internetmonitor.InternetMonitor)(nil)
