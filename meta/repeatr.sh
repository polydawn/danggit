#!/bin/bash
##
## Source this to get `$repeatr` and `$rciflags` env vars.
##
## - `$repeatr` -- the full path command to exec.  If you have one on your
##    path, it's used; if not, we'll toolstrap one.
## - `$rciflags` -- flags you should give to `$repeatr run`.  It's mostly blank,
##    but will set some modes that make things work in CI environments (e.g.
##    places known not to be able to support all container features).
##
eval "$(
	cd "$( dirname "${BASH_SOURCE[0]}" )"
	source ./toolstrap.sh

	repeatr="`which repeatr`" || {
		toolstrap \
			repeatr v0.12 \
			61ef917c7988d985629a4818858dbc614cb7a6da6c37c2a6bcf6cf97781fc5c83f028243d4c11a2b7d958a1c78fa6c6b \
			https://github.com/polydawn/repeatr/releases/download/release%2Fv0.12/repeatr-linux-amd64-v0.12.tar.gz
		repeatr="$PWD/tools/repeatr/v0.12/repeatr"
	}

	rciflags=""
	[ -z "${TRAVIS:-}" ] || rciflags+=" --executor=chroot"

	printf 'repeatr=%q\n'  "$repeatr"
	printf 'rciflags=%q\n' "$rciflags"
)"
