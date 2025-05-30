// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * User Service API
 *
 * API of User microservice
 *
 * API version: 0.0.1
 */

package userApi

import (
	"context"
	"net/http"
	"errors"
)

// UserAPIService is a service that implements the logic for the UserAPIServicer
// This service should implement the business logic for every endpoint for the UserAPI API.
// Include any external packages or services that will be required by this service.
type UserAPIService struct {
}

// NewUserAPIService creates a default api service
func NewUserAPIService() *UserAPIService {
	return &UserAPIService{}
}

// GetAllUsers - Get all users.
func (s *UserAPIService) GetAllUsers(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetAllUsers with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []User{}) or use other options such as http.Ok ...
	// return Response(200, []User{}), nil

	// TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	// return Response(422, nil),nil

	// TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	// return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAllUsers method not implemented")
}

// CreateUser - Create User.
func (s *UserAPIService) CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (ImplResponse, error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	// return Response(422, nil),nil

	// TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	// return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateUser method not implemented")
}

// GetUserById - Get User.
func (s *UserAPIService) GetUserById(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update GetUserById with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	// return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserById method not implemented")
}

// UpdateUserById - Update User.
func (s *UserAPIService) UpdateUserById(ctx context.Context, userId string, updateUserRequest UpdateUserRequest) (ImplResponse, error) {
	// TODO - update UpdateUserById with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	// return Response(422, nil),nil

	// TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	// return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateUserById method not implemented")
}

// DeleteUserById - Deletes a user
func (s *UserAPIService) DeleteUserById(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update DeleteUserById with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	// return Response(422, nil),nil

	// TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	// return Response(0, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUserById method not implemented")
}
