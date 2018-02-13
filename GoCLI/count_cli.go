package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
)

func main(){

	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage: "Count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "stop, s",
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error{
					start := c.Int("stop")
					if start <= 0 {
						fmt.Println("Stop cannot be zero or negative.")
					}else{
						for i := 1; i <= start ; i++  {
							fmt.Println(i)
						}
					}
					return nil
			},
		},
		{
			Name: "down",
			ShortName: "d",
			Usage: "count down",
			Flags:[]cli.Flag{
				cli.IntFlag{
					Name: "start, s",
					Usage: "Value to count dow from",
					Value: 10,
				},
			},
			Action: func(c * cli.Context) error {
				start := c.Int("start")
				if start <= 0 {
					fmt.Println("Start cannont be zero or negative")
				}else{
					for ; start >= 0 ; start--  {
						fmt.Println(start)
					}
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}