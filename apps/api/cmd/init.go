package main

import (
	"fmt"
	"os"

	"github.com/esc-chula/intania-vote/apps/api/config"
	"github.com/esc-chula/intania-vote/apps/api/database"
	"github.com/esc-chula/intania-vote/apps/api/model"
	"github.com/esc-chula/intania-vote/apps/api/repository"
	"github.com/esc-chula/intania-vote/apps/api/server"
	"github.com/esc-chula/intania-vote/apps/api/service"
)

func InitializeApp() (App, error) {
	os.Setenv("TZ", "Asia/Bangkok")

	cfg, err := config.LoadConfig()
	if err != nil {
		return App{}, fmt.Errorf("failed to load config: %w", err)
	}

	gormDB, err := database.NewGormDatabase(cfg)
	if err != nil {
		return App{}, err
	}

	if err := gormDB.Exec(`
		DO $$ 
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'owner') THEN
				CREATE TYPE owner AS ENUM ('USER', 'ESC', 'ISESC');
			END IF;
		END $$;
	`).Error; err != nil {
		return App{}, fmt.Errorf("failed to create enum type: %w", err)
	}

	gormDB.AutoMigrate(model.User{}, model.Vote{}, model.Choice{}, model.Ballot{})

	userRepository := repository.NewUserRepository(gormDB)
	voteRepository := repository.NewVoteRepository(gormDB)
	choiceRepository := repository.NewChoiceRepository(gormDB)
	ballotRepository := repository.NewBallotRepository(gormDB)
	userService := service.NewUserService(userRepository, cfg)
	voteService := service.NewVoteService(voteRepository, userRepository, ballotRepository, cfg)
	ballotService := service.NewBallotService(ballotRepository, userRepository, voteRepository, choiceRepository, cfg)

	userServiceServer := server.NewUserServer(userService)
	voteServiceServer := server.NewVoteServer(voteService, ballotService)
	ballotServiceServer := server.NewBallotServer(ballotService)

	return NewApp(userServiceServer, voteServiceServer, ballotServiceServer, cfg), nil
}
