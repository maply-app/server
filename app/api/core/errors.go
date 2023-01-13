package core

import "net/http"

var RouteNotFound = ErrorResponse{
	"error",
	http.StatusNotFound,
	ErrorInfo{
		404,
		"Route not found",
	}}

var MethodNotAllowed = ErrorResponse{
	"error",
	http.StatusMethodNotAllowed,
	ErrorInfo{
		405,
		"Method not allowed",
	}}

var InternalServerError = ErrorResponse{
	"error",
	http.StatusInternalServerError,
	ErrorInfo{
		500,
		"Internal server error",
	}}

var Unauthorized = ErrorResponse{
	"error",
	http.StatusUnauthorized,
	ErrorInfo{
		401,
		"Unauthorized",
	}}

var ValidationError = ErrorResponse{
	"error",
	http.StatusBadRequest,
	ErrorInfo{
		2000,
		"Validation error",
	}}

var ObjectNotFound = ErrorResponse{
	"error",
	http.StatusBadRequest,
	ErrorInfo{
		2010,
		"Object not found",
	}}

var ObjectAlreadyExists = ErrorResponse{
	"error",
	http.StatusBadRequest,
	ErrorInfo{
		2020,
		"Object already exists",
	}}
