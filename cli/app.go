package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type App struct {
	OS OSInterface
}

func (a *App) Run() {
	const inputFileFlagName = "file"
	const outputFileFlagName = "out"
	var useSpaces bool
	var filePath string
	var outFile string
	var writeToFile bool

	app := &cli.App{
		Name:                   "json-pprint",
		Usage:                  "Json to pretty print",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "write",
				Aliases:     []string{"w"},
				Usage:       "write result to file instead of stdout, by default overwrites source file",
				Value:       false,
				Destination: &writeToFile,
			},
			&cli.BoolFlag{
				Name:        "spaces",
				Aliases:     []string{"s"},
				Usage:       "use 2 spaces instead of tab",
				Value:       false,
				Destination: &useSpaces,
			},
			&cli.StringFlag{
				Name:        inputFileFlagName,
				Usage:       "file to read JSON from",
				Aliases:     []string{"f"},
				TakesFile:   true,
				Destination: &filePath,
			},
			&cli.StringFlag{
				Name:        outputFileFlagName,
				Aliases:     []string{"o"},
				Usage:       "output file for result",
				TakesFile:   true,
				Destination: &outFile,
			},
		},
		Action: func(cCtx *cli.Context) error {
			parsedString, err := retrieveJsonInput(cCtx.Args().First(), cCtx.String(inputFileFlagName), a.OS)

			if err != nil {
				fmt.Printf("Error while trying to retrieve json input: %s", err)
				cli.Exit(err, 1)
			}

			if (len(parsedString) == 0) {
				cli.ShowAppHelpAndExit(cCtx, 0)
			}

			prettyJson, err := indentJson(parsedString, useSpaces)
			if err != nil {
				fmt.Printf("Error while trying to indent json: %s\n", err)
				cli.Exit(err, 1)
			}

			return writeOutput(prettyJson, writeToFile, cCtx.String(inputFileFlagName), cCtx.String(outputFileFlagName), a.OS)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeOutput(content string, writeToFile bool, inputFilePath string, outputFilePath string, OS OSInterface) error {
	var writer io.Writer
	if writeToFile {
		if outputFilePath == "" {
			outputFilePath = inputFilePath
		}

		f, err := OS.Create(outputFilePath)
		if err != nil {
			return err
		}
		writer = f
		defer f.Close()
	} else {
		writer = os.Stdout
	}

	fmt.Fprintln(writer, content)
	return nil
}

func indentJson(rawJson string, useSpaces bool) (string, error) {
	indent := "\t" // using tab indentation by default
	if useSpaces {
		indent = "  "
	}

	const prefix = ""
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(rawJson), prefix, indent); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

func retrieveJsonInput(firstArg string, filePath string, _os OSInterface) (string, error) {
	var usingArgument = filePath == ""

	if usingArgument {
		return firstArg, nil
	}

	// using file option
	data, err := _os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
