# goautobahn

## _Disclaimer_

I am not associated with the Federal Ministry of Transport and Digital Infrastructure or the Autobahn GmbH in any way.

Use this code at your own risk.

## Introduction

goautobahn is a library for fetching the http endpoints which are used in the official [Autobahn app](https://www.autobahn.de/app) from the Autobahn GmbH written in go.
The official API documentation is available under https://autobahn.api.bund.dev/

There is a runnable basic example for fetching all roadworks for the A1 in the cmd directory.

## Basic example

Fetching all available roads is pretty easy and can be done this way

```go
// Create a custom http client with a timeout
http := &http.Client{
    Timeout: time.Second * 15,
}

// Create a new api instance
autobahn := goautobahn.New(goautobahn.WithHTTPClient(http))

// Get all roads
resp, err := autobahn.Roads.GetAll(context.TODO())

if err != nil {
    log.Fatalln(err)
}

// Print Roadworks as JSON
respDbg, _ := json.MarshalIndent(resp, "", "	")
fmt.Println(string(respDbg))
```
