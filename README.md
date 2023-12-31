# Portal Client for Go

A Go client for the Tyk Enterprise Dev Portal  API

[![Build Status](https://github.com/TykTechnologies/portal-go/actions/workflows/ci.yml/badge.svg)](https://github.com/TykTechnologies/portal-go/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/TykTechnologies/portal-go?status.svg)](https://godoc.org/github.com/TykTechnologies/portal-go) 
[![Go Report Card](https://goreportcard.com/badge/github.com/TykTechnologies/portal-go)](https://goreportcard.com/report/github.com/portal-go)

## Installation

```shell
go get github.com/TykTechnologies/portal-go
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

    "github.com/TykTechnologies/portal-go"
)

func main() {
    // new client
    client, err := portal.New(
        portal.WithBaseURL(os.Getenv("PORTAL_URL")),
        portal.WithToken(os.Getenv("PORTAL_TOKEN")),
    )

    // list organisations
    portal, err := client.Orgs().ListOrgs(context.Background())
    if err != nil {
        fmt.Printf("returned error: %v\n", err)
        os.Exit(1)
    }
}
```

## Contributing

For instructions about contributing and testing, visit the [CONTRIBUTING](CONTRIBUTING.md) file.
