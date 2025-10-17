package main

import (
	"database/sql"
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/config"
	"go_tutorial/bootdotdev/blog_aggregator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db   *database.Queries
	conf *config.Config
}

func main() {
	fmt.Println("Welcome to the blog aggregator.")
	fmt.Println("Reading config...")
	currentState := &state{
		conf: config.Read(),
	}

	db, err := sql.Open("postgres", currentState.conf.DbUrl)
	if err != nil {
		log.Fatalf("there was an error in opening the database: %v", err)
	}

	defer db.Close()
	currentState.db = database.New(db)

	allCmds := commands{}
	allCmds.register("login", handlerLogin)
	allCmds.register("register", handlerRegister)
	allCmds.register("reset", handlerReset)
	allCmds.register("users", handlerGetUsers)
	allCmds.register("agg", handlerRss)
	allCmds.register("addfeed", middlewareLoggedIn(handlerFeed))
	allCmds.register("feeds", handlerFeeds)
	allCmds.register("follow", middlewareLoggedIn(handlerFollow))
	allCmds.register("following", middlewareLoggedIn(handlerFollowing))
	allCmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	allCmds.register("browse", middlewareLoggedIn(handlerBrowse))
	if len(os.Args) <= 1 {
		log.Fatalf("please provide a command. e.g. 'login <username>'")
	}
	cmdArgs := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = allCmds.run(currentState, cmdArgs)
	if err != nil {
		log.Fatalf("there was an error running the command: %v", err)
	}
}
