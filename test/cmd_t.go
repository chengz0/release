package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "say"
	app.Commands = []cli.Command{
		{
			Name:        "hello",
			ShortName:   "hi",
			Usage:       "use it to see a description",
			Description: "This is how we describe hello the function",
			Subcommands: []cli.Command{
				{
					Name:        "english",
					ShortName:   "en",
					Usage:       "sends a greeting in english",
					Description: "greets someone in english",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Value: "Bob",
							Usage: "Name of the person to greet",
						},
					},
					Action: func(c *cli.Context) {
						println("Hello, ", c.String("name"))
					},
				}, {
					Name:      "spanish",
					ShortName: "sp",
					Usage:     "sends a greeting in spanish",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "surname",
							Value: "Jones",
							Usage: "Surname of the person to greet",
						},
					},
					Action: func(c *cli.Context) {
						println("Hola, ", c.String("surname"))
					},
				}, {
					Name:      "french",
					ShortName: "fr",
					Usage:     "sends a greeting in french",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "nickname",
							Value: "Stevie",
							Usage: "Nickname of the person to greet",
						},
					},
					Action: func(c *cli.Context) {
						println("Bonjour, ", c.String("nickname"))
					},
				},
			},
		}, {
			Name:  "bye",
			Usage: "says goodbye",
			Action: func(c *cli.Context) {
				println("bye")
			},
		},
	}

	app.Run(os.Args)
}
