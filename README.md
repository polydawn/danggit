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


Building
--------

Don't.  The whole point is to use a binary release and get on with your life.

But if you're developing with us...

### dependencies

You need:
- a go compiler
- gcc (for cgo to link libgit2 and the other c bits)
- repeatr (for containers, for building the parts)
- a smattering of dev headers on your host: 'zlib' and 'dl' (these are bugs waiting to be fixed).

Having a cc toolchain on your host is kind of onerous, even, so we'll be adding more containers in the future
for the final build & link stage as well.  But for now this is it.

### steps

```
goad init ## fetch our submodules and do other first-time init.
goad parts ## build the parts.  this uses containers heavily, and will require root.
goad ## build the whole thing and run tests.
./bin/danggit ## huzzah!
```

If you're only altering the golang components, from here you'll only need to run `goad`.
Run `goad final` to make a build with the race detector disabled (you want this in production builds).
