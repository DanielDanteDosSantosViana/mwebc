package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/DanielDanteDosSantosViana/mwebc/env"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mwebc"
	app.Usage = "Monitor HTTP Web Services Client"
	app.Commands = []cli.Command{
		{
			Name:        "init",
			Usage:       "use it to init client config",
			Description: "This function generate .mwebc file to init application",
			Action: func(c *cli.Context) {
				println("Load Configuration ", c.Args().First())
			},
		},
		{
			Name:  "auth",
			Usage: "Authentication",
			Action: func(c *cli.Context) {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("User: ")
				user, _ := reader.ReadString('\n')
				fmt.Print("Enter password: ")
				password, err := terminal.ReadPassword(0)
				if err == nil {
					fmt.Println("Password typed: "+string(password), user)
				}
			},
		},
		{
			Name:  "proxy",
			Usage: "create proxy monitor",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host, hs",
					Value: "",
					Usage: "host to proxy",
				},
				cli.StringFlag{
					Name:  "port, p",
					Value: "",
					Usage: "port to proxy",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:  "stop",
					Usage: "stop proxy",
					Action: func(c *cli.Context) error {
						fmt.Println("stop proxy...")
						enviroment := env.NewUnsetEnv()
						enviroment.UnsetGSettings()
						return nil
					},
				},
			},
			Action: func(c *cli.Context) {
				host := c.String("host")
				port := c.String("port")
				first := c.Args().First()

				if first == "stop" {
					return
				}
				if host == "" && port == "" {
					fmt.Println("[mwebc] - Error : host and port are necessary")
					fmt.Println("Plx , check de help command :' mwebc proxy -h' ")
					return
				}
				fmt.Println("start proxy...")
				enviroment := env.NewEnv(host, port)
				enviroment.GSettings()
			},
		},
		{
			Name:  "start",
			Usage: "Start send data to center",
			Action: func(c *cli.Context) {
				log.Print("Send data")
			},
		},
	}

	app.Run(os.Args)
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.

}
