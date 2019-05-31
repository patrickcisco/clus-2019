package main

import (
	"os"
	"time"

	"github.com/patrickcisco/clus-2019/go-curl/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-curl"
	app.Version = "1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "patrickcisco",
		},
	}
	app.Usage = "go-curl https://developer.cisco.com"
	app.Description = "go-curl is a minimal http tool"
	app.Commands = cmd.Init()
	app.Run(os.Args)

}
