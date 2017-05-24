package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/DanielDanteDosSantosViana/mwebc/env"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	RED    = "\x1b[31m"
	GREEN  = "\x1b[32m"
	YELLOW = "\x1b[33m"
	BLUE   = "\x1b[34m"
	CYAN   = "\x1b[36m"
	PURPLE = "\x1b[35m"
	GRAY   = "\x1b[37m"
	NONE   = "\x1b[0m"
)

func color(color string) string {
	if isTTY {
		return color
	} else {
		return ""
	}
}

var isTTY bool

const successResponsePrefix = "time"

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
					Usage: "- Stop proxy server",
					Action: func(c *cli.Context) error {
						enviroment := env.NewUnsetEnv()
						err := enviroment.UnsetGSettings()
						if err != nil {
							logrus.WithFields(logrus.Fields{
								"enviroment": enviroment,
							}).Info(err)
							return err
						}
						if isTTY {
							fmt.Printf("%s[MWEBC] - %sStop proxy server\n%s", color(GREEN), color(BLUE), color(NONE))
						}
						return nil
					},
				},
			},
			Action: func(c *cli.Context) {
				host := c.String("host")
				port := c.String("port")

				if host == "" || port == "" {

					if isTTY {
						fmt.Printf("%s[MWEBC] - %shost and port are necessary, check de help command: %s' mwebc proxy -h'\n%s", color(GREEN), color(RED), color(BLUE), color(NONE))
					}

					logrus.WithFields(logrus.Fields{
						"Port ": port,
						"Host":  host,
					}).Error("host and port are necessary")
					return
				}

				if isTTY {
					fmt.Printf("%s[MWEBC] - %s Start proxy ...\n%s", color(GREEN), color(BLUE), color(NONE))
					enviroment := env.NewEnv(host, port)
					err := enviroment.GSettings()
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"Port ": port,
							"Host":  host,
						}).Error(err)
						return
					}

					log.Println(err)

					fmt.Printf("%sHost\t\t\t\t\t%sPort\t\t\t\t%sEnabled\n%s", color(GREEN), color(BLUE),
						color(YELLOW), color(NONE))
					fmt.Printf("%s======================================================================================\n", color(NONE))
					fmt.Printf("%s%s\t\t\t\t%s%s\t\t\t\t%s%s\n%s", color(GREEN), host, color(BLUE), port,
						color(YELLOW), "TRUE", color(NONE))
					Daemon()
					//go proxy.Start()
				}
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
	isTTY = terminal.IsTerminal(int(os.Stdout.Fd()))
	app.Run(os.Args)

}

func Daemon() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
