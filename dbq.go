package main

import (
	"log"

	"gopkg.in/urfave/cli.v1"

	"errors"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "dbq"
	app.Usage = "Check equality of two databases."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "source-database",
			Usage: "The source db for checkin equality",
		},
		cli.StringFlag{
			Name:  "target-database",
			Usage: "The target db for checkin equality",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("source-database") == "" {
			return errors.New("Required flag missing \"source-database\"")
		}

		if c.String("target-database") == "" {
			return errors.New("Required flag missing \"target-database\"")
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
