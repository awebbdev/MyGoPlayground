package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"time"
)

func main(){

	var language string

	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Usage = "Print a hello world message"
	app.Authors = []cli.Author{
		{
			Name:  "Adam R. Webb",
			Email: "adamwebb13@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{	Name: "name, n",
						Value: "World",
						Usage: "Who to say hello to.",
						},
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "yell",
			Aliases: []string{"y"},
			Usage: "Yell at the user.",
			Action: func(c *cli.Context) error {
				fmt.Println("YOU MORON,", c.Args().First())
				return nil
			},
		},
	}


	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		if language == "spanish"{
			fmt.Printf("Hola %s\n", name)
		}else{
			fmt.Printf("Hello %s!\n", name)
		}
		return nil
	}

	app.Run(os.Args)
}