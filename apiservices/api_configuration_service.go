/*
 * App template API
 *
 * API to access and configure the app template
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiservices

import (
	"context"
	"errors"
	"hailo/apiserver"
	"net/http"
)

// ConfigurationApiService is a service that implements the logic for the ConfigurationApiServicer
// This service should implement the business logic for every endpoint for the ConfigurationApi API.
// Include any external packages or services that will be required by this service.
type ConfigurationApiService struct {
}

// NewConfigurationApiService creates a default api service
func NewConfigurationApiService() apiserver.ConfigurationApiServicer {
	return &ConfigurationApiService{}
}

// GetExamples - Get example configuration
func (s *ConfigurationApiService) GetExamples(ctx context.Context) (apiserver.ImplResponse, error) {
	// TODO - update GetExamples with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Example{}) or use other options such as http.Ok ...
	//return Response(200, []Example{}), nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GetExamples method not implemented")
}

// PostExample - Creates an example configuration
func (s *ConfigurationApiService) PostExample(ctx context.Context, example apiserver.Example) (apiserver.ImplResponse, error) {
	// TODO - update PostExample with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Example{}) or use other options such as http.Ok ...
	//return Response(201, Example{}), nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("PostExample method not implemented")
}