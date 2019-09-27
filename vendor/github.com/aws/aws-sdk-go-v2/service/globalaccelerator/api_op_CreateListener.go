// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/CreateListenerRequest
type CreateListenerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of your accelerator.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of
	// the client request. Clienty affinity gives you control over whether to always
	// route each client to the same specific endpoint.
	//
	// AWS Global Accelerator uses a consistent-flow hashing algorithm to choose
	// the optimal endpoint for a connection. If client affinity is NONE, Global
	// Accelerator uses the "five-tuple" (5-tuple) properties—source IP address,
	// source port, destination IP address, destination port, and protocol—to
	// select the hash value, and then chooses the best endpoint. However, with
	// this setting, if someone uses different ports to connect to Global Accelerator,
	// their connections might not be always routed to the same endpoint because
	// the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set
	// client affinity to SOURCE_IP instead. When you use the SOURCE_IP setting,
	// Global Accelerator uses the "two-tuple" (2-tuple) properties— source (client)
	// IP address and destination IP address—to select the hash value.
	//
	// The default value is NONE.
	ClientAffinity Affinity `type:"string" enum:"true"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency—that
	// is, the uniqueness—of the request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `type:"string" required:"true"`

	// The list of port ranges to support for connections from clients to your accelerator.
	//
	// PortRanges is a required field
	PortRanges []PortRange `min:"1" type:"list" required:"true"`

	// The protocol for connections from clients to your accelerator.
	//
	// Protocol is a required field
	Protocol Protocol `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateListenerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateListenerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateListenerInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}

	if s.PortRanges == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortRanges"))
	}
	if s.PortRanges != nil && len(s.PortRanges) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortRanges", 1))
	}
	if len(s.Protocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}
	if s.PortRanges != nil {
		for i, v := range s.PortRanges {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PortRanges", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/CreateListenerResponse
type CreateListenerOutput struct {
	_ struct{} `type:"structure"`

	// The listener that you've created.
	Listener *Listener `type:"structure"`
}

// String returns the string representation
func (s CreateListenerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateListener = "CreateListener"

// CreateListenerRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Create a listener to process inbound connections from clients to an accelerator.
// Connections arrive to assigned static IP addresses on a port, port range,
// or list of port ranges that you specify. To see an AWS CLI example of creating
// a listener, scroll down to Example.
//
//    // Example sending a request using CreateListenerRequest.
//    req := client.CreateListenerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/CreateListener
func (c *Client) CreateListenerRequest(input *CreateListenerInput) CreateListenerRequest {
	op := &aws.Operation{
		Name:       opCreateListener,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateListenerInput{}
	}

	req := c.newRequest(op, input, &CreateListenerOutput{})
	return CreateListenerRequest{Request: req, Input: input, Copy: c.CreateListenerRequest}
}

// CreateListenerRequest is the request type for the
// CreateListener API operation.
type CreateListenerRequest struct {
	*aws.Request
	Input *CreateListenerInput
	Copy  func(*CreateListenerInput) CreateListenerRequest
}

// Send marshals and sends the CreateListener API request.
func (r CreateListenerRequest) Send(ctx context.Context) (*CreateListenerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateListenerResponse{
		CreateListenerOutput: r.Request.Data.(*CreateListenerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateListenerResponse is the response type for the
// CreateListener API operation.
type CreateListenerResponse struct {
	*CreateListenerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateListener request.
func (r *CreateListenerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}