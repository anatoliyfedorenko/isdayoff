# Isdayoff

Requests for [Isdayoff API](https://isdayoff.ru/)

## Install

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference isdayoff module in a Go program with `import`:

``` go
import (
    "github.com/anatoliyfedorenko/isdayoff"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```
go get -u github.com/anatoliyfedorenko/isdayoff
```

## Note: 
- TZ names should be taken from [IANA](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)