Coinbase Go Client V3
============================

## Notes
Coinbase is going to release a golang V3 client but I didn't want to wait so I wrote my own... it should work. Let me know. It's working for me =).

## Dependencies
```sh
make install
```

## Generate Mocks
```sh
make generate
```

## Installation
If using Go modules (Go version >= 1.19) simply import as needed.
```sh
go get github.com/happilymarrieddad/coinbase-go-client-v3
```

### Setup
```go
import (
    "fmt"
    "context"
    "net/http"

    coinbasegoclientv3 "github.com/happilymarrieddad/coinbase-go-client-v3"
)

func main() {
    ctx := context.Background()
    httpClient = &http.Client{
        Timeout: time.Second * 30,
    }

    client, _ = coinbasegoclientv3.NewClient(
        httpClient, "", 
        os.Getenv("COINBASE_TEST_API_KEY"), 
        os.Getenv("COINBASE_TEST_API_SECRET"),
    )
    client.HostURL = "some other URI if you want" // at the moment there is no sandbox but in case there ever is...
    ...

    productCandles, _ := client.GetProductCandles(
        ctx, "some-product-id",
        fmt.Sprintf("%d", time.Now().Add(time.Hour*-12).Unix()),
        fmt.Sprintf("%d", time.Now().Unix()),
        coinbasegoclientv3.FiveMinuteGranularity,
    )
    // Do something with productCandles
}
```

Take a look at the tests for more examples

## Tests
```sh
make test.integ
make test.unit
```

## Notes about tests
WARNING!!! Running the INTEG tests as they are at the moment run against the REAL coinbase API. Do not run them unless you have thoroughly gone through the tests, know how they work and are okay with live tests!!!!

## Build examples
```go
//go:build integ
// +build integ
```
