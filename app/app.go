package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generates the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP and Name Server searching"
	app.Usage = "Searches for IPs and Name Servers on the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Looks for IPs on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server-names",
			Usage:  "Looks for server names on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
