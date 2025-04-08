package main

import (
	"fmt"

	"github.com/esc-chula/intania-vote/apps/api/pkg/config"
	"github.com/esc-chula/intania-vote/apps/api/pkg/database"
)

func InitializeApp() (App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return App{}, fmt.Errorf("failed to load config: %w", err)
	}

	_, err = database.NewGormDatabase(cfg)
	if err != nil {
		return App{}, err
	}

	return NewApp(cfg), nil
}
