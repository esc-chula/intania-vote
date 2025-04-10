package server

import (
	"context"
	"fmt"

	"github.com/esc-chula/intania-vote/apps/api/service"
	grpcUser "github.com/esc-chula/intania-vote/libs/grpc-go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type userServerImpl struct {
	grpcUser.UnimplementedUserServiceServer
	svc service.UserService
}

func NewUserServer(s service.UserService) grpcUser.UserServiceServer {
	return userServerImpl{
		svc: s,
	}
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, server grpcUser.UserServiceServer) {
	grpcUser.RegisterUserServiceServer(s, server)
}

func (s userServerImpl) CreateUser(ctx context.Context, req *grpcUser.CreateUserRequest) (*grpcUser.CreateUserResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	studentId := req.GetStudentId()
	if studentId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing studentId")
	}

	user, err := s.svc.GetUserByOidcIdOrStudentId(ctx, oidcId, studentId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Error(codes.Internal, "failed to get user by oidcId")
		}
	}

	if user != nil {
		if user.OidcId.String() != oidcId {
			if err := s.svc.UpdateUserById(ctx, user.Id, oidcId, studentId); err != nil {
				return nil, status.Error(codes.Internal, "failed to update user")
			}
		}
		return &grpcUser.CreateUserResponse{}, nil
	}

	if err := s.svc.CreateUser(ctx, oidcId, studentId); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &grpcUser.CreateUserResponse{}, nil
}

func (s userServerImpl) GetUserByOidcId(ctx context.Context, req *grpcUser.GetUserByOidcIdRequest) (*grpcUser.GetUserByOidcIdResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}

	user, err := s.svc.GetUserByOidcIdOrStudentId(ctx, oidcId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user by oidcId")
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &grpcUser.GetUserByOidcIdResponse{
		Id:        fmt.Sprintf("%d", user.Id),
		StudentId: user.StudentId,
		OidcId:    user.OidcId.String(),
	}, nil
}
