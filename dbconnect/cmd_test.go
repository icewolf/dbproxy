package dbconnect

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/urfave/cli/v2"
)

func TestCmd(t *testing.T) {
	tests := [][]string{
		{"dbproxy", "start", "--port", "1234", "--playground"},
		{"dbproxy", "start", "--port", "1234", "--url", "sqlite3::memory:?cache=shared", "--insecure"},
		{"dbproxy", "start", "--port", "1234", "--url", "sqlite3::memory:?cache=shared", "--auth-domain", "mysite.cloudflareaccess.com", "--application-aud", "aud"},
	}

	cmd := Cmd()
	cmd.Action = func(*cli.Context) error { return nil }

	app := &cli.App{
		Name:     "dbproxy",
		Commands: []*cli.Command{cmd},
	}

	for _, test := range tests {
		assert.NoError(t, app.Run(test))
	}
}
