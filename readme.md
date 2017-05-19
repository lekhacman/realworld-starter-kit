# ![RealWorld Example App](logo.png)

> ### Go + Gorilla codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld-example-apps) spec and API.


### [Demo]()&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)


This codebase was created to demonstrate a fully fledged fullstack application built with **Go + Gorilla** including CRUD operations, authentication, routing, pagination, and more.

We've gone to great lengths to adhere to the **Go** community styleguides & best practices.

For more information on how to this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.

## Other Go implementations
There are other GO implementations utilising the different web frameworks available to **Golang** developers.
+ [Go net/http (https://github.com/JackyChiu/realworld-starter-kit)](https://github.com/JackyChiu/realworld-starter-kit)
+ [Go + GIN (https://github.com/chrislewispac/realworld-starter-kit)](https://github.com/chrislewispac/realworld-starter-kit)

# Progress
API route status
- [x] Authentication (`POST /api/users/login`)
- [x] Registration (`POST /api/users`)
- [x] Get Current User (`GET /api/user`)
- [x] Update User (`PUT /api/user`)
- [x] Get Profile (`GET/api/profiles/:username`)
- [x] Follow user (`POST /api/profiles/:username/follow`)
- [x] Unfollow user (`DELETE /api/profiles/:username/follow`)
- [x] List Articles (`GET /api/articles`)
- [x] Feed Articles (`GET /api/articles/feed`)
- [x] Get Article (`GET /api/articles/:slug`)
- [x] Create Article (`POST /api/articles`)
- [x] Update Article (`PUT /api/articles/:slug`)
- [x] Delete Article (`DELETE /api/articles/:slug`)
- [x] Add Comments to an Article (`POST /api/articles/:slug/comments`)
- [x] Get Comments from an Article (`GET /api/articles/:slug/comments`)
- [x] Delete Comment (`DELETE /api/articles/:slug/comments/:id`)
- [x] Favourite Article (`POST /api/articles/:slug/favorite`)
- [x] Unfavourite Article (`DELETE /api/articles/:slug/favorite`)
- [x] Get Tags (`GET /api/tags`)

## Post API completion
- [ ] TESTS !!!
- [ ] Make simpler - maybe another branch to show simpler use of ``func (w http.ResponseWriter, r *http.Request) ``?
- [ ] Better/simpler error / response handling - just ``w.WriteHeader`` on error
- [ ] alice/negroni chaining?
- [ ] Optimise / Refactor / Data model Refining
- [ ] Docker?

# How it works

> TODO

# Getting started

### Installing and setting up Go
- Installation instructions for Go: [Getting Started](https://golang.org/doc/install)
- Getting familiar with the environment: [How to write Go Code](https://golang.org/doc/code.html)

### Getting the Project
```
go get github.com/chilledoj/realworld-starter-kit
cd $GOPATH/src/github.com/chilledoj/realworld-starter-kit
```

### Installing dependencies
TODO: Vendor dependencies
```
go get ./..
```

### Database (MYSQL)
Currently this codebase is utilising MYSQL as the database. Ensure you have created the database and tables (``mysql_init.sql``  provided in models directory)

### Building and Running
You can build using your own build flags, or use ``MAKE`` as there is a ``makefile`` provided.
```
make build
```
A run command is provided in the ``makefile`` which will first run the build step.
```
make run
```
**N.B.** The database URL (containing user+pass+host+port) can be set as an environment variable, or passed as a flag to the executable
```
./bin/backend -dburl "root:password@/conduit?parseTime=true"
// or
DBURL=root:password@/conduit?parseTime=true ./bin/backend
```

Additionally the IP host and ports are currently defaulted to run on ":8080", but can be passed in as ENV variables or command line flags ``HOST or -host`` and ``PORT or -port``.

# Contributing
Raise an issue and/or Pull Request. Alternatively create your own better fork.
