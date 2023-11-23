// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appstreamiface provides an interface to enable mocking the Amazon AppStream service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appstreamiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appstream"
)

// AppStreamAPI provides an interface to enable mocking the
// appstream.AppStream service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon AppStream.
//	func myFunc(svc appstreamiface.AppStreamAPI) bool {
//	    // Make svc.AssociateAppBlockBuilderAppBlock request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := appstream.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAppStreamClient struct {
//	    appstreamiface.AppStreamAPI
//	}
//	func (m *mockAppStreamClient) AssociateAppBlockBuilderAppBlock(input *appstream.AssociateAppBlockBuilderAppBlockInput) (*appstream.AssociateAppBlockBuilderAppBlockOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAppStreamClient{}
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
type AppStreamAPI interface {
	AssociateAppBlockBuilderAppBlock(*appstream.AssociateAppBlockBuilderAppBlockInput) (*appstream.AssociateAppBlockBuilderAppBlockOutput, error)
	AssociateAppBlockBuilderAppBlockWithContext(aws.Context, *appstream.AssociateAppBlockBuilderAppBlockInput, ...request.Option) (*appstream.AssociateAppBlockBuilderAppBlockOutput, error)
	AssociateAppBlockBuilderAppBlockRequest(*appstream.AssociateAppBlockBuilderAppBlockInput) (*request.Request, *appstream.AssociateAppBlockBuilderAppBlockOutput)

	AssociateApplicationFleet(*appstream.AssociateApplicationFleetInput) (*appstream.AssociateApplicationFleetOutput, error)
	AssociateApplicationFleetWithContext(aws.Context, *appstream.AssociateApplicationFleetInput, ...request.Option) (*appstream.AssociateApplicationFleetOutput, error)
	AssociateApplicationFleetRequest(*appstream.AssociateApplicationFleetInput) (*request.Request, *appstream.AssociateApplicationFleetOutput)

	AssociateApplicationToEntitlement(*appstream.AssociateApplicationToEntitlementInput) (*appstream.AssociateApplicationToEntitlementOutput, error)
	AssociateApplicationToEntitlementWithContext(aws.Context, *appstream.AssociateApplicationToEntitlementInput, ...request.Option) (*appstream.AssociateApplicationToEntitlementOutput, error)
	AssociateApplicationToEntitlementRequest(*appstream.AssociateApplicationToEntitlementInput) (*request.Request, *appstream.AssociateApplicationToEntitlementOutput)

	AssociateFleet(*appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error)
	AssociateFleetWithContext(aws.Context, *appstream.AssociateFleetInput, ...request.Option) (*appstream.AssociateFleetOutput, error)
	AssociateFleetRequest(*appstream.AssociateFleetInput) (*request.Request, *appstream.AssociateFleetOutput)

	BatchAssociateUserStack(*appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error)
	BatchAssociateUserStackWithContext(aws.Context, *appstream.BatchAssociateUserStackInput, ...request.Option) (*appstream.BatchAssociateUserStackOutput, error)
	BatchAssociateUserStackRequest(*appstream.BatchAssociateUserStackInput) (*request.Request, *appstream.BatchAssociateUserStackOutput)

	BatchDisassociateUserStack(*appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error)
	BatchDisassociateUserStackWithContext(aws.Context, *appstream.BatchDisassociateUserStackInput, ...request.Option) (*appstream.BatchDisassociateUserStackOutput, error)
	BatchDisassociateUserStackRequest(*appstream.BatchDisassociateUserStackInput) (*request.Request, *appstream.BatchDisassociateUserStackOutput)

	CopyImage(*appstream.CopyImageInput) (*appstream.CopyImageOutput, error)
	CopyImageWithContext(aws.Context, *appstream.CopyImageInput, ...request.Option) (*appstream.CopyImageOutput, error)
	CopyImageRequest(*appstream.CopyImageInput) (*request.Request, *appstream.CopyImageOutput)

	CreateAppBlock(*appstream.CreateAppBlockInput) (*appstream.CreateAppBlockOutput, error)
	CreateAppBlockWithContext(aws.Context, *appstream.CreateAppBlockInput, ...request.Option) (*appstream.CreateAppBlockOutput, error)
	CreateAppBlockRequest(*appstream.CreateAppBlockInput) (*request.Request, *appstream.CreateAppBlockOutput)

	CreateAppBlockBuilder(*appstream.CreateAppBlockBuilderInput) (*appstream.CreateAppBlockBuilderOutput, error)
	CreateAppBlockBuilderWithContext(aws.Context, *appstream.CreateAppBlockBuilderInput, ...request.Option) (*appstream.CreateAppBlockBuilderOutput, error)
	CreateAppBlockBuilderRequest(*appstream.CreateAppBlockBuilderInput) (*request.Request, *appstream.CreateAppBlockBuilderOutput)

	CreateAppBlockBuilderStreamingURL(*appstream.CreateAppBlockBuilderStreamingURLInput) (*appstream.CreateAppBlockBuilderStreamingURLOutput, error)
	CreateAppBlockBuilderStreamingURLWithContext(aws.Context, *appstream.CreateAppBlockBuilderStreamingURLInput, ...request.Option) (*appstream.CreateAppBlockBuilderStreamingURLOutput, error)
	CreateAppBlockBuilderStreamingURLRequest(*appstream.CreateAppBlockBuilderStreamingURLInput) (*request.Request, *appstream.CreateAppBlockBuilderStreamingURLOutput)

	CreateApplication(*appstream.CreateApplicationInput) (*appstream.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *appstream.CreateApplicationInput, ...request.Option) (*appstream.CreateApplicationOutput, error)
	CreateApplicationRequest(*appstream.CreateApplicationInput) (*request.Request, *appstream.CreateApplicationOutput)

	CreateDirectoryConfig(*appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error)
	CreateDirectoryConfigWithContext(aws.Context, *appstream.CreateDirectoryConfigInput, ...request.Option) (*appstream.CreateDirectoryConfigOutput, error)
	CreateDirectoryConfigRequest(*appstream.CreateDirectoryConfigInput) (*request.Request, *appstream.CreateDirectoryConfigOutput)

	CreateEntitlement(*appstream.CreateEntitlementInput) (*appstream.CreateEntitlementOutput, error)
	CreateEntitlementWithContext(aws.Context, *appstream.CreateEntitlementInput, ...request.Option) (*appstream.CreateEntitlementOutput, error)
	CreateEntitlementRequest(*appstream.CreateEntitlementInput) (*request.Request, *appstream.CreateEntitlementOutput)

	CreateFleet(*appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error)
	CreateFleetWithContext(aws.Context, *appstream.CreateFleetInput, ...request.Option) (*appstream.CreateFleetOutput, error)
	CreateFleetRequest(*appstream.CreateFleetInput) (*request.Request, *appstream.CreateFleetOutput)

	CreateImageBuilder(*appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderWithContext(aws.Context, *appstream.CreateImageBuilderInput, ...request.Option) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderRequest(*appstream.CreateImageBuilderInput) (*request.Request, *appstream.CreateImageBuilderOutput)

	CreateImageBuilderStreamingURL(*appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateImageBuilderStreamingURLWithContext(aws.Context, *appstream.CreateImageBuilderStreamingURLInput, ...request.Option) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateImageBuilderStreamingURLRequest(*appstream.CreateImageBuilderStreamingURLInput) (*request.Request, *appstream.CreateImageBuilderStreamingURLOutput)

	CreateStack(*appstream.CreateStackInput) (*appstream.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *appstream.CreateStackInput, ...request.Option) (*appstream.CreateStackOutput, error)
	CreateStackRequest(*appstream.CreateStackInput) (*request.Request, *appstream.CreateStackOutput)

	CreateStreamingURL(*appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error)
	CreateStreamingURLWithContext(aws.Context, *appstream.CreateStreamingURLInput, ...request.Option) (*appstream.CreateStreamingURLOutput, error)
	CreateStreamingURLRequest(*appstream.CreateStreamingURLInput) (*request.Request, *appstream.CreateStreamingURLOutput)

	CreateUpdatedImage(*appstream.CreateUpdatedImageInput) (*appstream.CreateUpdatedImageOutput, error)
	CreateUpdatedImageWithContext(aws.Context, *appstream.CreateUpdatedImageInput, ...request.Option) (*appstream.CreateUpdatedImageOutput, error)
	CreateUpdatedImageRequest(*appstream.CreateUpdatedImageInput) (*request.Request, *appstream.CreateUpdatedImageOutput)

	CreateUsageReportSubscription(*appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error)
	CreateUsageReportSubscriptionWithContext(aws.Context, *appstream.CreateUsageReportSubscriptionInput, ...request.Option) (*appstream.CreateUsageReportSubscriptionOutput, error)
	CreateUsageReportSubscriptionRequest(*appstream.CreateUsageReportSubscriptionInput) (*request.Request, *appstream.CreateUsageReportSubscriptionOutput)

	CreateUser(*appstream.CreateUserInput) (*appstream.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *appstream.CreateUserInput, ...request.Option) (*appstream.CreateUserOutput, error)
	CreateUserRequest(*appstream.CreateUserInput) (*request.Request, *appstream.CreateUserOutput)

	DeleteAppBlock(*appstream.DeleteAppBlockInput) (*appstream.DeleteAppBlockOutput, error)
	DeleteAppBlockWithContext(aws.Context, *appstream.DeleteAppBlockInput, ...request.Option) (*appstream.DeleteAppBlockOutput, error)
	DeleteAppBlockRequest(*appstream.DeleteAppBlockInput) (*request.Request, *appstream.DeleteAppBlockOutput)

	DeleteAppBlockBuilder(*appstream.DeleteAppBlockBuilderInput) (*appstream.DeleteAppBlockBuilderOutput, error)
	DeleteAppBlockBuilderWithContext(aws.Context, *appstream.DeleteAppBlockBuilderInput, ...request.Option) (*appstream.DeleteAppBlockBuilderOutput, error)
	DeleteAppBlockBuilderRequest(*appstream.DeleteAppBlockBuilderInput) (*request.Request, *appstream.DeleteAppBlockBuilderOutput)

	DeleteApplication(*appstream.DeleteApplicationInput) (*appstream.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *appstream.DeleteApplicationInput, ...request.Option) (*appstream.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*appstream.DeleteApplicationInput) (*request.Request, *appstream.DeleteApplicationOutput)

	DeleteDirectoryConfig(*appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteDirectoryConfigWithContext(aws.Context, *appstream.DeleteDirectoryConfigInput, ...request.Option) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteDirectoryConfigRequest(*appstream.DeleteDirectoryConfigInput) (*request.Request, *appstream.DeleteDirectoryConfigOutput)

	DeleteEntitlement(*appstream.DeleteEntitlementInput) (*appstream.DeleteEntitlementOutput, error)
	DeleteEntitlementWithContext(aws.Context, *appstream.DeleteEntitlementInput, ...request.Option) (*appstream.DeleteEntitlementOutput, error)
	DeleteEntitlementRequest(*appstream.DeleteEntitlementInput) (*request.Request, *appstream.DeleteEntitlementOutput)

	DeleteFleet(*appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error)
	DeleteFleetWithContext(aws.Context, *appstream.DeleteFleetInput, ...request.Option) (*appstream.DeleteFleetOutput, error)
	DeleteFleetRequest(*appstream.DeleteFleetInput) (*request.Request, *appstream.DeleteFleetOutput)

	DeleteImage(*appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error)
	DeleteImageWithContext(aws.Context, *appstream.DeleteImageInput, ...request.Option) (*appstream.DeleteImageOutput, error)
	DeleteImageRequest(*appstream.DeleteImageInput) (*request.Request, *appstream.DeleteImageOutput)

	DeleteImageBuilder(*appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImageBuilderWithContext(aws.Context, *appstream.DeleteImageBuilderInput, ...request.Option) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImageBuilderRequest(*appstream.DeleteImageBuilderInput) (*request.Request, *appstream.DeleteImageBuilderOutput)

	DeleteImagePermissions(*appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error)
	DeleteImagePermissionsWithContext(aws.Context, *appstream.DeleteImagePermissionsInput, ...request.Option) (*appstream.DeleteImagePermissionsOutput, error)
	DeleteImagePermissionsRequest(*appstream.DeleteImagePermissionsInput) (*request.Request, *appstream.DeleteImagePermissionsOutput)

	DeleteStack(*appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *appstream.DeleteStackInput, ...request.Option) (*appstream.DeleteStackOutput, error)
	DeleteStackRequest(*appstream.DeleteStackInput) (*request.Request, *appstream.DeleteStackOutput)

	DeleteUsageReportSubscription(*appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error)
	DeleteUsageReportSubscriptionWithContext(aws.Context, *appstream.DeleteUsageReportSubscriptionInput, ...request.Option) (*appstream.DeleteUsageReportSubscriptionOutput, error)
	DeleteUsageReportSubscriptionRequest(*appstream.DeleteUsageReportSubscriptionInput) (*request.Request, *appstream.DeleteUsageReportSubscriptionOutput)

	DeleteUser(*appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *appstream.DeleteUserInput, ...request.Option) (*appstream.DeleteUserOutput, error)
	DeleteUserRequest(*appstream.DeleteUserInput) (*request.Request, *appstream.DeleteUserOutput)

	DescribeAppBlockBuilderAppBlockAssociations(*appstream.DescribeAppBlockBuilderAppBlockAssociationsInput) (*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, error)
	DescribeAppBlockBuilderAppBlockAssociationsWithContext(aws.Context, *appstream.DescribeAppBlockBuilderAppBlockAssociationsInput, ...request.Option) (*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, error)
	DescribeAppBlockBuilderAppBlockAssociationsRequest(*appstream.DescribeAppBlockBuilderAppBlockAssociationsInput) (*request.Request, *appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput)

	DescribeAppBlockBuilderAppBlockAssociationsPages(*appstream.DescribeAppBlockBuilderAppBlockAssociationsInput, func(*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, bool) bool) error
	DescribeAppBlockBuilderAppBlockAssociationsPagesWithContext(aws.Context, *appstream.DescribeAppBlockBuilderAppBlockAssociationsInput, func(*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, bool) bool, ...request.Option) error

	DescribeAppBlockBuilders(*appstream.DescribeAppBlockBuildersInput) (*appstream.DescribeAppBlockBuildersOutput, error)
	DescribeAppBlockBuildersWithContext(aws.Context, *appstream.DescribeAppBlockBuildersInput, ...request.Option) (*appstream.DescribeAppBlockBuildersOutput, error)
	DescribeAppBlockBuildersRequest(*appstream.DescribeAppBlockBuildersInput) (*request.Request, *appstream.DescribeAppBlockBuildersOutput)

	DescribeAppBlockBuildersPages(*appstream.DescribeAppBlockBuildersInput, func(*appstream.DescribeAppBlockBuildersOutput, bool) bool) error
	DescribeAppBlockBuildersPagesWithContext(aws.Context, *appstream.DescribeAppBlockBuildersInput, func(*appstream.DescribeAppBlockBuildersOutput, bool) bool, ...request.Option) error

	DescribeAppBlocks(*appstream.DescribeAppBlocksInput) (*appstream.DescribeAppBlocksOutput, error)
	DescribeAppBlocksWithContext(aws.Context, *appstream.DescribeAppBlocksInput, ...request.Option) (*appstream.DescribeAppBlocksOutput, error)
	DescribeAppBlocksRequest(*appstream.DescribeAppBlocksInput) (*request.Request, *appstream.DescribeAppBlocksOutput)

	DescribeApplicationFleetAssociations(*appstream.DescribeApplicationFleetAssociationsInput) (*appstream.DescribeApplicationFleetAssociationsOutput, error)
	DescribeApplicationFleetAssociationsWithContext(aws.Context, *appstream.DescribeApplicationFleetAssociationsInput, ...request.Option) (*appstream.DescribeApplicationFleetAssociationsOutput, error)
	DescribeApplicationFleetAssociationsRequest(*appstream.DescribeApplicationFleetAssociationsInput) (*request.Request, *appstream.DescribeApplicationFleetAssociationsOutput)

	DescribeApplications(*appstream.DescribeApplicationsInput) (*appstream.DescribeApplicationsOutput, error)
	DescribeApplicationsWithContext(aws.Context, *appstream.DescribeApplicationsInput, ...request.Option) (*appstream.DescribeApplicationsOutput, error)
	DescribeApplicationsRequest(*appstream.DescribeApplicationsInput) (*request.Request, *appstream.DescribeApplicationsOutput)

	DescribeDirectoryConfigs(*appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeDirectoryConfigsWithContext(aws.Context, *appstream.DescribeDirectoryConfigsInput, ...request.Option) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeDirectoryConfigsRequest(*appstream.DescribeDirectoryConfigsInput) (*request.Request, *appstream.DescribeDirectoryConfigsOutput)

	DescribeEntitlements(*appstream.DescribeEntitlementsInput) (*appstream.DescribeEntitlementsOutput, error)
	DescribeEntitlementsWithContext(aws.Context, *appstream.DescribeEntitlementsInput, ...request.Option) (*appstream.DescribeEntitlementsOutput, error)
	DescribeEntitlementsRequest(*appstream.DescribeEntitlementsInput) (*request.Request, *appstream.DescribeEntitlementsOutput)

	DescribeFleets(*appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error)
	DescribeFleetsWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.Option) (*appstream.DescribeFleetsOutput, error)
	DescribeFleetsRequest(*appstream.DescribeFleetsInput) (*request.Request, *appstream.DescribeFleetsOutput)

	DescribeImageBuilders(*appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImageBuildersWithContext(aws.Context, *appstream.DescribeImageBuildersInput, ...request.Option) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImageBuildersRequest(*appstream.DescribeImageBuildersInput) (*request.Request, *appstream.DescribeImageBuildersOutput)

	DescribeImagePermissions(*appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImagePermissionsWithContext(aws.Context, *appstream.DescribeImagePermissionsInput, ...request.Option) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImagePermissionsRequest(*appstream.DescribeImagePermissionsInput) (*request.Request, *appstream.DescribeImagePermissionsOutput)

	DescribeImagePermissionsPages(*appstream.DescribeImagePermissionsInput, func(*appstream.DescribeImagePermissionsOutput, bool) bool) error
	DescribeImagePermissionsPagesWithContext(aws.Context, *appstream.DescribeImagePermissionsInput, func(*appstream.DescribeImagePermissionsOutput, bool) bool, ...request.Option) error

	DescribeImages(*appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error)
	DescribeImagesWithContext(aws.Context, *appstream.DescribeImagesInput, ...request.Option) (*appstream.DescribeImagesOutput, error)
	DescribeImagesRequest(*appstream.DescribeImagesInput) (*request.Request, *appstream.DescribeImagesOutput)

	DescribeImagesPages(*appstream.DescribeImagesInput, func(*appstream.DescribeImagesOutput, bool) bool) error
	DescribeImagesPagesWithContext(aws.Context, *appstream.DescribeImagesInput, func(*appstream.DescribeImagesOutput, bool) bool, ...request.Option) error

	DescribeSessions(*appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error)
	DescribeSessionsWithContext(aws.Context, *appstream.DescribeSessionsInput, ...request.Option) (*appstream.DescribeSessionsOutput, error)
	DescribeSessionsRequest(*appstream.DescribeSessionsInput) (*request.Request, *appstream.DescribeSessionsOutput)

	DescribeStacks(*appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *appstream.DescribeStacksInput, ...request.Option) (*appstream.DescribeStacksOutput, error)
	DescribeStacksRequest(*appstream.DescribeStacksInput) (*request.Request, *appstream.DescribeStacksOutput)

	DescribeUsageReportSubscriptions(*appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUsageReportSubscriptionsWithContext(aws.Context, *appstream.DescribeUsageReportSubscriptionsInput, ...request.Option) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUsageReportSubscriptionsRequest(*appstream.DescribeUsageReportSubscriptionsInput) (*request.Request, *appstream.DescribeUsageReportSubscriptionsOutput)

	DescribeUserStackAssociations(*appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUserStackAssociationsWithContext(aws.Context, *appstream.DescribeUserStackAssociationsInput, ...request.Option) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUserStackAssociationsRequest(*appstream.DescribeUserStackAssociationsInput) (*request.Request, *appstream.DescribeUserStackAssociationsOutput)

	DescribeUsers(*appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error)
	DescribeUsersWithContext(aws.Context, *appstream.DescribeUsersInput, ...request.Option) (*appstream.DescribeUsersOutput, error)
	DescribeUsersRequest(*appstream.DescribeUsersInput) (*request.Request, *appstream.DescribeUsersOutput)

	DisableUser(*appstream.DisableUserInput) (*appstream.DisableUserOutput, error)
	DisableUserWithContext(aws.Context, *appstream.DisableUserInput, ...request.Option) (*appstream.DisableUserOutput, error)
	DisableUserRequest(*appstream.DisableUserInput) (*request.Request, *appstream.DisableUserOutput)

	DisassociateAppBlockBuilderAppBlock(*appstream.DisassociateAppBlockBuilderAppBlockInput) (*appstream.DisassociateAppBlockBuilderAppBlockOutput, error)
	DisassociateAppBlockBuilderAppBlockWithContext(aws.Context, *appstream.DisassociateAppBlockBuilderAppBlockInput, ...request.Option) (*appstream.DisassociateAppBlockBuilderAppBlockOutput, error)
	DisassociateAppBlockBuilderAppBlockRequest(*appstream.DisassociateAppBlockBuilderAppBlockInput) (*request.Request, *appstream.DisassociateAppBlockBuilderAppBlockOutput)

	DisassociateApplicationFleet(*appstream.DisassociateApplicationFleetInput) (*appstream.DisassociateApplicationFleetOutput, error)
	DisassociateApplicationFleetWithContext(aws.Context, *appstream.DisassociateApplicationFleetInput, ...request.Option) (*appstream.DisassociateApplicationFleetOutput, error)
	DisassociateApplicationFleetRequest(*appstream.DisassociateApplicationFleetInput) (*request.Request, *appstream.DisassociateApplicationFleetOutput)

	DisassociateApplicationFromEntitlement(*appstream.DisassociateApplicationFromEntitlementInput) (*appstream.DisassociateApplicationFromEntitlementOutput, error)
	DisassociateApplicationFromEntitlementWithContext(aws.Context, *appstream.DisassociateApplicationFromEntitlementInput, ...request.Option) (*appstream.DisassociateApplicationFromEntitlementOutput, error)
	DisassociateApplicationFromEntitlementRequest(*appstream.DisassociateApplicationFromEntitlementInput) (*request.Request, *appstream.DisassociateApplicationFromEntitlementOutput)

	DisassociateFleet(*appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error)
	DisassociateFleetWithContext(aws.Context, *appstream.DisassociateFleetInput, ...request.Option) (*appstream.DisassociateFleetOutput, error)
	DisassociateFleetRequest(*appstream.DisassociateFleetInput) (*request.Request, *appstream.DisassociateFleetOutput)

	EnableUser(*appstream.EnableUserInput) (*appstream.EnableUserOutput, error)
	EnableUserWithContext(aws.Context, *appstream.EnableUserInput, ...request.Option) (*appstream.EnableUserOutput, error)
	EnableUserRequest(*appstream.EnableUserInput) (*request.Request, *appstream.EnableUserOutput)

	ExpireSession(*appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error)
	ExpireSessionWithContext(aws.Context, *appstream.ExpireSessionInput, ...request.Option) (*appstream.ExpireSessionOutput, error)
	ExpireSessionRequest(*appstream.ExpireSessionInput) (*request.Request, *appstream.ExpireSessionOutput)

	ListAssociatedFleets(*appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedFleetsWithContext(aws.Context, *appstream.ListAssociatedFleetsInput, ...request.Option) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedFleetsRequest(*appstream.ListAssociatedFleetsInput) (*request.Request, *appstream.ListAssociatedFleetsOutput)

	ListAssociatedStacks(*appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error)
	ListAssociatedStacksWithContext(aws.Context, *appstream.ListAssociatedStacksInput, ...request.Option) (*appstream.ListAssociatedStacksOutput, error)
	ListAssociatedStacksRequest(*appstream.ListAssociatedStacksInput) (*request.Request, *appstream.ListAssociatedStacksOutput)

	ListEntitledApplications(*appstream.ListEntitledApplicationsInput) (*appstream.ListEntitledApplicationsOutput, error)
	ListEntitledApplicationsWithContext(aws.Context, *appstream.ListEntitledApplicationsInput, ...request.Option) (*appstream.ListEntitledApplicationsOutput, error)
	ListEntitledApplicationsRequest(*appstream.ListEntitledApplicationsInput) (*request.Request, *appstream.ListEntitledApplicationsOutput)

	ListTagsForResource(*appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *appstream.ListTagsForResourceInput, ...request.Option) (*appstream.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*appstream.ListTagsForResourceInput) (*request.Request, *appstream.ListTagsForResourceOutput)

	StartAppBlockBuilder(*appstream.StartAppBlockBuilderInput) (*appstream.StartAppBlockBuilderOutput, error)
	StartAppBlockBuilderWithContext(aws.Context, *appstream.StartAppBlockBuilderInput, ...request.Option) (*appstream.StartAppBlockBuilderOutput, error)
	StartAppBlockBuilderRequest(*appstream.StartAppBlockBuilderInput) (*request.Request, *appstream.StartAppBlockBuilderOutput)

	StartFleet(*appstream.StartFleetInput) (*appstream.StartFleetOutput, error)
	StartFleetWithContext(aws.Context, *appstream.StartFleetInput, ...request.Option) (*appstream.StartFleetOutput, error)
	StartFleetRequest(*appstream.StartFleetInput) (*request.Request, *appstream.StartFleetOutput)

	StartImageBuilder(*appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error)
	StartImageBuilderWithContext(aws.Context, *appstream.StartImageBuilderInput, ...request.Option) (*appstream.StartImageBuilderOutput, error)
	StartImageBuilderRequest(*appstream.StartImageBuilderInput) (*request.Request, *appstream.StartImageBuilderOutput)

	StopAppBlockBuilder(*appstream.StopAppBlockBuilderInput) (*appstream.StopAppBlockBuilderOutput, error)
	StopAppBlockBuilderWithContext(aws.Context, *appstream.StopAppBlockBuilderInput, ...request.Option) (*appstream.StopAppBlockBuilderOutput, error)
	StopAppBlockBuilderRequest(*appstream.StopAppBlockBuilderInput) (*request.Request, *appstream.StopAppBlockBuilderOutput)

	StopFleet(*appstream.StopFleetInput) (*appstream.StopFleetOutput, error)
	StopFleetWithContext(aws.Context, *appstream.StopFleetInput, ...request.Option) (*appstream.StopFleetOutput, error)
	StopFleetRequest(*appstream.StopFleetInput) (*request.Request, *appstream.StopFleetOutput)

	StopImageBuilder(*appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error)
	StopImageBuilderWithContext(aws.Context, *appstream.StopImageBuilderInput, ...request.Option) (*appstream.StopImageBuilderOutput, error)
	StopImageBuilderRequest(*appstream.StopImageBuilderInput) (*request.Request, *appstream.StopImageBuilderOutput)

	TagResource(*appstream.TagResourceInput) (*appstream.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *appstream.TagResourceInput, ...request.Option) (*appstream.TagResourceOutput, error)
	TagResourceRequest(*appstream.TagResourceInput) (*request.Request, *appstream.TagResourceOutput)

	UntagResource(*appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *appstream.UntagResourceInput, ...request.Option) (*appstream.UntagResourceOutput, error)
	UntagResourceRequest(*appstream.UntagResourceInput) (*request.Request, *appstream.UntagResourceOutput)

	UpdateAppBlockBuilder(*appstream.UpdateAppBlockBuilderInput) (*appstream.UpdateAppBlockBuilderOutput, error)
	UpdateAppBlockBuilderWithContext(aws.Context, *appstream.UpdateAppBlockBuilderInput, ...request.Option) (*appstream.UpdateAppBlockBuilderOutput, error)
	UpdateAppBlockBuilderRequest(*appstream.UpdateAppBlockBuilderInput) (*request.Request, *appstream.UpdateAppBlockBuilderOutput)

	UpdateApplication(*appstream.UpdateApplicationInput) (*appstream.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *appstream.UpdateApplicationInput, ...request.Option) (*appstream.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*appstream.UpdateApplicationInput) (*request.Request, *appstream.UpdateApplicationOutput)

	UpdateDirectoryConfig(*appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateDirectoryConfigWithContext(aws.Context, *appstream.UpdateDirectoryConfigInput, ...request.Option) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateDirectoryConfigRequest(*appstream.UpdateDirectoryConfigInput) (*request.Request, *appstream.UpdateDirectoryConfigOutput)

	UpdateEntitlement(*appstream.UpdateEntitlementInput) (*appstream.UpdateEntitlementOutput, error)
	UpdateEntitlementWithContext(aws.Context, *appstream.UpdateEntitlementInput, ...request.Option) (*appstream.UpdateEntitlementOutput, error)
	UpdateEntitlementRequest(*appstream.UpdateEntitlementInput) (*request.Request, *appstream.UpdateEntitlementOutput)

	UpdateFleet(*appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error)
	UpdateFleetWithContext(aws.Context, *appstream.UpdateFleetInput, ...request.Option) (*appstream.UpdateFleetOutput, error)
	UpdateFleetRequest(*appstream.UpdateFleetInput) (*request.Request, *appstream.UpdateFleetOutput)

	UpdateImagePermissions(*appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error)
	UpdateImagePermissionsWithContext(aws.Context, *appstream.UpdateImagePermissionsInput, ...request.Option) (*appstream.UpdateImagePermissionsOutput, error)
	UpdateImagePermissionsRequest(*appstream.UpdateImagePermissionsInput) (*request.Request, *appstream.UpdateImagePermissionsOutput)

	UpdateStack(*appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *appstream.UpdateStackInput, ...request.Option) (*appstream.UpdateStackOutput, error)
	UpdateStackRequest(*appstream.UpdateStackInput) (*request.Request, *appstream.UpdateStackOutput)

	WaitUntilFleetStarted(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStartedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.WaiterOption) error

	WaitUntilFleetStopped(*appstream.DescribeFleetsInput) error
	WaitUntilFleetStoppedWithContext(aws.Context, *appstream.DescribeFleetsInput, ...request.WaiterOption) error
}

var _ AppStreamAPI = (*appstream.AppStream)(nil)
