// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package paymentcryptographydataiface provides an interface to enable mocking the Payment Cryptography Data Plane service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package paymentcryptographydataiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/paymentcryptographydata"
)

// PaymentCryptographyDataAPI provides an interface to enable mocking the
// paymentcryptographydata.PaymentCryptographyData service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Payment Cryptography Data Plane.
//	func myFunc(svc paymentcryptographydataiface.PaymentCryptographyDataAPI) bool {
//	    // Make svc.DecryptData request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := paymentcryptographydata.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPaymentCryptographyDataClient struct {
//	    paymentcryptographydataiface.PaymentCryptographyDataAPI
//	}
//	func (m *mockPaymentCryptographyDataClient) DecryptData(input *paymentcryptographydata.DecryptDataInput) (*paymentcryptographydata.DecryptDataOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPaymentCryptographyDataClient{}
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
type PaymentCryptographyDataAPI interface {
	DecryptData(*paymentcryptographydata.DecryptDataInput) (*paymentcryptographydata.DecryptDataOutput, error)
	DecryptDataWithContext(aws.Context, *paymentcryptographydata.DecryptDataInput, ...request.Option) (*paymentcryptographydata.DecryptDataOutput, error)
	DecryptDataRequest(*paymentcryptographydata.DecryptDataInput) (*request.Request, *paymentcryptographydata.DecryptDataOutput)

	EncryptData(*paymentcryptographydata.EncryptDataInput) (*paymentcryptographydata.EncryptDataOutput, error)
	EncryptDataWithContext(aws.Context, *paymentcryptographydata.EncryptDataInput, ...request.Option) (*paymentcryptographydata.EncryptDataOutput, error)
	EncryptDataRequest(*paymentcryptographydata.EncryptDataInput) (*request.Request, *paymentcryptographydata.EncryptDataOutput)

	GenerateCardValidationData(*paymentcryptographydata.GenerateCardValidationDataInput) (*paymentcryptographydata.GenerateCardValidationDataOutput, error)
	GenerateCardValidationDataWithContext(aws.Context, *paymentcryptographydata.GenerateCardValidationDataInput, ...request.Option) (*paymentcryptographydata.GenerateCardValidationDataOutput, error)
	GenerateCardValidationDataRequest(*paymentcryptographydata.GenerateCardValidationDataInput) (*request.Request, *paymentcryptographydata.GenerateCardValidationDataOutput)

	GenerateMac(*paymentcryptographydata.GenerateMacInput) (*paymentcryptographydata.GenerateMacOutput, error)
	GenerateMacWithContext(aws.Context, *paymentcryptographydata.GenerateMacInput, ...request.Option) (*paymentcryptographydata.GenerateMacOutput, error)
	GenerateMacRequest(*paymentcryptographydata.GenerateMacInput) (*request.Request, *paymentcryptographydata.GenerateMacOutput)

	GeneratePinData(*paymentcryptographydata.GeneratePinDataInput) (*paymentcryptographydata.GeneratePinDataOutput, error)
	GeneratePinDataWithContext(aws.Context, *paymentcryptographydata.GeneratePinDataInput, ...request.Option) (*paymentcryptographydata.GeneratePinDataOutput, error)
	GeneratePinDataRequest(*paymentcryptographydata.GeneratePinDataInput) (*request.Request, *paymentcryptographydata.GeneratePinDataOutput)

	ReEncryptData(*paymentcryptographydata.ReEncryptDataInput) (*paymentcryptographydata.ReEncryptDataOutput, error)
	ReEncryptDataWithContext(aws.Context, *paymentcryptographydata.ReEncryptDataInput, ...request.Option) (*paymentcryptographydata.ReEncryptDataOutput, error)
	ReEncryptDataRequest(*paymentcryptographydata.ReEncryptDataInput) (*request.Request, *paymentcryptographydata.ReEncryptDataOutput)

	TranslatePinData(*paymentcryptographydata.TranslatePinDataInput) (*paymentcryptographydata.TranslatePinDataOutput, error)
	TranslatePinDataWithContext(aws.Context, *paymentcryptographydata.TranslatePinDataInput, ...request.Option) (*paymentcryptographydata.TranslatePinDataOutput, error)
	TranslatePinDataRequest(*paymentcryptographydata.TranslatePinDataInput) (*request.Request, *paymentcryptographydata.TranslatePinDataOutput)

	VerifyAuthRequestCryptogram(*paymentcryptographydata.VerifyAuthRequestCryptogramInput) (*paymentcryptographydata.VerifyAuthRequestCryptogramOutput, error)
	VerifyAuthRequestCryptogramWithContext(aws.Context, *paymentcryptographydata.VerifyAuthRequestCryptogramInput, ...request.Option) (*paymentcryptographydata.VerifyAuthRequestCryptogramOutput, error)
	VerifyAuthRequestCryptogramRequest(*paymentcryptographydata.VerifyAuthRequestCryptogramInput) (*request.Request, *paymentcryptographydata.VerifyAuthRequestCryptogramOutput)

	VerifyCardValidationData(*paymentcryptographydata.VerifyCardValidationDataInput) (*paymentcryptographydata.VerifyCardValidationDataOutput, error)
	VerifyCardValidationDataWithContext(aws.Context, *paymentcryptographydata.VerifyCardValidationDataInput, ...request.Option) (*paymentcryptographydata.VerifyCardValidationDataOutput, error)
	VerifyCardValidationDataRequest(*paymentcryptographydata.VerifyCardValidationDataInput) (*request.Request, *paymentcryptographydata.VerifyCardValidationDataOutput)

	VerifyMac(*paymentcryptographydata.VerifyMacInput) (*paymentcryptographydata.VerifyMacOutput, error)
	VerifyMacWithContext(aws.Context, *paymentcryptographydata.VerifyMacInput, ...request.Option) (*paymentcryptographydata.VerifyMacOutput, error)
	VerifyMacRequest(*paymentcryptographydata.VerifyMacInput) (*request.Request, *paymentcryptographydata.VerifyMacOutput)

	VerifyPinData(*paymentcryptographydata.VerifyPinDataInput) (*paymentcryptographydata.VerifyPinDataOutput, error)
	VerifyPinDataWithContext(aws.Context, *paymentcryptographydata.VerifyPinDataInput, ...request.Option) (*paymentcryptographydata.VerifyPinDataOutput, error)
	VerifyPinDataRequest(*paymentcryptographydata.VerifyPinDataInput) (*request.Request, *paymentcryptographydata.VerifyPinDataOutput)
}

var _ PaymentCryptographyDataAPI = (*paymentcryptographydata.PaymentCryptographyData)(nil)
