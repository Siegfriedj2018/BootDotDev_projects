package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	comFunc, ok := c.cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("that is an invalid command, try again")
	}

	err := comFunc(s, cmd)
	return err
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.cmds == nil {
		c.cmds = make(map[string]func(*state, command) error)
	}

	c.cmds[name] = f
}
