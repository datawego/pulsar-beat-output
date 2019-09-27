// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/GetReservedNodeExchangeOfferingsInputMessage
type GetReservedNodeExchangeOfferingsInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of ReservedNodeOfferings.
	Marker *string `type:"string"`

	// An integer setting the maximum number of ReservedNodeOfferings to retrieve.
	MaxRecords *int64 `type:"integer"`

	// A string representing the node identifier for the DC1 Reserved Node to be
	// exchanged.
	//
	// ReservedNodeId is a required field
	ReservedNodeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetReservedNodeExchangeOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservedNodeExchangeOfferingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservedNodeExchangeOfferingsInput"}

	if s.ReservedNodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedNodeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/GetReservedNodeExchangeOfferingsOutputMessage
type GetReservedNodeExchangeOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point for returning a set
	// of response records. When the results of a GetReservedNodeExchangeOfferings
	// request exceed the value specified in MaxRecords, Amazon Redshift returns
	// a value in the marker field of the response. You can retrieve the next set
	// of response records by providing the returned marker value in the marker
	// parameter and retrying the request.
	Marker *string `type:"string"`

	// Returns an array of ReservedNodeOffering objects.
	ReservedNodeOfferings []ReservedNodeOffering `locationNameList:"ReservedNodeOffering" type:"list"`
}

// String returns the string representation
func (s GetReservedNodeExchangeOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReservedNodeExchangeOfferings = "GetReservedNodeExchangeOfferings"

// GetReservedNodeExchangeOfferingsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns an array of DC2 ReservedNodeOfferings that matches the payment type,
// term, and usage price of the given DC1 reserved node.
//
//    // Example sending a request using GetReservedNodeExchangeOfferingsRequest.
//    req := client.GetReservedNodeExchangeOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/GetReservedNodeExchangeOfferings
func (c *Client) GetReservedNodeExchangeOfferingsRequest(input *GetReservedNodeExchangeOfferingsInput) GetReservedNodeExchangeOfferingsRequest {
	op := &aws.Operation{
		Name:       opGetReservedNodeExchangeOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservedNodeExchangeOfferingsInput{}
	}

	req := c.newRequest(op, input, &GetReservedNodeExchangeOfferingsOutput{})
	return GetReservedNodeExchangeOfferingsRequest{Request: req, Input: input, Copy: c.GetReservedNodeExchangeOfferingsRequest}
}

// GetReservedNodeExchangeOfferingsRequest is the request type for the
// GetReservedNodeExchangeOfferings API operation.
type GetReservedNodeExchangeOfferingsRequest struct {
	*aws.Request
	Input *GetReservedNodeExchangeOfferingsInput
	Copy  func(*GetReservedNodeExchangeOfferingsInput) GetReservedNodeExchangeOfferingsRequest
}

// Send marshals and sends the GetReservedNodeExchangeOfferings API request.
func (r GetReservedNodeExchangeOfferingsRequest) Send(ctx context.Context) (*GetReservedNodeExchangeOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservedNodeExchangeOfferingsResponse{
		GetReservedNodeExchangeOfferingsOutput: r.Request.Data.(*GetReservedNodeExchangeOfferingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservedNodeExchangeOfferingsResponse is the response type for the
// GetReservedNodeExchangeOfferings API operation.
type GetReservedNodeExchangeOfferingsResponse struct {
	*GetReservedNodeExchangeOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservedNodeExchangeOfferings request.
func (r *GetReservedNodeExchangeOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}