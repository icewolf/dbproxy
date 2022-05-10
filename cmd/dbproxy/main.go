package main

import (
	"os"

	"github.com/icewolf/cloudflare-db-proxy/dbconnect"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Name = "dbproxy"
	app.Usage = "Standalone implementation of Cloudflare's db-connect"
	app.UsageText = "dbproxy [global options] [command] [command options]"
	app.Commands = commands()

	_ = app.Run(os.Args)
}

func commands() []*cli.Command {
	cmds := []*cli.Command{dbconnect.Cmd()}
	return cmds
}
