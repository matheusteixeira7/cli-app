package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "cli app"
	app.Usage = "Find ips and domains"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devnine.tech",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find ips",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "domain",
			Usage:  "Find domains",
			Flags:  flags,
			Action: findDomains,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findDomains(c *cli.Context) {
	host := c.String("host")

	domains, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range domains {
		fmt.Println(domain.Host)
	}
}
