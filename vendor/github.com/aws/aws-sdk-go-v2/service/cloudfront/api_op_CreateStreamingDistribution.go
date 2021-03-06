// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to create a new streaming distribution.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateStreamingDistributionRequest
type CreateStreamingDistributionInput struct {
	_ struct{} `type:"structure" payload:"StreamingDistributionConfig"`

	// The streaming distribution's configuration information.
	//
	// StreamingDistributionConfig is a required field
	StreamingDistributionConfig *StreamingDistributionConfig `locationName:"StreamingDistributionConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`
}

// String returns the string representation
func (s CreateStreamingDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStreamingDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStreamingDistributionInput"}

	if s.StreamingDistributionConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamingDistributionConfig"))
	}
	if s.StreamingDistributionConfig != nil {
		if err := s.StreamingDistributionConfig.Validate(); err != nil {
			invalidParams.AddNested("StreamingDistributionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamingDistributionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.StreamingDistributionConfig != nil {
		v := s.StreamingDistributionConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "StreamingDistributionConfig", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateStreamingDistributionResult
type CreateStreamingDistributionOutput struct {
	_ struct{} `type:"structure" payload:"StreamingDistribution"`

	// The current version of the streaming distribution created.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the new streaming distribution resource just created.
	// For example: https://cloudfront.amazonaws.com/2010-11-01/streaming-distribution/EGTXBD79H29TRA8.
	Location *string `location:"header" locationName:"Location" type:"string"`

	// The streaming distribution's information.
	StreamingDistribution *StreamingDistribution `type:"structure"`
}

// String returns the string representation
func (s CreateStreamingDistributionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamingDistributionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.StreamingDistribution != nil {
		v := s.StreamingDistribution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "StreamingDistribution", v, metadata)
	}
	return nil
}

const opCreateStreamingDistribution = "CreateStreamingDistribution2019_03_26"

// CreateStreamingDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Creates a new RTMP distribution. An RTMP distribution is similar to a web
// distribution, but an RTMP distribution streams media files using the Adobe
// Real-Time Messaging Protocol (RTMP) instead of serving files using HTTP.
//
// To create a new distribution, submit a POST request to the CloudFront API
// version/distribution resource. The request body must include a document with
// a StreamingDistributionConfig element. The response echoes the StreamingDistributionConfig
// element and returns other information about the RTMP distribution.
//
// To get the status of your request, use the GET StreamingDistribution API
// action. When the value of Enabled is true and the value of Status is Deployed,
// your distribution is ready. A distribution usually deploys in less than 15
// minutes.
//
// For more information about web distributions, see Working with RTMP Distributions
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-rtmp.html)
// in the Amazon CloudFront Developer Guide.
//
// Beginning with the 2012-05-05 version of the CloudFront API, we made substantial
// changes to the format of the XML document that you include in the request
// body when you create or update a web distribution or an RTMP distribution,
// and when you invalidate objects. With previous versions of the API, we discovered
// that it was too easy to accidentally delete one or more values for an element
// that accepts multiple values, for example, CNAMEs and trusted signers. Our
// changes for the 2012-05-05 release are intended to prevent these accidental
// deletions and to notify you when there's a mismatch between the number of
// values you say you're specifying in the Quantity element and the number of
// values specified.
//
//    // Example sending a request using CreateStreamingDistributionRequest.
//    req := client.CreateStreamingDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateStreamingDistribution
func (c *Client) CreateStreamingDistributionRequest(input *CreateStreamingDistributionInput) CreateStreamingDistributionRequest {
	op := &aws.Operation{
		Name:       opCreateStreamingDistribution,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/streaming-distribution",
	}

	if input == nil {
		input = &CreateStreamingDistributionInput{}
	}

	req := c.newRequest(op, input, &CreateStreamingDistributionOutput{})
	return CreateStreamingDistributionRequest{Request: req, Input: input, Copy: c.CreateStreamingDistributionRequest}
}

// CreateStreamingDistributionRequest is the request type for the
// CreateStreamingDistribution API operation.
type CreateStreamingDistributionRequest struct {
	*aws.Request
	Input *CreateStreamingDistributionInput
	Copy  func(*CreateStreamingDistributionInput) CreateStreamingDistributionRequest
}

// Send marshals and sends the CreateStreamingDistribution API request.
func (r CreateStreamingDistributionRequest) Send(ctx context.Context) (*CreateStreamingDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamingDistributionResponse{
		CreateStreamingDistributionOutput: r.Request.Data.(*CreateStreamingDistributionOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamingDistributionResponse is the response type for the
// CreateStreamingDistribution API operation.
type CreateStreamingDistributionResponse struct {
	*CreateStreamingDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStreamingDistribution request.
func (r *CreateStreamingDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
