package app_error

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppError = error

var (
	ErrConfigInvalidPort  AppError = errors.New("invalid port")
	ErrUnauthorized       AppError = errors.New("unauthorized")
	ErrInvalidAccountType AppError = errors.New("invalid account type")
	ErrResourceNotFound   AppError = errors.New("resource not found")
)

var ErrGrpcInternal = status.Error(codes.Internal, "internal server error")
