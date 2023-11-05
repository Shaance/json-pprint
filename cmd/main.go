package main

import pp "github.com/Shaance/json-pprint/v2/pprint"

func main() {
	app := &pp.App{
		OS: pp.ActualOS{},
	}
	app.Run()
}
