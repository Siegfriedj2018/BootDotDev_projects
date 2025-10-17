# gator

Welcome to gator
This is my attempt at making a rss feed aggragator following boot.devs guided project. 
I know this is a little rough and there are probably some issues and bugs. I might try to clean it up later.
I have a list of things I want to fix/clean up. 

## Prerequisites

To use this program, you will need to install:

[Postgres](https://www.postgresql.org/download/)

[Go](https://go.dev/dl/)

Just follow the install instructions on the software links above

## Install Gator

Once you have installed Go and Postgres and have set up each, you can install gator by running the command:
go install <this-repo-url>

## Setup Config file

Once installed, you will need to setup the config file that will have the current user and postgres connection string
I have uploaded my config(.gatorconfig) as a reference for later use or as an example. I was working in wsl,so I wanted to keep everything together
for later reference. This might not work unless you fix the write config function as it save in two locations for the projects test.


