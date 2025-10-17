package main

import (
	"context"
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUser)
		if err != nil {
			return fmt.Errorf("failed to retrieve current user: %w", err)
		}

		return handler(s, c, currentUser)
	}

}