package main

import (
	"fmt"
	"log"
	"net"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/server"
	ballot "github.com/esc-chula/intania-vote/libs/grpc-go/ballot"
	user "github.com/esc-chula/intania-vote/libs/grpc-go/user"
	vote "github.com/esc-chula/intania-vote/libs/grpc-go/vote"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	userServer   user.UserServiceServer
	voteServer   vote.VoteServiceServer
	ballotServer ballot.BallotServiceServer

	cfg *config.Config
}

func NewApp(userServer user.UserServiceServer, voteServer vote.VoteServiceServer, ballotServer ballot.BallotServiceServer, cfg *config.Config) App {
	return App{
		userServer:   userServer,
		voteServer:   voteServer,
		ballotServer: ballotServer,
		cfg:          cfg,
	}
}

func (a *App) Start() {
	s := grpc.NewServer()

	server.RegisterUserServiceServer(s, a.userServer)
	server.RegisterVoteServiceServer(s, a.voteServer)
	server.RegisterBallotServiceServer(s, a.ballotServer)
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
