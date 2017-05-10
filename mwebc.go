package main

import (
	"bufio"
	"fmt"
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
			Name:  "init",
			Usage: "Load configuration",
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
			Action: func(c *cli.Context) {
				enviroment := env.NewEnv("127.0.0.1", "8080")
				//enviroment.GSettings()
				enviroment.UnsetGSettings()
			},
		},
	}

	app.Run(os.Args)
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.

}
