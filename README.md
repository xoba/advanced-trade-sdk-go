# Advanced Trade Go SDK README

[![GoDoc](https://godoc.org/github.com/coinbase-samples/advanced-trade-sdk-go?status.svg)](https://godoc.org/github.com/coinbase-samples/advanced-trade-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/coinbase-samples/advanced-trade-sdk-go)](https://goreportcard.com/report/coinbase-samples/advanced-trade-sdk-go)

## Overview

The *Advanced Trade Go SDK* is a sample library that demonstrates the structure of a [Coinbase Advanced Trade](https://www.coinbase.com/advanced-trade/) driver for
the [REST APIs](https://docs.cloud.coinbase.com/advanced-trade/docs/welcome).

Coinbase Advanced Trade offers a comprehensive API for traders, providing access to real-time market data, order management, and execution. Elevate your trading strategies and develop sophisticated solutions using our powerful tools and features.

## License

The *Advanced Trade Go SDK* sample library is free and open source and released under the [Apache License, Version 2.0](LICENSE).

The application and code are only available for demonstration purposes.

## Usage

There are multiple ways to initialize the *Advanced Trade Go SDK* with your credentials. Choose the approach that best fits your needs.

### Option 1: Using Environment Variables for Credentials JSON

You can use the JSON format in an environment variable:

```go
// Load credentials from ADV_CREDENTIALS environment variable
credentials, err := credentials.ReadEnvCredentials("ADV_CREDENTIALS")
if err != nil {
    return nil, fmt.Errorf("unable to load credentials: %w", err)
}

httpClient, err := client.DefaultHttpClient()
if err != nil {
    panic(fmt.Sprintf("unable to load default http client: %v", err))
}

restClient := client.NewRestClient(credentials, httpClient)
```

The JSON format expected is:

```json
{
  "accessKey": "your-api-key-name",
  "privatePemKey": "-----BEGIN EC PRIVATE KEY-----\n...your private key...\n-----END EC PRIVATE KEY-----"
}
```

### Option 2: Direct Credential Input

You can also create the credentials struct directly:

```go
// Get credentials from environment variables or any other source
keyName := os.Getenv("KEY_NAME")
privateKey := os.Getenv("KEY_VALUE")

// Create credentials manually
credentials := &credentials.Credentials{
    AccessKey:     keyName,
    PrivatePemKey: privateKey,
}

httpClient, err := client.DefaultHttpClient()
if err != nil {
    panic(fmt.Sprintf("unable to load default http client: %v", err))
}

restClient := client.NewRestClient(credentials, httpClient)
```

### Getting Started

Coinbase Advanced Trade API credentials can be created in the [CDP web portal](https://portal.cdp.coinbase.com/).

Once the client is initialized, create a service for the specific API you want to use. For example, to [list portfolios](https://github.com/coinbase-samples/advanced-trade-sdk-go/blob/main/portfolios/list_portfolios.go):

```go
service := portfolios.NewPortfoliosService(restClient)

response, err := service.ListPortfolios(ctx, &portfolios.ListPortfoliosRequest{})
```

### Examples

Check out the [example1](./example1) directory for a complete working example of selling Bitcoin at market price.

## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build ./...
```
