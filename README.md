danggit
=======

`import danggit`: What you do after you got burned one too many times exec-wrapping git.

`danggit` is a pre-built, statically-linked, easily-redistributable, distro-agnostic build of libgit2.
It's distributed as a full executable.
Ship easily.

Interact with `danggit` via a simple JSON (or CBOR) RPC API.
Enjoy clean and clear error codes -- `danggit` API messages simplify and standardize far beyond exec'ing git and attempting to parse exit codes and stdout/stderr.

`danggit` can be used as a daemon, responding to many commands before exiting.
If you execute many git processes that are doing small amounts of work, `danggit` may be more performant because it eschews creating new processes every time.

`danggit` can be easily run with low privileges, in a sandbox, defanged by seccomp, or whatever you can dream up.
Since it's an executable instead of a shared library linked into your program, privilege dropping is easy.


Why
---

↑↑↑ that's why
