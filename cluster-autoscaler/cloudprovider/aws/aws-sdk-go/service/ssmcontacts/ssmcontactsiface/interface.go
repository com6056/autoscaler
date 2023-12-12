// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ssmcontactsiface provides an interface to enable mocking the AWS Systems Manager Incident Manager Contacts service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ssmcontactsiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/ssmcontacts"
)

// SSMContactsAPI provides an interface to enable mocking the
// ssmcontacts.SSMContacts service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Systems Manager Incident Manager Contacts.
//	func myFunc(svc ssmcontactsiface.SSMContactsAPI) bool {
//	    // Make svc.AcceptPage request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := ssmcontacts.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSSMContactsClient struct {
//	    ssmcontactsiface.SSMContactsAPI
//	}
//	func (m *mockSSMContactsClient) AcceptPage(input *ssmcontacts.AcceptPageInput) (*ssmcontacts.AcceptPageOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSSMContactsClient{}
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
type SSMContactsAPI interface {
	AcceptPage(*ssmcontacts.AcceptPageInput) (*ssmcontacts.AcceptPageOutput, error)
	AcceptPageWithContext(aws.Context, *ssmcontacts.AcceptPageInput, ...request.Option) (*ssmcontacts.AcceptPageOutput, error)
	AcceptPageRequest(*ssmcontacts.AcceptPageInput) (*request.Request, *ssmcontacts.AcceptPageOutput)

	ActivateContactChannel(*ssmcontacts.ActivateContactChannelInput) (*ssmcontacts.ActivateContactChannelOutput, error)
	ActivateContactChannelWithContext(aws.Context, *ssmcontacts.ActivateContactChannelInput, ...request.Option) (*ssmcontacts.ActivateContactChannelOutput, error)
	ActivateContactChannelRequest(*ssmcontacts.ActivateContactChannelInput) (*request.Request, *ssmcontacts.ActivateContactChannelOutput)

	CreateContact(*ssmcontacts.CreateContactInput) (*ssmcontacts.CreateContactOutput, error)
	CreateContactWithContext(aws.Context, *ssmcontacts.CreateContactInput, ...request.Option) (*ssmcontacts.CreateContactOutput, error)
	CreateContactRequest(*ssmcontacts.CreateContactInput) (*request.Request, *ssmcontacts.CreateContactOutput)

	CreateContactChannel(*ssmcontacts.CreateContactChannelInput) (*ssmcontacts.CreateContactChannelOutput, error)
	CreateContactChannelWithContext(aws.Context, *ssmcontacts.CreateContactChannelInput, ...request.Option) (*ssmcontacts.CreateContactChannelOutput, error)
	CreateContactChannelRequest(*ssmcontacts.CreateContactChannelInput) (*request.Request, *ssmcontacts.CreateContactChannelOutput)

	CreateRotation(*ssmcontacts.CreateRotationInput) (*ssmcontacts.CreateRotationOutput, error)
	CreateRotationWithContext(aws.Context, *ssmcontacts.CreateRotationInput, ...request.Option) (*ssmcontacts.CreateRotationOutput, error)
	CreateRotationRequest(*ssmcontacts.CreateRotationInput) (*request.Request, *ssmcontacts.CreateRotationOutput)

	CreateRotationOverride(*ssmcontacts.CreateRotationOverrideInput) (*ssmcontacts.CreateRotationOverrideOutput, error)
	CreateRotationOverrideWithContext(aws.Context, *ssmcontacts.CreateRotationOverrideInput, ...request.Option) (*ssmcontacts.CreateRotationOverrideOutput, error)
	CreateRotationOverrideRequest(*ssmcontacts.CreateRotationOverrideInput) (*request.Request, *ssmcontacts.CreateRotationOverrideOutput)

	DeactivateContactChannel(*ssmcontacts.DeactivateContactChannelInput) (*ssmcontacts.DeactivateContactChannelOutput, error)
	DeactivateContactChannelWithContext(aws.Context, *ssmcontacts.DeactivateContactChannelInput, ...request.Option) (*ssmcontacts.DeactivateContactChannelOutput, error)
	DeactivateContactChannelRequest(*ssmcontacts.DeactivateContactChannelInput) (*request.Request, *ssmcontacts.DeactivateContactChannelOutput)

	DeleteContact(*ssmcontacts.DeleteContactInput) (*ssmcontacts.DeleteContactOutput, error)
	DeleteContactWithContext(aws.Context, *ssmcontacts.DeleteContactInput, ...request.Option) (*ssmcontacts.DeleteContactOutput, error)
	DeleteContactRequest(*ssmcontacts.DeleteContactInput) (*request.Request, *ssmcontacts.DeleteContactOutput)

	DeleteContactChannel(*ssmcontacts.DeleteContactChannelInput) (*ssmcontacts.DeleteContactChannelOutput, error)
	DeleteContactChannelWithContext(aws.Context, *ssmcontacts.DeleteContactChannelInput, ...request.Option) (*ssmcontacts.DeleteContactChannelOutput, error)
	DeleteContactChannelRequest(*ssmcontacts.DeleteContactChannelInput) (*request.Request, *ssmcontacts.DeleteContactChannelOutput)

	DeleteRotation(*ssmcontacts.DeleteRotationInput) (*ssmcontacts.DeleteRotationOutput, error)
	DeleteRotationWithContext(aws.Context, *ssmcontacts.DeleteRotationInput, ...request.Option) (*ssmcontacts.DeleteRotationOutput, error)
	DeleteRotationRequest(*ssmcontacts.DeleteRotationInput) (*request.Request, *ssmcontacts.DeleteRotationOutput)

	DeleteRotationOverride(*ssmcontacts.DeleteRotationOverrideInput) (*ssmcontacts.DeleteRotationOverrideOutput, error)
	DeleteRotationOverrideWithContext(aws.Context, *ssmcontacts.DeleteRotationOverrideInput, ...request.Option) (*ssmcontacts.DeleteRotationOverrideOutput, error)
	DeleteRotationOverrideRequest(*ssmcontacts.DeleteRotationOverrideInput) (*request.Request, *ssmcontacts.DeleteRotationOverrideOutput)

	DescribeEngagement(*ssmcontacts.DescribeEngagementInput) (*ssmcontacts.DescribeEngagementOutput, error)
	DescribeEngagementWithContext(aws.Context, *ssmcontacts.DescribeEngagementInput, ...request.Option) (*ssmcontacts.DescribeEngagementOutput, error)
	DescribeEngagementRequest(*ssmcontacts.DescribeEngagementInput) (*request.Request, *ssmcontacts.DescribeEngagementOutput)

	DescribePage(*ssmcontacts.DescribePageInput) (*ssmcontacts.DescribePageOutput, error)
	DescribePageWithContext(aws.Context, *ssmcontacts.DescribePageInput, ...request.Option) (*ssmcontacts.DescribePageOutput, error)
	DescribePageRequest(*ssmcontacts.DescribePageInput) (*request.Request, *ssmcontacts.DescribePageOutput)

	GetContact(*ssmcontacts.GetContactInput) (*ssmcontacts.GetContactOutput, error)
	GetContactWithContext(aws.Context, *ssmcontacts.GetContactInput, ...request.Option) (*ssmcontacts.GetContactOutput, error)
	GetContactRequest(*ssmcontacts.GetContactInput) (*request.Request, *ssmcontacts.GetContactOutput)

	GetContactChannel(*ssmcontacts.GetContactChannelInput) (*ssmcontacts.GetContactChannelOutput, error)
	GetContactChannelWithContext(aws.Context, *ssmcontacts.GetContactChannelInput, ...request.Option) (*ssmcontacts.GetContactChannelOutput, error)
	GetContactChannelRequest(*ssmcontacts.GetContactChannelInput) (*request.Request, *ssmcontacts.GetContactChannelOutput)

	GetContactPolicy(*ssmcontacts.GetContactPolicyInput) (*ssmcontacts.GetContactPolicyOutput, error)
	GetContactPolicyWithContext(aws.Context, *ssmcontacts.GetContactPolicyInput, ...request.Option) (*ssmcontacts.GetContactPolicyOutput, error)
	GetContactPolicyRequest(*ssmcontacts.GetContactPolicyInput) (*request.Request, *ssmcontacts.GetContactPolicyOutput)

	GetRotation(*ssmcontacts.GetRotationInput) (*ssmcontacts.GetRotationOutput, error)
	GetRotationWithContext(aws.Context, *ssmcontacts.GetRotationInput, ...request.Option) (*ssmcontacts.GetRotationOutput, error)
	GetRotationRequest(*ssmcontacts.GetRotationInput) (*request.Request, *ssmcontacts.GetRotationOutput)

	GetRotationOverride(*ssmcontacts.GetRotationOverrideInput) (*ssmcontacts.GetRotationOverrideOutput, error)
	GetRotationOverrideWithContext(aws.Context, *ssmcontacts.GetRotationOverrideInput, ...request.Option) (*ssmcontacts.GetRotationOverrideOutput, error)
	GetRotationOverrideRequest(*ssmcontacts.GetRotationOverrideInput) (*request.Request, *ssmcontacts.GetRotationOverrideOutput)

	ListContactChannels(*ssmcontacts.ListContactChannelsInput) (*ssmcontacts.ListContactChannelsOutput, error)
	ListContactChannelsWithContext(aws.Context, *ssmcontacts.ListContactChannelsInput, ...request.Option) (*ssmcontacts.ListContactChannelsOutput, error)
	ListContactChannelsRequest(*ssmcontacts.ListContactChannelsInput) (*request.Request, *ssmcontacts.ListContactChannelsOutput)

	ListContactChannelsPages(*ssmcontacts.ListContactChannelsInput, func(*ssmcontacts.ListContactChannelsOutput, bool) bool) error
	ListContactChannelsPagesWithContext(aws.Context, *ssmcontacts.ListContactChannelsInput, func(*ssmcontacts.ListContactChannelsOutput, bool) bool, ...request.Option) error

	ListContacts(*ssmcontacts.ListContactsInput) (*ssmcontacts.ListContactsOutput, error)
	ListContactsWithContext(aws.Context, *ssmcontacts.ListContactsInput, ...request.Option) (*ssmcontacts.ListContactsOutput, error)
	ListContactsRequest(*ssmcontacts.ListContactsInput) (*request.Request, *ssmcontacts.ListContactsOutput)

	ListContactsPages(*ssmcontacts.ListContactsInput, func(*ssmcontacts.ListContactsOutput, bool) bool) error
	ListContactsPagesWithContext(aws.Context, *ssmcontacts.ListContactsInput, func(*ssmcontacts.ListContactsOutput, bool) bool, ...request.Option) error

	ListEngagements(*ssmcontacts.ListEngagementsInput) (*ssmcontacts.ListEngagementsOutput, error)
	ListEngagementsWithContext(aws.Context, *ssmcontacts.ListEngagementsInput, ...request.Option) (*ssmcontacts.ListEngagementsOutput, error)
	ListEngagementsRequest(*ssmcontacts.ListEngagementsInput) (*request.Request, *ssmcontacts.ListEngagementsOutput)

	ListEngagementsPages(*ssmcontacts.ListEngagementsInput, func(*ssmcontacts.ListEngagementsOutput, bool) bool) error
	ListEngagementsPagesWithContext(aws.Context, *ssmcontacts.ListEngagementsInput, func(*ssmcontacts.ListEngagementsOutput, bool) bool, ...request.Option) error

	ListPageReceipts(*ssmcontacts.ListPageReceiptsInput) (*ssmcontacts.ListPageReceiptsOutput, error)
	ListPageReceiptsWithContext(aws.Context, *ssmcontacts.ListPageReceiptsInput, ...request.Option) (*ssmcontacts.ListPageReceiptsOutput, error)
	ListPageReceiptsRequest(*ssmcontacts.ListPageReceiptsInput) (*request.Request, *ssmcontacts.ListPageReceiptsOutput)

	ListPageReceiptsPages(*ssmcontacts.ListPageReceiptsInput, func(*ssmcontacts.ListPageReceiptsOutput, bool) bool) error
	ListPageReceiptsPagesWithContext(aws.Context, *ssmcontacts.ListPageReceiptsInput, func(*ssmcontacts.ListPageReceiptsOutput, bool) bool, ...request.Option) error

	ListPageResolutions(*ssmcontacts.ListPageResolutionsInput) (*ssmcontacts.ListPageResolutionsOutput, error)
	ListPageResolutionsWithContext(aws.Context, *ssmcontacts.ListPageResolutionsInput, ...request.Option) (*ssmcontacts.ListPageResolutionsOutput, error)
	ListPageResolutionsRequest(*ssmcontacts.ListPageResolutionsInput) (*request.Request, *ssmcontacts.ListPageResolutionsOutput)

	ListPageResolutionsPages(*ssmcontacts.ListPageResolutionsInput, func(*ssmcontacts.ListPageResolutionsOutput, bool) bool) error
	ListPageResolutionsPagesWithContext(aws.Context, *ssmcontacts.ListPageResolutionsInput, func(*ssmcontacts.ListPageResolutionsOutput, bool) bool, ...request.Option) error

	ListPagesByContact(*ssmcontacts.ListPagesByContactInput) (*ssmcontacts.ListPagesByContactOutput, error)
	ListPagesByContactWithContext(aws.Context, *ssmcontacts.ListPagesByContactInput, ...request.Option) (*ssmcontacts.ListPagesByContactOutput, error)
	ListPagesByContactRequest(*ssmcontacts.ListPagesByContactInput) (*request.Request, *ssmcontacts.ListPagesByContactOutput)

	ListPagesByContactPages(*ssmcontacts.ListPagesByContactInput, func(*ssmcontacts.ListPagesByContactOutput, bool) bool) error
	ListPagesByContactPagesWithContext(aws.Context, *ssmcontacts.ListPagesByContactInput, func(*ssmcontacts.ListPagesByContactOutput, bool) bool, ...request.Option) error

	ListPagesByEngagement(*ssmcontacts.ListPagesByEngagementInput) (*ssmcontacts.ListPagesByEngagementOutput, error)
	ListPagesByEngagementWithContext(aws.Context, *ssmcontacts.ListPagesByEngagementInput, ...request.Option) (*ssmcontacts.ListPagesByEngagementOutput, error)
	ListPagesByEngagementRequest(*ssmcontacts.ListPagesByEngagementInput) (*request.Request, *ssmcontacts.ListPagesByEngagementOutput)

	ListPagesByEngagementPages(*ssmcontacts.ListPagesByEngagementInput, func(*ssmcontacts.ListPagesByEngagementOutput, bool) bool) error
	ListPagesByEngagementPagesWithContext(aws.Context, *ssmcontacts.ListPagesByEngagementInput, func(*ssmcontacts.ListPagesByEngagementOutput, bool) bool, ...request.Option) error

	ListPreviewRotationShifts(*ssmcontacts.ListPreviewRotationShiftsInput) (*ssmcontacts.ListPreviewRotationShiftsOutput, error)
	ListPreviewRotationShiftsWithContext(aws.Context, *ssmcontacts.ListPreviewRotationShiftsInput, ...request.Option) (*ssmcontacts.ListPreviewRotationShiftsOutput, error)
	ListPreviewRotationShiftsRequest(*ssmcontacts.ListPreviewRotationShiftsInput) (*request.Request, *ssmcontacts.ListPreviewRotationShiftsOutput)

	ListPreviewRotationShiftsPages(*ssmcontacts.ListPreviewRotationShiftsInput, func(*ssmcontacts.ListPreviewRotationShiftsOutput, bool) bool) error
	ListPreviewRotationShiftsPagesWithContext(aws.Context, *ssmcontacts.ListPreviewRotationShiftsInput, func(*ssmcontacts.ListPreviewRotationShiftsOutput, bool) bool, ...request.Option) error

	ListRotationOverrides(*ssmcontacts.ListRotationOverridesInput) (*ssmcontacts.ListRotationOverridesOutput, error)
	ListRotationOverridesWithContext(aws.Context, *ssmcontacts.ListRotationOverridesInput, ...request.Option) (*ssmcontacts.ListRotationOverridesOutput, error)
	ListRotationOverridesRequest(*ssmcontacts.ListRotationOverridesInput) (*request.Request, *ssmcontacts.ListRotationOverridesOutput)

	ListRotationOverridesPages(*ssmcontacts.ListRotationOverridesInput, func(*ssmcontacts.ListRotationOverridesOutput, bool) bool) error
	ListRotationOverridesPagesWithContext(aws.Context, *ssmcontacts.ListRotationOverridesInput, func(*ssmcontacts.ListRotationOverridesOutput, bool) bool, ...request.Option) error

	ListRotationShifts(*ssmcontacts.ListRotationShiftsInput) (*ssmcontacts.ListRotationShiftsOutput, error)
	ListRotationShiftsWithContext(aws.Context, *ssmcontacts.ListRotationShiftsInput, ...request.Option) (*ssmcontacts.ListRotationShiftsOutput, error)
	ListRotationShiftsRequest(*ssmcontacts.ListRotationShiftsInput) (*request.Request, *ssmcontacts.ListRotationShiftsOutput)

	ListRotationShiftsPages(*ssmcontacts.ListRotationShiftsInput, func(*ssmcontacts.ListRotationShiftsOutput, bool) bool) error
	ListRotationShiftsPagesWithContext(aws.Context, *ssmcontacts.ListRotationShiftsInput, func(*ssmcontacts.ListRotationShiftsOutput, bool) bool, ...request.Option) error

	ListRotations(*ssmcontacts.ListRotationsInput) (*ssmcontacts.ListRotationsOutput, error)
	ListRotationsWithContext(aws.Context, *ssmcontacts.ListRotationsInput, ...request.Option) (*ssmcontacts.ListRotationsOutput, error)
	ListRotationsRequest(*ssmcontacts.ListRotationsInput) (*request.Request, *ssmcontacts.ListRotationsOutput)

	ListRotationsPages(*ssmcontacts.ListRotationsInput, func(*ssmcontacts.ListRotationsOutput, bool) bool) error
	ListRotationsPagesWithContext(aws.Context, *ssmcontacts.ListRotationsInput, func(*ssmcontacts.ListRotationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*ssmcontacts.ListTagsForResourceInput) (*ssmcontacts.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *ssmcontacts.ListTagsForResourceInput, ...request.Option) (*ssmcontacts.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*ssmcontacts.ListTagsForResourceInput) (*request.Request, *ssmcontacts.ListTagsForResourceOutput)

	PutContactPolicy(*ssmcontacts.PutContactPolicyInput) (*ssmcontacts.PutContactPolicyOutput, error)
	PutContactPolicyWithContext(aws.Context, *ssmcontacts.PutContactPolicyInput, ...request.Option) (*ssmcontacts.PutContactPolicyOutput, error)
	PutContactPolicyRequest(*ssmcontacts.PutContactPolicyInput) (*request.Request, *ssmcontacts.PutContactPolicyOutput)

	SendActivationCode(*ssmcontacts.SendActivationCodeInput) (*ssmcontacts.SendActivationCodeOutput, error)
	SendActivationCodeWithContext(aws.Context, *ssmcontacts.SendActivationCodeInput, ...request.Option) (*ssmcontacts.SendActivationCodeOutput, error)
	SendActivationCodeRequest(*ssmcontacts.SendActivationCodeInput) (*request.Request, *ssmcontacts.SendActivationCodeOutput)

	StartEngagement(*ssmcontacts.StartEngagementInput) (*ssmcontacts.StartEngagementOutput, error)
	StartEngagementWithContext(aws.Context, *ssmcontacts.StartEngagementInput, ...request.Option) (*ssmcontacts.StartEngagementOutput, error)
	StartEngagementRequest(*ssmcontacts.StartEngagementInput) (*request.Request, *ssmcontacts.StartEngagementOutput)

	StopEngagement(*ssmcontacts.StopEngagementInput) (*ssmcontacts.StopEngagementOutput, error)
	StopEngagementWithContext(aws.Context, *ssmcontacts.StopEngagementInput, ...request.Option) (*ssmcontacts.StopEngagementOutput, error)
	StopEngagementRequest(*ssmcontacts.StopEngagementInput) (*request.Request, *ssmcontacts.StopEngagementOutput)

	TagResource(*ssmcontacts.TagResourceInput) (*ssmcontacts.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ssmcontacts.TagResourceInput, ...request.Option) (*ssmcontacts.TagResourceOutput, error)
	TagResourceRequest(*ssmcontacts.TagResourceInput) (*request.Request, *ssmcontacts.TagResourceOutput)

	UntagResource(*ssmcontacts.UntagResourceInput) (*ssmcontacts.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ssmcontacts.UntagResourceInput, ...request.Option) (*ssmcontacts.UntagResourceOutput, error)
	UntagResourceRequest(*ssmcontacts.UntagResourceInput) (*request.Request, *ssmcontacts.UntagResourceOutput)

	UpdateContact(*ssmcontacts.UpdateContactInput) (*ssmcontacts.UpdateContactOutput, error)
	UpdateContactWithContext(aws.Context, *ssmcontacts.UpdateContactInput, ...request.Option) (*ssmcontacts.UpdateContactOutput, error)
	UpdateContactRequest(*ssmcontacts.UpdateContactInput) (*request.Request, *ssmcontacts.UpdateContactOutput)

	UpdateContactChannel(*ssmcontacts.UpdateContactChannelInput) (*ssmcontacts.UpdateContactChannelOutput, error)
	UpdateContactChannelWithContext(aws.Context, *ssmcontacts.UpdateContactChannelInput, ...request.Option) (*ssmcontacts.UpdateContactChannelOutput, error)
	UpdateContactChannelRequest(*ssmcontacts.UpdateContactChannelInput) (*request.Request, *ssmcontacts.UpdateContactChannelOutput)

	UpdateRotation(*ssmcontacts.UpdateRotationInput) (*ssmcontacts.UpdateRotationOutput, error)
	UpdateRotationWithContext(aws.Context, *ssmcontacts.UpdateRotationInput, ...request.Option) (*ssmcontacts.UpdateRotationOutput, error)
	UpdateRotationRequest(*ssmcontacts.UpdateRotationInput) (*request.Request, *ssmcontacts.UpdateRotationOutput)
}

var _ SSMContactsAPI = (*ssmcontacts.SSMContacts)(nil)
