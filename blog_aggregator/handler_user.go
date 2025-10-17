package main

import (
	"context"
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("a username is required, usage: %v <name>", cmd.Name)
	}
  
	ctx := context.Background()
	cmdName := cmd.Args[0]
  user, err := s.db.GetUser(ctx, cmdName)
	if err != nil {
		log.Fatalf("user is not found: %v", err)
		return err
	}

	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	log.Printf("User %v has been logged in\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("name is required, usage: 'register <name>'")
	}
	
	nameReg := cmd.Args[0]
	ctx := context.Background()
	params := database.CreateUserParams{
		ID:					uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:				nameReg,
	}
	
	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("there was an error creating the user: %v", err)
	}

	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("there was an error setting the user, %w", err)
	}
	fmt.Printf("User created successfully. User: %v\n", user.Name)
	return nil
}

func handlerReset(s *state, _ command) error {
	ctx := context.Background()

	err := s.db.DeleteUsers(ctx)
	if err != nil {
		log.Fatalf("there was an error deleting users: %v", err)
	}

	err = s.conf.SetUser("")
	if err != nil {
		return err
	}
	log.Println("Removed current user")
	log.Println("All tables were reset")
	return nil
}

func handlerGetUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatalf("there was an error getting users: %v", err)
	}

	for _, user := range users {
		if user.Name == s.conf.CurrentUser {
			fmt.Printf(" * %v (current)\n", user.Name)
		} else {
			fmt.Printf(" - %v\n", user.Name)
		}
	}

	return nil
}