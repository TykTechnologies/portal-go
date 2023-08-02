# Portal Go Client

A Go client for the Tyk Enterprise Dev Portal  API

[![Build Status](https://github.com/edsonmichaque/edp-go/actions/workflows/ci.yml/badge.svg)](https://github.com/dnsimple/edp-go/actions/workflows/ci.yml)

## Installation

```shell
go get github.com/edsonmichaque/edp-go
```

## Usage

This library is a Go client you can use to interact with the Tyk Enterprise Dev Portal  API. Here are some examples.

```go
package main

import (
    "context"
    "fmt"
    "os"
    "strconv"

    "github.com/edsonmichaque/edp-go"
)

func main() {
    // new client
    client, err := edp.New(
        edp.WithBaseURL("http://localhost:3000"),
        edp.WithToken("your token"),
    )

    // list organisations
    resp, err := client.Orgs().ListOrgs(context.Background())
    if err != nil {
        fmt.Printf("returned error: %v\n", err)
        os.Exit(1)
    }
}
```

## Contributing

For instructions about contributing and testing, visit the [CONTRIBUTING](CONTRIBUTING.md) file.
