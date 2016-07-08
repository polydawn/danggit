package gat

import (
	"github.com/polydawn/gosh"
)

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

var execgit gosh.Command = gosh.Gosh(
	"git",
	gosh.NullIO,
	gosh.Opts{
		Env: map[string]string{
			"GIT_CONFIG_NOSYSTEM": "true",
			"HOME":                "/dev/null",
			"GIT_ASKPASS":         "/bin/true",
		},
	},
)
