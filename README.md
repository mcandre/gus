# gus: clobber git submodules

# EXAMPLE

```console
$ gus -init
$ gus -add https://github.com/golang/mobile.git -path vendor/golang.org/x/mobile
$ gus -list
https://github.com/golang/mobile.git vendor/golang.org/x/mobile
$ gus -remove https://github.com/golang/mobile.git
$ gus -list

$ git status
On branch master
nothing to commit, working tree clean
```

See `gus -help` for more options.

![magic rat](https://raw.githubusercontent.com/mcandre/gus/master/gus.jpg)

# ABOUT

gus automates submodule management for safer, more reliable edits to your git repositories. This is especially helpful for managing dependencies in old-style Go 1.12- projects.

# DOWNLOAD

https://github.com/mcandre/gus/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/gus

# RUNTIME REQUIREMENTS

* [git](https://git-scm.com/)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.12+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/gus/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/gus.git $GOPATH/src/github.com/mcandre/gus
$ cd $GOPATH/src/github.com/mcandre/gus
$ git submodule update --init --recursive
$ go install ./...
```

# (INTEGRATION) TEST LOCALLY

```console
$ go test
```

# PORT

```console
$ mage port
```

# LINT

Keep the code tidy:

```console
$ mage lint
```

# DANK SOURCE MATERIAL

* [Stack Overflow](https://stackoverflow.com/questions/1260748/how-do-i-remove-a-submodule/1260982#1260982)
* [go-git](https://github.com/src-d/go-git)
* Duh, duh, duh... Happy Birthday!
