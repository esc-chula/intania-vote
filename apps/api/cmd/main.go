package main

import (
	"fmt"
	"log"
	"net"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/server"
	user "github.com/esc-chula/intania-vote/libs/grpc-go/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	userServer user.UserServiceServer

	cfg *config.Config
}

func NewApp(userServer user.UserServiceServer, cfg *config.Config) App {
	return App{
		userServer: userServer,
		cfg:        cfg,
	}
}

func (a *App) Start() {
	s := grpc.NewServer()

	server.RegisterUserServiceServer(s, a.userServer)
	server.RegisterHealthCheckServer(s)

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	a, err := InitializeApp()
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	a.Start()
}
