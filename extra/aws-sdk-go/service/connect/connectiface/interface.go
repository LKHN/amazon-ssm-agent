// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectiface provides an interface to enable mocking the Amazon Connect Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/connect"
)

// ConnectAPI provides an interface to enable mocking the
// connect.Connect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect Service.
//    func myFunc(svc connectiface.ConnectAPI) bool {
//        // Make svc.AssociateRoutingProfileQueues request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := connect.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConnectClient struct {
//        connectiface.ConnectAPI
//    }
//    func (m *mockConnectClient) AssociateRoutingProfileQueues(input *connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConnectClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ConnectAPI interface {
	AssociateRoutingProfileQueues(*connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error)
	AssociateRoutingProfileQueuesWithContext(aws.Context, *connect.AssociateRoutingProfileQueuesInput, ...request.Option) (*connect.AssociateRoutingProfileQueuesOutput, error)
	AssociateRoutingProfileQueuesRequest(*connect.AssociateRoutingProfileQueuesInput) (*request.Request, *connect.AssociateRoutingProfileQueuesOutput)

	CreateContactFlow(*connect.CreateContactFlowInput) (*connect.CreateContactFlowOutput, error)
	CreateContactFlowWithContext(aws.Context, *connect.CreateContactFlowInput, ...request.Option) (*connect.CreateContactFlowOutput, error)
	CreateContactFlowRequest(*connect.CreateContactFlowInput) (*request.Request, *connect.CreateContactFlowOutput)

	CreateRoutingProfile(*connect.CreateRoutingProfileInput) (*connect.CreateRoutingProfileOutput, error)
	CreateRoutingProfileWithContext(aws.Context, *connect.CreateRoutingProfileInput, ...request.Option) (*connect.CreateRoutingProfileOutput, error)
	CreateRoutingProfileRequest(*connect.CreateRoutingProfileInput) (*request.Request, *connect.CreateRoutingProfileOutput)

	CreateUser(*connect.CreateUserInput) (*connect.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *connect.CreateUserInput, ...request.Option) (*connect.CreateUserOutput, error)
	CreateUserRequest(*connect.CreateUserInput) (*request.Request, *connect.CreateUserOutput)

	DeleteUser(*connect.DeleteUserInput) (*connect.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *connect.DeleteUserInput, ...request.Option) (*connect.DeleteUserOutput, error)
	DeleteUserRequest(*connect.DeleteUserInput) (*request.Request, *connect.DeleteUserOutput)

	DescribeContactFlow(*connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error)
	DescribeContactFlowWithContext(aws.Context, *connect.DescribeContactFlowInput, ...request.Option) (*connect.DescribeContactFlowOutput, error)
	DescribeContactFlowRequest(*connect.DescribeContactFlowInput) (*request.Request, *connect.DescribeContactFlowOutput)

	DescribeRoutingProfile(*connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error)
	DescribeRoutingProfileWithContext(aws.Context, *connect.DescribeRoutingProfileInput, ...request.Option) (*connect.DescribeRoutingProfileOutput, error)
	DescribeRoutingProfileRequest(*connect.DescribeRoutingProfileInput) (*request.Request, *connect.DescribeRoutingProfileOutput)

	DescribeUser(*connect.DescribeUserInput) (*connect.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *connect.DescribeUserInput, ...request.Option) (*connect.DescribeUserOutput, error)
	DescribeUserRequest(*connect.DescribeUserInput) (*request.Request, *connect.DescribeUserOutput)

	DescribeUserHierarchyGroup(*connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupWithContext(aws.Context, *connect.DescribeUserHierarchyGroupInput, ...request.Option) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupRequest(*connect.DescribeUserHierarchyGroupInput) (*request.Request, *connect.DescribeUserHierarchyGroupOutput)

	DescribeUserHierarchyStructure(*connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureWithContext(aws.Context, *connect.DescribeUserHierarchyStructureInput, ...request.Option) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureRequest(*connect.DescribeUserHierarchyStructureInput) (*request.Request, *connect.DescribeUserHierarchyStructureOutput)

	DisassociateRoutingProfileQueues(*connect.DisassociateRoutingProfileQueuesInput) (*connect.DisassociateRoutingProfileQueuesOutput, error)
	DisassociateRoutingProfileQueuesWithContext(aws.Context, *connect.DisassociateRoutingProfileQueuesInput, ...request.Option) (*connect.DisassociateRoutingProfileQueuesOutput, error)
	DisassociateRoutingProfileQueuesRequest(*connect.DisassociateRoutingProfileQueuesInput) (*request.Request, *connect.DisassociateRoutingProfileQueuesOutput)

	GetContactAttributes(*connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesWithContext(aws.Context, *connect.GetContactAttributesInput, ...request.Option) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesRequest(*connect.GetContactAttributesInput) (*request.Request, *connect.GetContactAttributesOutput)

	GetCurrentMetricData(*connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataWithContext(aws.Context, *connect.GetCurrentMetricDataInput, ...request.Option) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataRequest(*connect.GetCurrentMetricDataInput) (*request.Request, *connect.GetCurrentMetricDataOutput)

	GetCurrentMetricDataPages(*connect.GetCurrentMetricDataInput, func(*connect.GetCurrentMetricDataOutput, bool) bool) error
	GetCurrentMetricDataPagesWithContext(aws.Context, *connect.GetCurrentMetricDataInput, func(*connect.GetCurrentMetricDataOutput, bool) bool, ...request.Option) error

	GetFederationToken(*connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenWithContext(aws.Context, *connect.GetFederationTokenInput, ...request.Option) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenRequest(*connect.GetFederationTokenInput) (*request.Request, *connect.GetFederationTokenOutput)

	GetMetricData(*connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error)
	GetMetricDataWithContext(aws.Context, *connect.GetMetricDataInput, ...request.Option) (*connect.GetMetricDataOutput, error)
	GetMetricDataRequest(*connect.GetMetricDataInput) (*request.Request, *connect.GetMetricDataOutput)

	GetMetricDataPages(*connect.GetMetricDataInput, func(*connect.GetMetricDataOutput, bool) bool) error
	GetMetricDataPagesWithContext(aws.Context, *connect.GetMetricDataInput, func(*connect.GetMetricDataOutput, bool) bool, ...request.Option) error

	ListContactFlows(*connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error)
	ListContactFlowsWithContext(aws.Context, *connect.ListContactFlowsInput, ...request.Option) (*connect.ListContactFlowsOutput, error)
	ListContactFlowsRequest(*connect.ListContactFlowsInput) (*request.Request, *connect.ListContactFlowsOutput)

	ListContactFlowsPages(*connect.ListContactFlowsInput, func(*connect.ListContactFlowsOutput, bool) bool) error
	ListContactFlowsPagesWithContext(aws.Context, *connect.ListContactFlowsInput, func(*connect.ListContactFlowsOutput, bool) bool, ...request.Option) error

	ListHoursOfOperations(*connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error)
	ListHoursOfOperationsWithContext(aws.Context, *connect.ListHoursOfOperationsInput, ...request.Option) (*connect.ListHoursOfOperationsOutput, error)
	ListHoursOfOperationsRequest(*connect.ListHoursOfOperationsInput) (*request.Request, *connect.ListHoursOfOperationsOutput)

	ListHoursOfOperationsPages(*connect.ListHoursOfOperationsInput, func(*connect.ListHoursOfOperationsOutput, bool) bool) error
	ListHoursOfOperationsPagesWithContext(aws.Context, *connect.ListHoursOfOperationsInput, func(*connect.ListHoursOfOperationsOutput, bool) bool, ...request.Option) error

	ListPhoneNumbers(*connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error)
	ListPhoneNumbersWithContext(aws.Context, *connect.ListPhoneNumbersInput, ...request.Option) (*connect.ListPhoneNumbersOutput, error)
	ListPhoneNumbersRequest(*connect.ListPhoneNumbersInput) (*request.Request, *connect.ListPhoneNumbersOutput)

	ListPhoneNumbersPages(*connect.ListPhoneNumbersInput, func(*connect.ListPhoneNumbersOutput, bool) bool) error
	ListPhoneNumbersPagesWithContext(aws.Context, *connect.ListPhoneNumbersInput, func(*connect.ListPhoneNumbersOutput, bool) bool, ...request.Option) error

	ListPrompts(*connect.ListPromptsInput) (*connect.ListPromptsOutput, error)
	ListPromptsWithContext(aws.Context, *connect.ListPromptsInput, ...request.Option) (*connect.ListPromptsOutput, error)
	ListPromptsRequest(*connect.ListPromptsInput) (*request.Request, *connect.ListPromptsOutput)

	ListPromptsPages(*connect.ListPromptsInput, func(*connect.ListPromptsOutput, bool) bool) error
	ListPromptsPagesWithContext(aws.Context, *connect.ListPromptsInput, func(*connect.ListPromptsOutput, bool) bool, ...request.Option) error

	ListQueues(*connect.ListQueuesInput) (*connect.ListQueuesOutput, error)
	ListQueuesWithContext(aws.Context, *connect.ListQueuesInput, ...request.Option) (*connect.ListQueuesOutput, error)
	ListQueuesRequest(*connect.ListQueuesInput) (*request.Request, *connect.ListQueuesOutput)

	ListQueuesPages(*connect.ListQueuesInput, func(*connect.ListQueuesOutput, bool) bool) error
	ListQueuesPagesWithContext(aws.Context, *connect.ListQueuesInput, func(*connect.ListQueuesOutput, bool) bool, ...request.Option) error

	ListRoutingProfileQueues(*connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error)
	ListRoutingProfileQueuesWithContext(aws.Context, *connect.ListRoutingProfileQueuesInput, ...request.Option) (*connect.ListRoutingProfileQueuesOutput, error)
	ListRoutingProfileQueuesRequest(*connect.ListRoutingProfileQueuesInput) (*request.Request, *connect.ListRoutingProfileQueuesOutput)

	ListRoutingProfileQueuesPages(*connect.ListRoutingProfileQueuesInput, func(*connect.ListRoutingProfileQueuesOutput, bool) bool) error
	ListRoutingProfileQueuesPagesWithContext(aws.Context, *connect.ListRoutingProfileQueuesInput, func(*connect.ListRoutingProfileQueuesOutput, bool) bool, ...request.Option) error

	ListRoutingProfiles(*connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesWithContext(aws.Context, *connect.ListRoutingProfilesInput, ...request.Option) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesRequest(*connect.ListRoutingProfilesInput) (*request.Request, *connect.ListRoutingProfilesOutput)

	ListRoutingProfilesPages(*connect.ListRoutingProfilesInput, func(*connect.ListRoutingProfilesOutput, bool) bool) error
	ListRoutingProfilesPagesWithContext(aws.Context, *connect.ListRoutingProfilesInput, func(*connect.ListRoutingProfilesOutput, bool) bool, ...request.Option) error

	ListSecurityProfiles(*connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesWithContext(aws.Context, *connect.ListSecurityProfilesInput, ...request.Option) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesRequest(*connect.ListSecurityProfilesInput) (*request.Request, *connect.ListSecurityProfilesOutput)

	ListSecurityProfilesPages(*connect.ListSecurityProfilesInput, func(*connect.ListSecurityProfilesOutput, bool) bool) error
	ListSecurityProfilesPagesWithContext(aws.Context, *connect.ListSecurityProfilesInput, func(*connect.ListSecurityProfilesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *connect.ListTagsForResourceInput, ...request.Option) (*connect.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*connect.ListTagsForResourceInput) (*request.Request, *connect.ListTagsForResourceOutput)

	ListUserHierarchyGroups(*connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsWithContext(aws.Context, *connect.ListUserHierarchyGroupsInput, ...request.Option) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsRequest(*connect.ListUserHierarchyGroupsInput) (*request.Request, *connect.ListUserHierarchyGroupsOutput)

	ListUserHierarchyGroupsPages(*connect.ListUserHierarchyGroupsInput, func(*connect.ListUserHierarchyGroupsOutput, bool) bool) error
	ListUserHierarchyGroupsPagesWithContext(aws.Context, *connect.ListUserHierarchyGroupsInput, func(*connect.ListUserHierarchyGroupsOutput, bool) bool, ...request.Option) error

	ListUsers(*connect.ListUsersInput) (*connect.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *connect.ListUsersInput, ...request.Option) (*connect.ListUsersOutput, error)
	ListUsersRequest(*connect.ListUsersInput) (*request.Request, *connect.ListUsersOutput)

	ListUsersPages(*connect.ListUsersInput, func(*connect.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *connect.ListUsersInput, func(*connect.ListUsersOutput, bool) bool, ...request.Option) error

	ResumeContactRecording(*connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error)
	ResumeContactRecordingWithContext(aws.Context, *connect.ResumeContactRecordingInput, ...request.Option) (*connect.ResumeContactRecordingOutput, error)
	ResumeContactRecordingRequest(*connect.ResumeContactRecordingInput) (*request.Request, *connect.ResumeContactRecordingOutput)

	StartChatContact(*connect.StartChatContactInput) (*connect.StartChatContactOutput, error)
	StartChatContactWithContext(aws.Context, *connect.StartChatContactInput, ...request.Option) (*connect.StartChatContactOutput, error)
	StartChatContactRequest(*connect.StartChatContactInput) (*request.Request, *connect.StartChatContactOutput)

	StartContactRecording(*connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error)
	StartContactRecordingWithContext(aws.Context, *connect.StartContactRecordingInput, ...request.Option) (*connect.StartContactRecordingOutput, error)
	StartContactRecordingRequest(*connect.StartContactRecordingInput) (*request.Request, *connect.StartContactRecordingOutput)

	StartOutboundVoiceContact(*connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactWithContext(aws.Context, *connect.StartOutboundVoiceContactInput, ...request.Option) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactRequest(*connect.StartOutboundVoiceContactInput) (*request.Request, *connect.StartOutboundVoiceContactOutput)

	StopContact(*connect.StopContactInput) (*connect.StopContactOutput, error)
	StopContactWithContext(aws.Context, *connect.StopContactInput, ...request.Option) (*connect.StopContactOutput, error)
	StopContactRequest(*connect.StopContactInput) (*request.Request, *connect.StopContactOutput)

	StopContactRecording(*connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error)
	StopContactRecordingWithContext(aws.Context, *connect.StopContactRecordingInput, ...request.Option) (*connect.StopContactRecordingOutput, error)
	StopContactRecordingRequest(*connect.StopContactRecordingInput) (*request.Request, *connect.StopContactRecordingOutput)

	SuspendContactRecording(*connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error)
	SuspendContactRecordingWithContext(aws.Context, *connect.SuspendContactRecordingInput, ...request.Option) (*connect.SuspendContactRecordingOutput, error)
	SuspendContactRecordingRequest(*connect.SuspendContactRecordingInput) (*request.Request, *connect.SuspendContactRecordingOutput)

	TagResource(*connect.TagResourceInput) (*connect.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *connect.TagResourceInput, ...request.Option) (*connect.TagResourceOutput, error)
	TagResourceRequest(*connect.TagResourceInput) (*request.Request, *connect.TagResourceOutput)

	UntagResource(*connect.UntagResourceInput) (*connect.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *connect.UntagResourceInput, ...request.Option) (*connect.UntagResourceOutput, error)
	UntagResourceRequest(*connect.UntagResourceInput) (*request.Request, *connect.UntagResourceOutput)

	UpdateContactAttributes(*connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesWithContext(aws.Context, *connect.UpdateContactAttributesInput, ...request.Option) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesRequest(*connect.UpdateContactAttributesInput) (*request.Request, *connect.UpdateContactAttributesOutput)

	UpdateContactFlowContent(*connect.UpdateContactFlowContentInput) (*connect.UpdateContactFlowContentOutput, error)
	UpdateContactFlowContentWithContext(aws.Context, *connect.UpdateContactFlowContentInput, ...request.Option) (*connect.UpdateContactFlowContentOutput, error)
	UpdateContactFlowContentRequest(*connect.UpdateContactFlowContentInput) (*request.Request, *connect.UpdateContactFlowContentOutput)

	UpdateContactFlowName(*connect.UpdateContactFlowNameInput) (*connect.UpdateContactFlowNameOutput, error)
	UpdateContactFlowNameWithContext(aws.Context, *connect.UpdateContactFlowNameInput, ...request.Option) (*connect.UpdateContactFlowNameOutput, error)
	UpdateContactFlowNameRequest(*connect.UpdateContactFlowNameInput) (*request.Request, *connect.UpdateContactFlowNameOutput)

	UpdateRoutingProfileConcurrency(*connect.UpdateRoutingProfileConcurrencyInput) (*connect.UpdateRoutingProfileConcurrencyOutput, error)
	UpdateRoutingProfileConcurrencyWithContext(aws.Context, *connect.UpdateRoutingProfileConcurrencyInput, ...request.Option) (*connect.UpdateRoutingProfileConcurrencyOutput, error)
	UpdateRoutingProfileConcurrencyRequest(*connect.UpdateRoutingProfileConcurrencyInput) (*request.Request, *connect.UpdateRoutingProfileConcurrencyOutput)

	UpdateRoutingProfileDefaultOutboundQueue(*connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error)
	UpdateRoutingProfileDefaultOutboundQueueWithContext(aws.Context, *connect.UpdateRoutingProfileDefaultOutboundQueueInput, ...request.Option) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error)
	UpdateRoutingProfileDefaultOutboundQueueRequest(*connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*request.Request, *connect.UpdateRoutingProfileDefaultOutboundQueueOutput)

	UpdateRoutingProfileName(*connect.UpdateRoutingProfileNameInput) (*connect.UpdateRoutingProfileNameOutput, error)
	UpdateRoutingProfileNameWithContext(aws.Context, *connect.UpdateRoutingProfileNameInput, ...request.Option) (*connect.UpdateRoutingProfileNameOutput, error)
	UpdateRoutingProfileNameRequest(*connect.UpdateRoutingProfileNameInput) (*request.Request, *connect.UpdateRoutingProfileNameOutput)

	UpdateRoutingProfileQueues(*connect.UpdateRoutingProfileQueuesInput) (*connect.UpdateRoutingProfileQueuesOutput, error)
	UpdateRoutingProfileQueuesWithContext(aws.Context, *connect.UpdateRoutingProfileQueuesInput, ...request.Option) (*connect.UpdateRoutingProfileQueuesOutput, error)
	UpdateRoutingProfileQueuesRequest(*connect.UpdateRoutingProfileQueuesInput) (*request.Request, *connect.UpdateRoutingProfileQueuesOutput)

	UpdateUserHierarchy(*connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyWithContext(aws.Context, *connect.UpdateUserHierarchyInput, ...request.Option) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyRequest(*connect.UpdateUserHierarchyInput) (*request.Request, *connect.UpdateUserHierarchyOutput)

	UpdateUserIdentityInfo(*connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoWithContext(aws.Context, *connect.UpdateUserIdentityInfoInput, ...request.Option) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoRequest(*connect.UpdateUserIdentityInfoInput) (*request.Request, *connect.UpdateUserIdentityInfoOutput)

	UpdateUserPhoneConfig(*connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigWithContext(aws.Context, *connect.UpdateUserPhoneConfigInput, ...request.Option) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigRequest(*connect.UpdateUserPhoneConfigInput) (*request.Request, *connect.UpdateUserPhoneConfigOutput)

	UpdateUserRoutingProfile(*connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileWithContext(aws.Context, *connect.UpdateUserRoutingProfileInput, ...request.Option) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileRequest(*connect.UpdateUserRoutingProfileInput) (*request.Request, *connect.UpdateUserRoutingProfileOutput)

	UpdateUserSecurityProfiles(*connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesWithContext(aws.Context, *connect.UpdateUserSecurityProfilesInput, ...request.Option) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesRequest(*connect.UpdateUserSecurityProfilesInput) (*request.Request, *connect.UpdateUserSecurityProfilesOutput)
}

var _ ConnectAPI = (*connect.Connect)(nil)