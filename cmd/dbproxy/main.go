package main

import (
	"os"

	"github.com/icewolf/cloudflare-db-proxy/dbconnect"
	"github.com/urfave/cli/v2"
)

var version = "dev"

func main() {
	app := &cli.App{}
	app.Name = "dbproxy"
	app.Version = version
	app.Usage = "Standalone implementation of Cloudflare's db-connect"
	app.UsageText = "dbproxy [global options] [command] [command options]"
	app.Commands = commands()

	_ = app.Run(os.Args)
}

func commands() []*cli.Command {
	cmds := []*cli.Command{dbconnect.Cmd()}
	return cmds
}
