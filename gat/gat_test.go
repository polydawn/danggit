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

var execgitcommitheaders = gosh.Opts{
	Env: map[string]string{
		"GIT_AUTHOR_NAME":     "author",
		"GIT_AUTHOR_EMAIL":    "email@domain.wow",
		"GIT_AUTHOR_DATE":     "17 Oct 2009 12:00:00 -0000",
		"GIT_COMMITTER_NAME":  "author",
		"GIT_COMMITTER_EMAIL": "email@domain.wow",
		"GIT_COMMITTER_DATE":  "17 Oct 2009 12:00:00 -0000",
	},
}
