package server

import (
	"github.com/esc-chula/intania-vote/apps/api/service"
	user "github.com/esc-chula/intania-vote/libs/grpc-go/user"
	"google.golang.org/grpc"
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
