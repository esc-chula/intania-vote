package main

import (
	"fmt"
	"log"
	"net"

	"github.com/esc-chula/intania-vote/apps/api/pkg/config"
	"github.com/esc-chula/intania-vote/apps/api/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) App {
	return App{
		cfg: cfg,
	}
}

func (a *App) Start() {
	if a.cfg == nil {
		log.Fatal("Config is not initialized")
	}

	address := fmt.Sprintf(":%d", a.cfg.Port)

	s := grpc.NewServer()
	reflection.Register(s)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
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

	logger := logger.NewLogger(a.cfg)
	defer logger.Sync()

	a.Start()

}
