// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package snowdevicemanagementiface provides an interface to enable mocking the AWS Snow Device Management service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package snowdevicemanagementiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/snowdevicemanagement"
)

// SnowDeviceManagementAPI provides an interface to enable mocking the
// snowdevicemanagement.SnowDeviceManagement service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Snow Device Management.
//	func myFunc(svc snowdevicemanagementiface.SnowDeviceManagementAPI) bool {
//	    // Make svc.CancelTask request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := snowdevicemanagement.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSnowDeviceManagementClient struct {
//	    snowdevicemanagementiface.SnowDeviceManagementAPI
//	}
//	func (m *mockSnowDeviceManagementClient) CancelTask(input *snowdevicemanagement.CancelTaskInput) (*snowdevicemanagement.CancelTaskOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSnowDeviceManagementClient{}
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
type SnowDeviceManagementAPI interface {
	CancelTask(*snowdevicemanagement.CancelTaskInput) (*snowdevicemanagement.CancelTaskOutput, error)
	CancelTaskWithContext(aws.Context, *snowdevicemanagement.CancelTaskInput, ...request.Option) (*snowdevicemanagement.CancelTaskOutput, error)
	CancelTaskRequest(*snowdevicemanagement.CancelTaskInput) (*request.Request, *snowdevicemanagement.CancelTaskOutput)

	CreateTask(*snowdevicemanagement.CreateTaskInput) (*snowdevicemanagement.CreateTaskOutput, error)
	CreateTaskWithContext(aws.Context, *snowdevicemanagement.CreateTaskInput, ...request.Option) (*snowdevicemanagement.CreateTaskOutput, error)
	CreateTaskRequest(*snowdevicemanagement.CreateTaskInput) (*request.Request, *snowdevicemanagement.CreateTaskOutput)

	DescribeDevice(*snowdevicemanagement.DescribeDeviceInput) (*snowdevicemanagement.DescribeDeviceOutput, error)
	DescribeDeviceWithContext(aws.Context, *snowdevicemanagement.DescribeDeviceInput, ...request.Option) (*snowdevicemanagement.DescribeDeviceOutput, error)
	DescribeDeviceRequest(*snowdevicemanagement.DescribeDeviceInput) (*request.Request, *snowdevicemanagement.DescribeDeviceOutput)

	DescribeDeviceEc2Instances(*snowdevicemanagement.DescribeDeviceEc2InstancesInput) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error)
	DescribeDeviceEc2InstancesWithContext(aws.Context, *snowdevicemanagement.DescribeDeviceEc2InstancesInput, ...request.Option) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error)
	DescribeDeviceEc2InstancesRequest(*snowdevicemanagement.DescribeDeviceEc2InstancesInput) (*request.Request, *snowdevicemanagement.DescribeDeviceEc2InstancesOutput)

	DescribeExecution(*snowdevicemanagement.DescribeExecutionInput) (*snowdevicemanagement.DescribeExecutionOutput, error)
	DescribeExecutionWithContext(aws.Context, *snowdevicemanagement.DescribeExecutionInput, ...request.Option) (*snowdevicemanagement.DescribeExecutionOutput, error)
	DescribeExecutionRequest(*snowdevicemanagement.DescribeExecutionInput) (*request.Request, *snowdevicemanagement.DescribeExecutionOutput)

	DescribeTask(*snowdevicemanagement.DescribeTaskInput) (*snowdevicemanagement.DescribeTaskOutput, error)
	DescribeTaskWithContext(aws.Context, *snowdevicemanagement.DescribeTaskInput, ...request.Option) (*snowdevicemanagement.DescribeTaskOutput, error)
	DescribeTaskRequest(*snowdevicemanagement.DescribeTaskInput) (*request.Request, *snowdevicemanagement.DescribeTaskOutput)

	ListDeviceResources(*snowdevicemanagement.ListDeviceResourcesInput) (*snowdevicemanagement.ListDeviceResourcesOutput, error)
	ListDeviceResourcesWithContext(aws.Context, *snowdevicemanagement.ListDeviceResourcesInput, ...request.Option) (*snowdevicemanagement.ListDeviceResourcesOutput, error)
	ListDeviceResourcesRequest(*snowdevicemanagement.ListDeviceResourcesInput) (*request.Request, *snowdevicemanagement.ListDeviceResourcesOutput)

	ListDeviceResourcesPages(*snowdevicemanagement.ListDeviceResourcesInput, func(*snowdevicemanagement.ListDeviceResourcesOutput, bool) bool) error
	ListDeviceResourcesPagesWithContext(aws.Context, *snowdevicemanagement.ListDeviceResourcesInput, func(*snowdevicemanagement.ListDeviceResourcesOutput, bool) bool, ...request.Option) error

	ListDevices(*snowdevicemanagement.ListDevicesInput) (*snowdevicemanagement.ListDevicesOutput, error)
	ListDevicesWithContext(aws.Context, *snowdevicemanagement.ListDevicesInput, ...request.Option) (*snowdevicemanagement.ListDevicesOutput, error)
	ListDevicesRequest(*snowdevicemanagement.ListDevicesInput) (*request.Request, *snowdevicemanagement.ListDevicesOutput)

	ListDevicesPages(*snowdevicemanagement.ListDevicesInput, func(*snowdevicemanagement.ListDevicesOutput, bool) bool) error
	ListDevicesPagesWithContext(aws.Context, *snowdevicemanagement.ListDevicesInput, func(*snowdevicemanagement.ListDevicesOutput, bool) bool, ...request.Option) error

	ListExecutions(*snowdevicemanagement.ListExecutionsInput) (*snowdevicemanagement.ListExecutionsOutput, error)
	ListExecutionsWithContext(aws.Context, *snowdevicemanagement.ListExecutionsInput, ...request.Option) (*snowdevicemanagement.ListExecutionsOutput, error)
	ListExecutionsRequest(*snowdevicemanagement.ListExecutionsInput) (*request.Request, *snowdevicemanagement.ListExecutionsOutput)

	ListExecutionsPages(*snowdevicemanagement.ListExecutionsInput, func(*snowdevicemanagement.ListExecutionsOutput, bool) bool) error
	ListExecutionsPagesWithContext(aws.Context, *snowdevicemanagement.ListExecutionsInput, func(*snowdevicemanagement.ListExecutionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*snowdevicemanagement.ListTagsForResourceInput) (*snowdevicemanagement.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *snowdevicemanagement.ListTagsForResourceInput, ...request.Option) (*snowdevicemanagement.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*snowdevicemanagement.ListTagsForResourceInput) (*request.Request, *snowdevicemanagement.ListTagsForResourceOutput)

	ListTasks(*snowdevicemanagement.ListTasksInput) (*snowdevicemanagement.ListTasksOutput, error)
	ListTasksWithContext(aws.Context, *snowdevicemanagement.ListTasksInput, ...request.Option) (*snowdevicemanagement.ListTasksOutput, error)
	ListTasksRequest(*snowdevicemanagement.ListTasksInput) (*request.Request, *snowdevicemanagement.ListTasksOutput)

	ListTasksPages(*snowdevicemanagement.ListTasksInput, func(*snowdevicemanagement.ListTasksOutput, bool) bool) error
	ListTasksPagesWithContext(aws.Context, *snowdevicemanagement.ListTasksInput, func(*snowdevicemanagement.ListTasksOutput, bool) bool, ...request.Option) error

	TagResource(*snowdevicemanagement.TagResourceInput) (*snowdevicemanagement.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *snowdevicemanagement.TagResourceInput, ...request.Option) (*snowdevicemanagement.TagResourceOutput, error)
	TagResourceRequest(*snowdevicemanagement.TagResourceInput) (*request.Request, *snowdevicemanagement.TagResourceOutput)

	UntagResource(*snowdevicemanagement.UntagResourceInput) (*snowdevicemanagement.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *snowdevicemanagement.UntagResourceInput, ...request.Option) (*snowdevicemanagement.UntagResourceOutput, error)
	UntagResourceRequest(*snowdevicemanagement.UntagResourceInput) (*request.Request, *snowdevicemanagement.UntagResourceOutput)
}

var _ SnowDeviceManagementAPI = (*snowdevicemanagement.SnowDeviceManagement)(nil)
