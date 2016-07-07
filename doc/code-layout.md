code layout
-----------

- `api`
  - lays out all the message api structures.
  - no external imports -- easy to link.
- `gat`
  - implements the `api` package interfaces.
  - calls libgit2 functions.  this is the real mccoy.
- `rpc/provider`
  - router to receive API messages, call git functions, and return API messages in response
    - actually it just shells out to anything that implements the `api` package interfaces, but it's prooobably the local one
- `rpc/client`
  - implements the `api` package interfaces.  turns the calls into messages that you should send to `rpc/provider`.
  - minimal external imports (just `api` and some serialization libraries) -- should be easy to link if you'd like to build a go process that shells out to danggit.
- `cmd/danggit`
  - the main method that builds the danggit command.
