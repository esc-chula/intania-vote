package server

import (
	"context"
	"fmt"

	"github.com/esc-chula/intania-vote/apps/api/service"
	user "github.com/esc-chula/intania-vote/libs/grpc-go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type userServerImpl struct {
	user.UnimplementedUserServiceServer
	svc service.UserService
}

func NewUserServer(s service.UserService) user.UserServiceServer {
	return userServerImpl{
		svc: s,
	}
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, server user.UserServiceServer) {
	user.RegisterUserServiceServer(s, server)
}

func (s userServerImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}
	studentId := req.GetStudentId()
	if studentId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing studentId")
	}

	existedUser, err := s.svc.GetUserByOidcIdOrStudentId(ctx, oidcId, studentId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Error(codes.Internal, "failed to get user by oidcId")
		}
	}

	if existedUser != nil {
		if existedUser.OidcId.String() != oidcId {
			if err := s.svc.UpdateUserById(ctx, existedUser.Id, oidcId, studentId); err != nil {
				return nil, status.Error(codes.Internal, "failed to update user")
			}
		}
		return &user.CreateUserResponse{}, nil
	}

	if err := s.svc.CreateUser(ctx, oidcId, studentId); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &user.CreateUserResponse{}, nil
}

func (s userServerImpl) GetUserByOidcId(ctx context.Context, req *user.GetUserByOidcIdRequest) (*user.GetUserByOidcIdResponse, error) {
	oidcId := req.GetOidcId()
	if oidcId == "" {
		return nil, status.Error(codes.FailedPrecondition, "missing oidcId")
	}

	foundedUser, err := s.svc.GetUserByOidcIdOrStudentId(ctx, oidcId, "")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user by oidcId")
	}
	if foundedUser == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &user.GetUserByOidcIdResponse{
		Id:        fmt.Sprintf("%d", foundedUser.Id),
		StudentId: foundedUser.StudentId,
		OidcId:    foundedUser.OidcId.String(),
	}, nil
}
