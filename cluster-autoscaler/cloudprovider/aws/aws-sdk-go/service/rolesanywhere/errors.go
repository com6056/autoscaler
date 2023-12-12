// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rolesanywhere

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// Too many tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Validation exception error.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"TooManyTagsException":      newErrorTooManyTagsException,
	"ValidationException":       newErrorValidationException,
}
