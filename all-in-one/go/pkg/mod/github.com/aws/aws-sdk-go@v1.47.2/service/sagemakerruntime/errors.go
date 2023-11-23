// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalDependencyException for service response error code
	// "InternalDependencyException".
	//
	// Your request caused an exception with an internal dependency. Contact customer
	// support.
	ErrCodeInternalDependencyException = "InternalDependencyException"

	// ErrCodeInternalFailure for service response error code
	// "InternalFailure".
	//
	// An internal failure occurred.
	ErrCodeInternalFailure = "InternalFailure"

	// ErrCodeInternalStreamFailure for service response error code
	// "InternalStreamFailure".
	//
	// The stream processing failed because of an unknown error, exception or failure.
	// Try your request again.
	ErrCodeInternalStreamFailure = "InternalStreamFailure"

	// ErrCodeModelError for service response error code
	// "ModelError".
	//
	// Model (owned by the customer in the container) returned 4xx or 5xx error
	// code.
	ErrCodeModelError = "ModelError"

	// ErrCodeModelNotReadyException for service response error code
	// "ModelNotReadyException".
	//
	// Either a serverless endpoint variant's resources are still being provisioned,
	// or a multi-model endpoint is still downloading or loading the target model.
	// Wait and try your request again.
	ErrCodeModelNotReadyException = "ModelNotReadyException"

	// ErrCodeModelStreamError for service response error code
	// "ModelStreamError".
	//
	// An error occurred while streaming the response body. This error can have
	// the following error codes:
	//
	// ModelInvocationTimeExceeded
	//
	// The model failed to finish sending the response within the timeout period
	// allowed by Amazon SageMaker.
	//
	// StreamBroken
	//
	// The Transmission Control Protocol (TCP) connection between the client and
	// the model was reset or closed.
	ErrCodeModelStreamError = "ModelStreamError"

	// ErrCodeServiceUnavailable for service response error code
	// "ServiceUnavailable".
	//
	// The service is unavailable. Try your call again.
	ErrCodeServiceUnavailable = "ServiceUnavailable"

	// ErrCodeValidationError for service response error code
	// "ValidationError".
	//
	// Inspect your request and try again.
	ErrCodeValidationError = "ValidationError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalDependencyException": newErrorInternalDependencyException,
	"InternalFailure":             newErrorInternalFailure,
	"InternalStreamFailure":       newErrorInternalStreamFailure,
	"ModelError":                  newErrorModelError,
	"ModelNotReadyException":      newErrorModelNotReadyException,
	"ModelStreamError":            newErrorModelStreamError,
	"ServiceUnavailable":          newErrorServiceUnavailable,
	"ValidationError":             newErrorValidationError,
}
