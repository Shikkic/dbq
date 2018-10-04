package main

import (
	"fmt"
	"log"

	"github.com/shikkic/dbq/equality"

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
			Name:  "source-database-url",
			Usage: "The source db url for checkin equality",
		},
		cli.StringFlag{
			Name:  "target-database-url",
			Usage: "The target db url for checkin equality",
		},
	}

	app.Action = func(c *cli.Context) error {
		sourceDBUrl := c.String("source-database-url")
		if sourceDBUrl == "" {
			return errors.New("Required flag missing \"source-database-url\"")
		}

		targetDBUrl := c.String("target-database-url")
		if targetDBUrl == "" {
			return errors.New("Required flag missing \"target-database-url\"")
		}

		fmt.Println(sourceDBUrl, targetDBUrl)

		return equality.CheckDBSubsetEquality(sourceDBUrl, targetDBUrl)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
