package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/overlordtm/uptimerobot"
)

func main() {
	app := cli.NewApp()
	app.Name = "uptimerobot"
	app.Usage = "uptimerobot cli"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "apiKey",
			Value:  "",
			Usage:  "master api key",
			EnvVar: "UPTIMEROBOT_API_KEY",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "monitor",
			Aliases: []string{"m", "mon"},
			Usage:   "monitor operations",
			Action:  getMonitors,
		},
		{
			Name:    "account",
			Aliases: []string{"a", "acct"},
			Usage:   "monitor operations",
			Action:  getAccountDetails,
		},
		{
			Name:    "contacts",
			Aliases: []string{"c", "cnts"},
			Usage:   "monitor operations",
			Action:  getAlertContacts,
		},
	}
	app.Run(os.Args)
}

func robot(c *cli.Context) *uptimerobot.UptimeRobot {
	return uptimerobot.New(c.GlobalString("apiKey"))
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
	os.Exit(1)
}

func getMonitors(c *cli.Context) {
	r := robot(c)
	mons, err := r.GetMonitors()
	if err != nil {
		handleError(err)
		return
	}

	for _, m := range mons {
		fmt.Println(m)
	}

}

func getAlertContacts(c *cli.Context) {
	r := robot(c)
	contacts, err := r.GetAlertContacts()
	if err != nil {
		handleError(err)
		return
	}

	for _, c := range contacts {
		fmt.Println(c)
	}

}

func getAccountDetails(c *cli.Context) {
	r := robot(c)
	acct, err := r.GetAccountDetails()
	if err != nil {
		handleError(err)
		return
	}

	fmt.Println(acct)

}
