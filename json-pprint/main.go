package main

import "github.com/Shaance/json-pprint/v2/cli"

func main() {
	app := &cli.App{
		OS: cli.ActualOS{},
	}
	app.Run()
}
