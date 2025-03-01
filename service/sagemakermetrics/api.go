// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakermetrics

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opBatchPutMetrics = "BatchPutMetrics"

// BatchPutMetricsRequest generates a "aws/request.Request" representing the
// client's request for the BatchPutMetrics operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See BatchPutMetrics for more information on using the BatchPutMetrics
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the BatchPutMetricsRequest method.
//	req, resp := client.BatchPutMetricsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sagemaker-metrics-2022-09-30/BatchPutMetrics
func (c *SageMakerMetrics) BatchPutMetricsRequest(input *BatchPutMetricsInput) (req *request.Request, output *BatchPutMetricsOutput) {
	op := &request.Operation{
		Name:       opBatchPutMetrics,
		HTTPMethod: "PUT",
		HTTPPath:   "/BatchPutMetrics",
	}

	if input == nil {
		input = &BatchPutMetricsInput{}
	}

	output = &BatchPutMetricsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// BatchPutMetrics API operation for Amazon SageMaker Metrics Service.
//
// Used to ingest training metrics into SageMaker which can be visualized in
// SageMaker Studio and retrieved with the GetMetrics API.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon SageMaker Metrics Service's
// API operation BatchPutMetrics for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sagemaker-metrics-2022-09-30/BatchPutMetrics
func (c *SageMakerMetrics) BatchPutMetrics(input *BatchPutMetricsInput) (*BatchPutMetricsOutput, error) {
	req, out := c.BatchPutMetricsRequest(input)
	return out, req.Send()
}

// BatchPutMetricsWithContext is the same as BatchPutMetrics with the addition of
// the ability to pass a context and additional request options.
//
// See BatchPutMetrics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SageMakerMetrics) BatchPutMetricsWithContext(ctx aws.Context, input *BatchPutMetricsInput, opts ...request.Option) (*BatchPutMetricsOutput, error) {
	req, out := c.BatchPutMetricsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// An error that occured when putting the metric data.
type BatchPutMetricsError struct {
	_ struct{} `type:"structure"`

	// The error code of an error that occured when attempting to put metrics.
	//
	//    * METRIC_LIMIT_EXCEEDED - The max amount of metrics per resource has been
	//    exceeded.
	//
	//    * INTERNAL_ERROR - An internal error occured.
	//
	//    * VALIDATION_ERROR - The metric data failed validation.
	//
	//    * CONFLICT_ERROR - Multiple requests attempted to modify the same data
	//    simultaneously.
	Code *string `type:"string" enum:"PutMetricsErrorCode"`

	// An index that corresponds to the metric in the request.
	MetricIndex *int64 `type:"integer"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsError) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsError) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *BatchPutMetricsError) SetCode(v string) *BatchPutMetricsError {
	s.Code = &v
	return s
}

// SetMetricIndex sets the MetricIndex field's value.
func (s *BatchPutMetricsError) SetMetricIndex(v int64) *BatchPutMetricsError {
	s.MetricIndex = &v
	return s
}

type BatchPutMetricsInput struct {
	_ struct{} `type:"structure"`

	// A list of raw metric values to put.
	//
	// MetricData is a required field
	MetricData []*RawMetricData `min:"1" type:"list" required:"true"`

	// The name of Trial Component to associate the metrics with.
	//
	// TrialComponentName is a required field
	TrialComponentName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutMetricsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "BatchPutMetricsInput"}
	if s.MetricData == nil {
		invalidParams.Add(request.NewErrParamRequired("MetricData"))
	}
	if s.MetricData != nil && len(s.MetricData) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MetricData", 1))
	}
	if s.TrialComponentName == nil {
		invalidParams.Add(request.NewErrParamRequired("TrialComponentName"))
	}
	if s.TrialComponentName != nil && len(*s.TrialComponentName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TrialComponentName", 1))
	}
	if s.MetricData != nil {
		for i, v := range s.MetricData {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MetricData", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMetricData sets the MetricData field's value.
func (s *BatchPutMetricsInput) SetMetricData(v []*RawMetricData) *BatchPutMetricsInput {
	s.MetricData = v
	return s
}

// SetTrialComponentName sets the TrialComponentName field's value.
func (s *BatchPutMetricsInput) SetTrialComponentName(v string) *BatchPutMetricsInput {
	s.TrialComponentName = &v
	return s
}

type BatchPutMetricsOutput struct {
	_ struct{} `type:"structure"`

	// Any errors that occur when inserting metric data will appear in this.
	Errors []*BatchPutMetricsError `min:"1" type:"list"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s BatchPutMetricsOutput) GoString() string {
	return s.String()
}

// SetErrors sets the Errors field's value.
func (s *BatchPutMetricsOutput) SetErrors(v []*BatchPutMetricsError) *BatchPutMetricsOutput {
	s.Errors = v
	return s
}

// The raw metric data to associate with the resource.
type RawMetricData struct {
	_ struct{} `type:"structure"`

	// The name of the metric.
	//
	// MetricName is a required field
	MetricName *string `min:"1" type:"string" required:"true"`

	// Metric step (aka Epoch).
	Step *int64 `type:"integer"`

	// The time when the metric was recorded.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" required:"true"`

	// The metric value.
	//
	// Value is a required field
	Value *float64 `type:"double" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RawMetricData) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RawMetricData) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RawMetricData) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RawMetricData"}
	if s.MetricName == nil {
		invalidParams.Add(request.NewErrParamRequired("MetricName"))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MetricName", 1))
	}
	if s.Timestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("Timestamp"))
	}
	if s.Value == nil {
		invalidParams.Add(request.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMetricName sets the MetricName field's value.
func (s *RawMetricData) SetMetricName(v string) *RawMetricData {
	s.MetricName = &v
	return s
}

// SetStep sets the Step field's value.
func (s *RawMetricData) SetStep(v int64) *RawMetricData {
	s.Step = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *RawMetricData) SetTimestamp(v time.Time) *RawMetricData {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *RawMetricData) SetValue(v float64) *RawMetricData {
	s.Value = &v
	return s
}

const (
	// PutMetricsErrorCodeMetricLimitExceeded is a PutMetricsErrorCode enum value
	PutMetricsErrorCodeMetricLimitExceeded = "METRIC_LIMIT_EXCEEDED"

	// PutMetricsErrorCodeInternalError is a PutMetricsErrorCode enum value
	PutMetricsErrorCodeInternalError = "INTERNAL_ERROR"

	// PutMetricsErrorCodeValidationError is a PutMetricsErrorCode enum value
	PutMetricsErrorCodeValidationError = "VALIDATION_ERROR"

	// PutMetricsErrorCodeConflictError is a PutMetricsErrorCode enum value
	PutMetricsErrorCodeConflictError = "CONFLICT_ERROR"
)

// PutMetricsErrorCode_Values returns all elements of the PutMetricsErrorCode enum
func PutMetricsErrorCode_Values() []string {
	return []string{
		PutMetricsErrorCodeMetricLimitExceeded,
		PutMetricsErrorCodeInternalError,
		PutMetricsErrorCodeValidationError,
		PutMetricsErrorCodeConflictError,
	}
}
