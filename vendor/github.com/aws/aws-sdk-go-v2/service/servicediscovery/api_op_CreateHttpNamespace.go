// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an HTTP namespace. Service instances registered using an HTTP namespace
// can be discovered using a DiscoverInstances request but can't be discovered
// using DNS.
//
// For the current quota on the number of namespaces that you can create using the
// same Amazon Web Services account, see [Cloud Map quotas]in the Cloud Map Developer Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
func (c *Client) CreateHttpNamespace(ctx context.Context, params *CreateHttpNamespaceInput, optFns ...func(*Options)) (*CreateHttpNamespaceOutput, error) {
	if params == nil {
		params = &CreateHttpNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHttpNamespace", params, optFns, c.addOperationCreateHttpNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHttpNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHttpNamespaceInput struct {

	// The name that you want to assign to this namespace.
	//
	// This member is required.
	Name *string

	// A unique string that identifies the request and that allows failed
	// CreateHttpNamespace requests to be retried without the risk of running the
	// operation twice. CreatorRequestId can be any unique string (for example, a
	// date/time stamp).
	CreatorRequestId *string

	// A description for the namespace.
	Description *string

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value that you define. Tags keys can be up to 128 characters in length, and tag
	// values can be up to 256 characters in length.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateHttpNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see [GetOperation].
	//
	// [GetOperation]: https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHttpNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHttpNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHttpNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHttpNamespace"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateHttpNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateHttpNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHttpNamespace(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateHttpNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateHttpNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateHttpNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateHttpNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateHttpNamespaceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateHttpNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateHttpNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateHttpNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHttpNamespace",
	}
}
