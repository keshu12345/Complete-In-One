// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directoryserviceiface provides an interface to enable mocking the AWS Directory Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directoryserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

// DirectoryServiceAPI provides an interface to enable mocking the
// directoryservice.DirectoryService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Directory Service.
//	func myFunc(svc directoryserviceiface.DirectoryServiceAPI) bool {
//	    // Make svc.AcceptSharedDirectory request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := directoryservice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockDirectoryServiceClient struct {
//	    directoryserviceiface.DirectoryServiceAPI
//	}
//	func (m *mockDirectoryServiceClient) AcceptSharedDirectory(input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockDirectoryServiceClient{}
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
type DirectoryServiceAPI interface {
	AcceptSharedDirectory(*directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error)
	AcceptSharedDirectoryWithContext(aws.Context, *directoryservice.AcceptSharedDirectoryInput, ...request.Option) (*directoryservice.AcceptSharedDirectoryOutput, error)
	AcceptSharedDirectoryRequest(*directoryservice.AcceptSharedDirectoryInput) (*request.Request, *directoryservice.AcceptSharedDirectoryOutput)

	AddIpRoutes(*directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error)
	AddIpRoutesWithContext(aws.Context, *directoryservice.AddIpRoutesInput, ...request.Option) (*directoryservice.AddIpRoutesOutput, error)
	AddIpRoutesRequest(*directoryservice.AddIpRoutesInput) (*request.Request, *directoryservice.AddIpRoutesOutput)

	AddRegion(*directoryservice.AddRegionInput) (*directoryservice.AddRegionOutput, error)
	AddRegionWithContext(aws.Context, *directoryservice.AddRegionInput, ...request.Option) (*directoryservice.AddRegionOutput, error)
	AddRegionRequest(*directoryservice.AddRegionInput) (*request.Request, *directoryservice.AddRegionOutput)

	AddTagsToResource(*directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(aws.Context, *directoryservice.AddTagsToResourceInput, ...request.Option) (*directoryservice.AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*directoryservice.AddTagsToResourceInput) (*request.Request, *directoryservice.AddTagsToResourceOutput)

	CancelSchemaExtension(*directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error)
	CancelSchemaExtensionWithContext(aws.Context, *directoryservice.CancelSchemaExtensionInput, ...request.Option) (*directoryservice.CancelSchemaExtensionOutput, error)
	CancelSchemaExtensionRequest(*directoryservice.CancelSchemaExtensionInput) (*request.Request, *directoryservice.CancelSchemaExtensionOutput)

	ConnectDirectory(*directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)
	ConnectDirectoryWithContext(aws.Context, *directoryservice.ConnectDirectoryInput, ...request.Option) (*directoryservice.ConnectDirectoryOutput, error)
	ConnectDirectoryRequest(*directoryservice.ConnectDirectoryInput) (*request.Request, *directoryservice.ConnectDirectoryOutput)

	CreateAlias(*directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)
	CreateAliasWithContext(aws.Context, *directoryservice.CreateAliasInput, ...request.Option) (*directoryservice.CreateAliasOutput, error)
	CreateAliasRequest(*directoryservice.CreateAliasInput) (*request.Request, *directoryservice.CreateAliasOutput)

	CreateComputer(*directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)
	CreateComputerWithContext(aws.Context, *directoryservice.CreateComputerInput, ...request.Option) (*directoryservice.CreateComputerOutput, error)
	CreateComputerRequest(*directoryservice.CreateComputerInput) (*request.Request, *directoryservice.CreateComputerOutput)

	CreateConditionalForwarder(*directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateConditionalForwarderWithContext(aws.Context, *directoryservice.CreateConditionalForwarderInput, ...request.Option) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateConditionalForwarderRequest(*directoryservice.CreateConditionalForwarderInput) (*request.Request, *directoryservice.CreateConditionalForwarderOutput)

	CreateDirectory(*directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)
	CreateDirectoryWithContext(aws.Context, *directoryservice.CreateDirectoryInput, ...request.Option) (*directoryservice.CreateDirectoryOutput, error)
	CreateDirectoryRequest(*directoryservice.CreateDirectoryInput) (*request.Request, *directoryservice.CreateDirectoryOutput)

	CreateLogSubscription(*directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error)
	CreateLogSubscriptionWithContext(aws.Context, *directoryservice.CreateLogSubscriptionInput, ...request.Option) (*directoryservice.CreateLogSubscriptionOutput, error)
	CreateLogSubscriptionRequest(*directoryservice.CreateLogSubscriptionInput) (*request.Request, *directoryservice.CreateLogSubscriptionOutput)

	CreateMicrosoftAD(*directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateMicrosoftADWithContext(aws.Context, *directoryservice.CreateMicrosoftADInput, ...request.Option) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateMicrosoftADRequest(*directoryservice.CreateMicrosoftADInput) (*request.Request, *directoryservice.CreateMicrosoftADOutput)

	CreateSnapshot(*directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)
	CreateSnapshotWithContext(aws.Context, *directoryservice.CreateSnapshotInput, ...request.Option) (*directoryservice.CreateSnapshotOutput, error)
	CreateSnapshotRequest(*directoryservice.CreateSnapshotInput) (*request.Request, *directoryservice.CreateSnapshotOutput)

	CreateTrust(*directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error)
	CreateTrustWithContext(aws.Context, *directoryservice.CreateTrustInput, ...request.Option) (*directoryservice.CreateTrustOutput, error)
	CreateTrustRequest(*directoryservice.CreateTrustInput) (*request.Request, *directoryservice.CreateTrustOutput)

	DeleteConditionalForwarder(*directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteConditionalForwarderWithContext(aws.Context, *directoryservice.DeleteConditionalForwarderInput, ...request.Option) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteConditionalForwarderRequest(*directoryservice.DeleteConditionalForwarderInput) (*request.Request, *directoryservice.DeleteConditionalForwarderOutput)

	DeleteDirectory(*directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteDirectoryWithContext(aws.Context, *directoryservice.DeleteDirectoryInput, ...request.Option) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteDirectoryRequest(*directoryservice.DeleteDirectoryInput) (*request.Request, *directoryservice.DeleteDirectoryOutput)

	DeleteLogSubscription(*directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error)
	DeleteLogSubscriptionWithContext(aws.Context, *directoryservice.DeleteLogSubscriptionInput, ...request.Option) (*directoryservice.DeleteLogSubscriptionOutput, error)
	DeleteLogSubscriptionRequest(*directoryservice.DeleteLogSubscriptionInput) (*request.Request, *directoryservice.DeleteLogSubscriptionOutput)

	DeleteSnapshot(*directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteSnapshotWithContext(aws.Context, *directoryservice.DeleteSnapshotInput, ...request.Option) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteSnapshotRequest(*directoryservice.DeleteSnapshotInput) (*request.Request, *directoryservice.DeleteSnapshotOutput)

	DeleteTrust(*directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error)
	DeleteTrustWithContext(aws.Context, *directoryservice.DeleteTrustInput, ...request.Option) (*directoryservice.DeleteTrustOutput, error)
	DeleteTrustRequest(*directoryservice.DeleteTrustInput) (*request.Request, *directoryservice.DeleteTrustOutput)

	DeregisterCertificate(*directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error)
	DeregisterCertificateWithContext(aws.Context, *directoryservice.DeregisterCertificateInput, ...request.Option) (*directoryservice.DeregisterCertificateOutput, error)
	DeregisterCertificateRequest(*directoryservice.DeregisterCertificateInput) (*request.Request, *directoryservice.DeregisterCertificateOutput)

	DeregisterEventTopic(*directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error)
	DeregisterEventTopicWithContext(aws.Context, *directoryservice.DeregisterEventTopicInput, ...request.Option) (*directoryservice.DeregisterEventTopicOutput, error)
	DeregisterEventTopicRequest(*directoryservice.DeregisterEventTopicInput) (*request.Request, *directoryservice.DeregisterEventTopicOutput)

	DescribeCertificate(*directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error)
	DescribeCertificateWithContext(aws.Context, *directoryservice.DescribeCertificateInput, ...request.Option) (*directoryservice.DescribeCertificateOutput, error)
	DescribeCertificateRequest(*directoryservice.DescribeCertificateInput) (*request.Request, *directoryservice.DescribeCertificateOutput)

	DescribeClientAuthenticationSettings(*directoryservice.DescribeClientAuthenticationSettingsInput) (*directoryservice.DescribeClientAuthenticationSettingsOutput, error)
	DescribeClientAuthenticationSettingsWithContext(aws.Context, *directoryservice.DescribeClientAuthenticationSettingsInput, ...request.Option) (*directoryservice.DescribeClientAuthenticationSettingsOutput, error)
	DescribeClientAuthenticationSettingsRequest(*directoryservice.DescribeClientAuthenticationSettingsInput) (*request.Request, *directoryservice.DescribeClientAuthenticationSettingsOutput)

	DescribeClientAuthenticationSettingsPages(*directoryservice.DescribeClientAuthenticationSettingsInput, func(*directoryservice.DescribeClientAuthenticationSettingsOutput, bool) bool) error
	DescribeClientAuthenticationSettingsPagesWithContext(aws.Context, *directoryservice.DescribeClientAuthenticationSettingsInput, func(*directoryservice.DescribeClientAuthenticationSettingsOutput, bool) bool, ...request.Option) error

	DescribeConditionalForwarders(*directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeConditionalForwardersWithContext(aws.Context, *directoryservice.DescribeConditionalForwardersInput, ...request.Option) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeConditionalForwardersRequest(*directoryservice.DescribeConditionalForwardersInput) (*request.Request, *directoryservice.DescribeConditionalForwardersOutput)

	DescribeDirectories(*directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDirectoriesWithContext(aws.Context, *directoryservice.DescribeDirectoriesInput, ...request.Option) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDirectoriesRequest(*directoryservice.DescribeDirectoriesInput) (*request.Request, *directoryservice.DescribeDirectoriesOutput)

	DescribeDirectoriesPages(*directoryservice.DescribeDirectoriesInput, func(*directoryservice.DescribeDirectoriesOutput, bool) bool) error
	DescribeDirectoriesPagesWithContext(aws.Context, *directoryservice.DescribeDirectoriesInput, func(*directoryservice.DescribeDirectoriesOutput, bool) bool, ...request.Option) error

	DescribeDomainControllers(*directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeDomainControllersWithContext(aws.Context, *directoryservice.DescribeDomainControllersInput, ...request.Option) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeDomainControllersRequest(*directoryservice.DescribeDomainControllersInput) (*request.Request, *directoryservice.DescribeDomainControllersOutput)

	DescribeDomainControllersPages(*directoryservice.DescribeDomainControllersInput, func(*directoryservice.DescribeDomainControllersOutput, bool) bool) error
	DescribeDomainControllersPagesWithContext(aws.Context, *directoryservice.DescribeDomainControllersInput, func(*directoryservice.DescribeDomainControllersOutput, bool) bool, ...request.Option) error

	DescribeEventTopics(*directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeEventTopicsWithContext(aws.Context, *directoryservice.DescribeEventTopicsInput, ...request.Option) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeEventTopicsRequest(*directoryservice.DescribeEventTopicsInput) (*request.Request, *directoryservice.DescribeEventTopicsOutput)

	DescribeLDAPSSettings(*directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error)
	DescribeLDAPSSettingsWithContext(aws.Context, *directoryservice.DescribeLDAPSSettingsInput, ...request.Option) (*directoryservice.DescribeLDAPSSettingsOutput, error)
	DescribeLDAPSSettingsRequest(*directoryservice.DescribeLDAPSSettingsInput) (*request.Request, *directoryservice.DescribeLDAPSSettingsOutput)

	DescribeLDAPSSettingsPages(*directoryservice.DescribeLDAPSSettingsInput, func(*directoryservice.DescribeLDAPSSettingsOutput, bool) bool) error
	DescribeLDAPSSettingsPagesWithContext(aws.Context, *directoryservice.DescribeLDAPSSettingsInput, func(*directoryservice.DescribeLDAPSSettingsOutput, bool) bool, ...request.Option) error

	DescribeRegions(*directoryservice.DescribeRegionsInput) (*directoryservice.DescribeRegionsOutput, error)
	DescribeRegionsWithContext(aws.Context, *directoryservice.DescribeRegionsInput, ...request.Option) (*directoryservice.DescribeRegionsOutput, error)
	DescribeRegionsRequest(*directoryservice.DescribeRegionsInput) (*request.Request, *directoryservice.DescribeRegionsOutput)

	DescribeRegionsPages(*directoryservice.DescribeRegionsInput, func(*directoryservice.DescribeRegionsOutput, bool) bool) error
	DescribeRegionsPagesWithContext(aws.Context, *directoryservice.DescribeRegionsInput, func(*directoryservice.DescribeRegionsOutput, bool) bool, ...request.Option) error

	DescribeSettings(*directoryservice.DescribeSettingsInput) (*directoryservice.DescribeSettingsOutput, error)
	DescribeSettingsWithContext(aws.Context, *directoryservice.DescribeSettingsInput, ...request.Option) (*directoryservice.DescribeSettingsOutput, error)
	DescribeSettingsRequest(*directoryservice.DescribeSettingsInput) (*request.Request, *directoryservice.DescribeSettingsOutput)

	DescribeSharedDirectories(*directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error)
	DescribeSharedDirectoriesWithContext(aws.Context, *directoryservice.DescribeSharedDirectoriesInput, ...request.Option) (*directoryservice.DescribeSharedDirectoriesOutput, error)
	DescribeSharedDirectoriesRequest(*directoryservice.DescribeSharedDirectoriesInput) (*request.Request, *directoryservice.DescribeSharedDirectoriesOutput)

	DescribeSharedDirectoriesPages(*directoryservice.DescribeSharedDirectoriesInput, func(*directoryservice.DescribeSharedDirectoriesOutput, bool) bool) error
	DescribeSharedDirectoriesPagesWithContext(aws.Context, *directoryservice.DescribeSharedDirectoriesInput, func(*directoryservice.DescribeSharedDirectoriesOutput, bool) bool, ...request.Option) error

	DescribeSnapshots(*directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeSnapshotsWithContext(aws.Context, *directoryservice.DescribeSnapshotsInput, ...request.Option) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequest(*directoryservice.DescribeSnapshotsInput) (*request.Request, *directoryservice.DescribeSnapshotsOutput)

	DescribeSnapshotsPages(*directoryservice.DescribeSnapshotsInput, func(*directoryservice.DescribeSnapshotsOutput, bool) bool) error
	DescribeSnapshotsPagesWithContext(aws.Context, *directoryservice.DescribeSnapshotsInput, func(*directoryservice.DescribeSnapshotsOutput, bool) bool, ...request.Option) error

	DescribeTrusts(*directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error)
	DescribeTrustsWithContext(aws.Context, *directoryservice.DescribeTrustsInput, ...request.Option) (*directoryservice.DescribeTrustsOutput, error)
	DescribeTrustsRequest(*directoryservice.DescribeTrustsInput) (*request.Request, *directoryservice.DescribeTrustsOutput)

	DescribeTrustsPages(*directoryservice.DescribeTrustsInput, func(*directoryservice.DescribeTrustsOutput, bool) bool) error
	DescribeTrustsPagesWithContext(aws.Context, *directoryservice.DescribeTrustsInput, func(*directoryservice.DescribeTrustsOutput, bool) bool, ...request.Option) error

	DescribeUpdateDirectory(*directoryservice.DescribeUpdateDirectoryInput) (*directoryservice.DescribeUpdateDirectoryOutput, error)
	DescribeUpdateDirectoryWithContext(aws.Context, *directoryservice.DescribeUpdateDirectoryInput, ...request.Option) (*directoryservice.DescribeUpdateDirectoryOutput, error)
	DescribeUpdateDirectoryRequest(*directoryservice.DescribeUpdateDirectoryInput) (*request.Request, *directoryservice.DescribeUpdateDirectoryOutput)

	DescribeUpdateDirectoryPages(*directoryservice.DescribeUpdateDirectoryInput, func(*directoryservice.DescribeUpdateDirectoryOutput, bool) bool) error
	DescribeUpdateDirectoryPagesWithContext(aws.Context, *directoryservice.DescribeUpdateDirectoryInput, func(*directoryservice.DescribeUpdateDirectoryOutput, bool) bool, ...request.Option) error

	DisableClientAuthentication(*directoryservice.DisableClientAuthenticationInput) (*directoryservice.DisableClientAuthenticationOutput, error)
	DisableClientAuthenticationWithContext(aws.Context, *directoryservice.DisableClientAuthenticationInput, ...request.Option) (*directoryservice.DisableClientAuthenticationOutput, error)
	DisableClientAuthenticationRequest(*directoryservice.DisableClientAuthenticationInput) (*request.Request, *directoryservice.DisableClientAuthenticationOutput)

	DisableLDAPS(*directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error)
	DisableLDAPSWithContext(aws.Context, *directoryservice.DisableLDAPSInput, ...request.Option) (*directoryservice.DisableLDAPSOutput, error)
	DisableLDAPSRequest(*directoryservice.DisableLDAPSInput) (*request.Request, *directoryservice.DisableLDAPSOutput)

	DisableRadius(*directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)
	DisableRadiusWithContext(aws.Context, *directoryservice.DisableRadiusInput, ...request.Option) (*directoryservice.DisableRadiusOutput, error)
	DisableRadiusRequest(*directoryservice.DisableRadiusInput) (*request.Request, *directoryservice.DisableRadiusOutput)

	DisableSso(*directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)
	DisableSsoWithContext(aws.Context, *directoryservice.DisableSsoInput, ...request.Option) (*directoryservice.DisableSsoOutput, error)
	DisableSsoRequest(*directoryservice.DisableSsoInput) (*request.Request, *directoryservice.DisableSsoOutput)

	EnableClientAuthentication(*directoryservice.EnableClientAuthenticationInput) (*directoryservice.EnableClientAuthenticationOutput, error)
	EnableClientAuthenticationWithContext(aws.Context, *directoryservice.EnableClientAuthenticationInput, ...request.Option) (*directoryservice.EnableClientAuthenticationOutput, error)
	EnableClientAuthenticationRequest(*directoryservice.EnableClientAuthenticationInput) (*request.Request, *directoryservice.EnableClientAuthenticationOutput)

	EnableLDAPS(*directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error)
	EnableLDAPSWithContext(aws.Context, *directoryservice.EnableLDAPSInput, ...request.Option) (*directoryservice.EnableLDAPSOutput, error)
	EnableLDAPSRequest(*directoryservice.EnableLDAPSInput) (*request.Request, *directoryservice.EnableLDAPSOutput)

	EnableRadius(*directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)
	EnableRadiusWithContext(aws.Context, *directoryservice.EnableRadiusInput, ...request.Option) (*directoryservice.EnableRadiusOutput, error)
	EnableRadiusRequest(*directoryservice.EnableRadiusInput) (*request.Request, *directoryservice.EnableRadiusOutput)

	EnableSso(*directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)
	EnableSsoWithContext(aws.Context, *directoryservice.EnableSsoInput, ...request.Option) (*directoryservice.EnableSsoOutput, error)
	EnableSsoRequest(*directoryservice.EnableSsoInput) (*request.Request, *directoryservice.EnableSsoOutput)

	GetDirectoryLimits(*directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetDirectoryLimitsWithContext(aws.Context, *directoryservice.GetDirectoryLimitsInput, ...request.Option) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetDirectoryLimitsRequest(*directoryservice.GetDirectoryLimitsInput) (*request.Request, *directoryservice.GetDirectoryLimitsOutput)

	GetSnapshotLimits(*directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)
	GetSnapshotLimitsWithContext(aws.Context, *directoryservice.GetSnapshotLimitsInput, ...request.Option) (*directoryservice.GetSnapshotLimitsOutput, error)
	GetSnapshotLimitsRequest(*directoryservice.GetSnapshotLimitsInput) (*request.Request, *directoryservice.GetSnapshotLimitsOutput)

	ListCertificates(*directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error)
	ListCertificatesWithContext(aws.Context, *directoryservice.ListCertificatesInput, ...request.Option) (*directoryservice.ListCertificatesOutput, error)
	ListCertificatesRequest(*directoryservice.ListCertificatesInput) (*request.Request, *directoryservice.ListCertificatesOutput)

	ListCertificatesPages(*directoryservice.ListCertificatesInput, func(*directoryservice.ListCertificatesOutput, bool) bool) error
	ListCertificatesPagesWithContext(aws.Context, *directoryservice.ListCertificatesInput, func(*directoryservice.ListCertificatesOutput, bool) bool, ...request.Option) error

	ListIpRoutes(*directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error)
	ListIpRoutesWithContext(aws.Context, *directoryservice.ListIpRoutesInput, ...request.Option) (*directoryservice.ListIpRoutesOutput, error)
	ListIpRoutesRequest(*directoryservice.ListIpRoutesInput) (*request.Request, *directoryservice.ListIpRoutesOutput)

	ListIpRoutesPages(*directoryservice.ListIpRoutesInput, func(*directoryservice.ListIpRoutesOutput, bool) bool) error
	ListIpRoutesPagesWithContext(aws.Context, *directoryservice.ListIpRoutesInput, func(*directoryservice.ListIpRoutesOutput, bool) bool, ...request.Option) error

	ListLogSubscriptions(*directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error)
	ListLogSubscriptionsWithContext(aws.Context, *directoryservice.ListLogSubscriptionsInput, ...request.Option) (*directoryservice.ListLogSubscriptionsOutput, error)
	ListLogSubscriptionsRequest(*directoryservice.ListLogSubscriptionsInput) (*request.Request, *directoryservice.ListLogSubscriptionsOutput)

	ListLogSubscriptionsPages(*directoryservice.ListLogSubscriptionsInput, func(*directoryservice.ListLogSubscriptionsOutput, bool) bool) error
	ListLogSubscriptionsPagesWithContext(aws.Context, *directoryservice.ListLogSubscriptionsInput, func(*directoryservice.ListLogSubscriptionsOutput, bool) bool, ...request.Option) error

	ListSchemaExtensions(*directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListSchemaExtensionsWithContext(aws.Context, *directoryservice.ListSchemaExtensionsInput, ...request.Option) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListSchemaExtensionsRequest(*directoryservice.ListSchemaExtensionsInput) (*request.Request, *directoryservice.ListSchemaExtensionsOutput)

	ListSchemaExtensionsPages(*directoryservice.ListSchemaExtensionsInput, func(*directoryservice.ListSchemaExtensionsOutput, bool) bool) error
	ListSchemaExtensionsPagesWithContext(aws.Context, *directoryservice.ListSchemaExtensionsInput, func(*directoryservice.ListSchemaExtensionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *directoryservice.ListTagsForResourceInput, ...request.Option) (*directoryservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*directoryservice.ListTagsForResourceInput) (*request.Request, *directoryservice.ListTagsForResourceOutput)

	ListTagsForResourcePages(*directoryservice.ListTagsForResourceInput, func(*directoryservice.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *directoryservice.ListTagsForResourceInput, func(*directoryservice.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	RegisterCertificate(*directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error)
	RegisterCertificateWithContext(aws.Context, *directoryservice.RegisterCertificateInput, ...request.Option) (*directoryservice.RegisterCertificateOutput, error)
	RegisterCertificateRequest(*directoryservice.RegisterCertificateInput) (*request.Request, *directoryservice.RegisterCertificateOutput)

	RegisterEventTopic(*directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error)
	RegisterEventTopicWithContext(aws.Context, *directoryservice.RegisterEventTopicInput, ...request.Option) (*directoryservice.RegisterEventTopicOutput, error)
	RegisterEventTopicRequest(*directoryservice.RegisterEventTopicInput) (*request.Request, *directoryservice.RegisterEventTopicOutput)

	RejectSharedDirectory(*directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error)
	RejectSharedDirectoryWithContext(aws.Context, *directoryservice.RejectSharedDirectoryInput, ...request.Option) (*directoryservice.RejectSharedDirectoryOutput, error)
	RejectSharedDirectoryRequest(*directoryservice.RejectSharedDirectoryInput) (*request.Request, *directoryservice.RejectSharedDirectoryOutput)

	RemoveIpRoutes(*directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveIpRoutesWithContext(aws.Context, *directoryservice.RemoveIpRoutesInput, ...request.Option) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveIpRoutesRequest(*directoryservice.RemoveIpRoutesInput) (*request.Request, *directoryservice.RemoveIpRoutesOutput)

	RemoveRegion(*directoryservice.RemoveRegionInput) (*directoryservice.RemoveRegionOutput, error)
	RemoveRegionWithContext(aws.Context, *directoryservice.RemoveRegionInput, ...request.Option) (*directoryservice.RemoveRegionOutput, error)
	RemoveRegionRequest(*directoryservice.RemoveRegionInput) (*request.Request, *directoryservice.RemoveRegionOutput)

	RemoveTagsFromResource(*directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(aws.Context, *directoryservice.RemoveTagsFromResourceInput, ...request.Option) (*directoryservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*directoryservice.RemoveTagsFromResourceInput) (*request.Request, *directoryservice.RemoveTagsFromResourceOutput)

	ResetUserPassword(*directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error)
	ResetUserPasswordWithContext(aws.Context, *directoryservice.ResetUserPasswordInput, ...request.Option) (*directoryservice.ResetUserPasswordOutput, error)
	ResetUserPasswordRequest(*directoryservice.ResetUserPasswordInput) (*request.Request, *directoryservice.ResetUserPasswordOutput)

	RestoreFromSnapshot(*directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)
	RestoreFromSnapshotWithContext(aws.Context, *directoryservice.RestoreFromSnapshotInput, ...request.Option) (*directoryservice.RestoreFromSnapshotOutput, error)
	RestoreFromSnapshotRequest(*directoryservice.RestoreFromSnapshotInput) (*request.Request, *directoryservice.RestoreFromSnapshotOutput)

	ShareDirectory(*directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error)
	ShareDirectoryWithContext(aws.Context, *directoryservice.ShareDirectoryInput, ...request.Option) (*directoryservice.ShareDirectoryOutput, error)
	ShareDirectoryRequest(*directoryservice.ShareDirectoryInput) (*request.Request, *directoryservice.ShareDirectoryOutput)

	StartSchemaExtension(*directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error)
	StartSchemaExtensionWithContext(aws.Context, *directoryservice.StartSchemaExtensionInput, ...request.Option) (*directoryservice.StartSchemaExtensionOutput, error)
	StartSchemaExtensionRequest(*directoryservice.StartSchemaExtensionInput) (*request.Request, *directoryservice.StartSchemaExtensionOutput)

	UnshareDirectory(*directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error)
	UnshareDirectoryWithContext(aws.Context, *directoryservice.UnshareDirectoryInput, ...request.Option) (*directoryservice.UnshareDirectoryOutput, error)
	UnshareDirectoryRequest(*directoryservice.UnshareDirectoryInput) (*request.Request, *directoryservice.UnshareDirectoryOutput)

	UpdateConditionalForwarder(*directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateConditionalForwarderWithContext(aws.Context, *directoryservice.UpdateConditionalForwarderInput, ...request.Option) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateConditionalForwarderRequest(*directoryservice.UpdateConditionalForwarderInput) (*request.Request, *directoryservice.UpdateConditionalForwarderOutput)

	UpdateDirectorySetup(*directoryservice.UpdateDirectorySetupInput) (*directoryservice.UpdateDirectorySetupOutput, error)
	UpdateDirectorySetupWithContext(aws.Context, *directoryservice.UpdateDirectorySetupInput, ...request.Option) (*directoryservice.UpdateDirectorySetupOutput, error)
	UpdateDirectorySetupRequest(*directoryservice.UpdateDirectorySetupInput) (*request.Request, *directoryservice.UpdateDirectorySetupOutput)

	UpdateNumberOfDomainControllers(*directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateNumberOfDomainControllersWithContext(aws.Context, *directoryservice.UpdateNumberOfDomainControllersInput, ...request.Option) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateNumberOfDomainControllersRequest(*directoryservice.UpdateNumberOfDomainControllersInput) (*request.Request, *directoryservice.UpdateNumberOfDomainControllersOutput)

	UpdateRadius(*directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
	UpdateRadiusWithContext(aws.Context, *directoryservice.UpdateRadiusInput, ...request.Option) (*directoryservice.UpdateRadiusOutput, error)
	UpdateRadiusRequest(*directoryservice.UpdateRadiusInput) (*request.Request, *directoryservice.UpdateRadiusOutput)

	UpdateSettings(*directoryservice.UpdateSettingsInput) (*directoryservice.UpdateSettingsOutput, error)
	UpdateSettingsWithContext(aws.Context, *directoryservice.UpdateSettingsInput, ...request.Option) (*directoryservice.UpdateSettingsOutput, error)
	UpdateSettingsRequest(*directoryservice.UpdateSettingsInput) (*request.Request, *directoryservice.UpdateSettingsOutput)

	UpdateTrust(*directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error)
	UpdateTrustWithContext(aws.Context, *directoryservice.UpdateTrustInput, ...request.Option) (*directoryservice.UpdateTrustOutput, error)
	UpdateTrustRequest(*directoryservice.UpdateTrustInput) (*request.Request, *directoryservice.UpdateTrustOutput)

	VerifyTrust(*directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error)
	VerifyTrustWithContext(aws.Context, *directoryservice.VerifyTrustInput, ...request.Option) (*directoryservice.VerifyTrustOutput, error)
	VerifyTrustRequest(*directoryservice.VerifyTrustInput) (*request.Request, *directoryservice.VerifyTrustOutput)
}

var _ DirectoryServiceAPI = (*directoryservice.DirectoryService)(nil)
