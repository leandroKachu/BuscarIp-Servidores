package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Gerando comando line"
	app.Usage = "Buscar ips"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "buscar ips",
			Flags:  flags,
			Action: BuscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "buscar servidores",
			Flags:  flags,
			Action: BuscarServidores,
		},
	}
	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func BuscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
