# Isdayoff

Requests client for [Isdayoff API](https://isdayoff.ru/)

## Requirements

Go 1.13+
## Install

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference isdayoff module in a Go program with `import`:

``` go
import (
    "github.com/rageofgods/isdayoff"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch module automatically.

Alternatively, you can also explicitly `go get` the package:

```
go get -u github.com/rageofgods/isdayoff
```

## Example

```go
package main

import (
	"fmt"
	"github.com/rageofgods/isdayoff"
)

package main

import (
"fmt"
"github.com/rageofgods/isdayoff"
"time"
)

func main() {
	dayOff := isdayoff.New()
	countryCode := isdayoff.CountryCodeRussia
	pre := false
	covid := false
	tn := time.Now()
	ay, by := tn.Year()
	am, bm := tn.Month()
	ad := tn.Day()
	bd := tn.AddDate(0, 0, 10).Day()
	day, err := dayOff.GetByRange(isdayoff.ParamsRange{
		AfterYear: &ay,
		AfterMonth: &am,
		AfterDay: &ad,
		BeforeYear: &by,
		BeforeMonth: &bm,
		BeforeDay: &bd,
		Params: isdayoff.Params{
			CountryCode: &countryCode,
			Pre:         &pre,
			Covid:       &covid,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(day) // 0
}
```

## Note: 
- TZ names should be taken from [IANA](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
