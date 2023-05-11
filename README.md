# cc-go-gin GIN

Gin Wrapper to adapt with [GoLib](https://github.com/Crosect/cc-go)

### Setup instruction

Base setup, see [GoLib Instruction](https://github.com/Crosect/cc-go/blob/main/README.md)

Both `go get` and `go mod` are supported.

```shell
go get github.com/crosect/cc-go-gin
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    "github.com/crosect/cc-go-gin"
    "go.uber.org/fx"
)

func main() {
    fx.New(
        // Using Gin as handler for Http Server,
        // Append startup method to Fx OnStart lifecycle.
        ccgin.GinHttpServerOpt(),

        // When you want to enable graceful shutdown.
        ccgin.OnStopHttpServerOpt(),
    ).Run()
}
```
