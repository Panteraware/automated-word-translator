package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var translateLink string
	var translateApiKey string
	var file string
	var fromLanguage string
	var toLanguage string
	var saveLocation string

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("Version=%s Commit=%s Build Date=%s\n", version, commit, date)
	}

	app := &cli.App{
		Name:    "Automated Word Translator",
		Usage:   "Automatically translate words in the contents of a file",
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Konotorii",
				Email: "github@konotorii.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "translateLink",
				Usage:       "Link to translation server",
				Destination: &translateLink,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "translateApiKey",
				Usage:       "API Key to translation server",
				Destination: &translateApiKey,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "file", // Not actual file, just file contents
				Usage:       "Contents of the file desired to read",
				Destination: &file,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "fromLanguage",
				Value:       "de",
				Usage:       "Language to translate from",
				Destination: &fromLanguage,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "toLanguage",
				Value:       "en",
				Usage:       "Language to translate to",
				Destination: &toLanguage,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "saveLocation",
				Usage:       "File location to save the final output at",
				Destination: &saveLocation,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			replaced := file

			fileStrings := strings.Split(file, " ")
			fileStringsLength := len(fileStrings)
			var translatedStrings []map[string]string

			println(fmt.Sprintf("Found %s words to translate...", FormatNumber(fileStringsLength)))

			startTime := time.Now()

			for i := 0; i < len(fileStrings); i++ {
				translated := Translate(c, Format(fileStrings[i]), translateApiKey, fromLanguage, toLanguage)

				translatedStrings = append(translatedStrings, map[string]string{"original": fileStrings[i], "new": translated})

				replaced = strings.ReplaceAll(replaced, fileStrings[i], translated)

				println(fmt.Sprintf("%s %s of %s words translated.", Percentage(float64(i), float64(fileStringsLength)), FormatNumber(i), FormatNumber(fileStringsLength)))
			}

			println(fmt.Sprintf("Translated %d words in %s.", len(translatedStrings), time.Since(startTime).String()))

			err := os.WriteFile(saveLocation, []byte(replaced), 0644)
			if err != nil {
				log.Fatal(err)
				return err
			}

			println(fmt.Sprintf("Saved translation to: %s", saveLocation))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
