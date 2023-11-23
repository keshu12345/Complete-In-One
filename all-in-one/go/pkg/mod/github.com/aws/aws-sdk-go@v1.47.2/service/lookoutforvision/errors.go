// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lookoutforvision

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You are not authorized to perform the action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The update or deletion of a resource caused an inconsistent state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Amazon Lookout for Vision experienced a service issue. Try your call again.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// A service quota was exceeded the allowed limit. For more information, see
	// Limits in Amazon Lookout for Vision in the Amazon Lookout for Vision Developer
	// Guide.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Amazon Lookout for Vision is temporarily unable to process the request. Try
	// your call again.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// An input validation error occured. For example, invalid characters in a project
	// name, or if a pagination token is invalid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
