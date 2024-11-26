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

To use the *Advanced Trade Go SDK*, initialize the [Credentials](credentials.go) struct and create a new client. The Credentials struct is JSON
enabled. Ensure that Advanced Trade API credentials are stored in a secure manner.

```
credentials := &credentials.Credentials{}
if err := json.Unmarshal([]byte(os.Getenv("ADV_CREDENTIALS")), credentials); err != nil {
    return nil, fmt.Errorf("unable to deserialize advanced trade credentials JSON: %w", err)
}

httpClient, err := client.DefaultHttpClient()
if err != nil {
    panic(fmt.Sprintf("unable to load default http client: %v", err))
}

restClient := client.NewRestClient(credentials, httpClient)
```

There are convenience functions to read the credentials as an environment variable (credentials.ReadEnvCredentials) and to deserialize the JSON structure (credentials.UnmarshalCredentials) if pulled from a different source. The JSON format expected by both is:

```
{
  "accessKey": "",
  "privatePemKey": ""
}
```

Coinbase Advanced Trade API credentials can be created in the [CDP web portal](https://portal.cdp.coinbase.com/). 

Once the client is initialized, initialize a service to make the desired call. For example, to [list portfolios](https://github.com/coinbase-samples/advanced-trade-sdk-go/blob/main/list_portfolios.go),
create the service, pass in the request object, check for an error, and if nil, process the response.


```
service := portfolios.NewPortfoliosService(restClient)


response, err := service.ListPortfolios(ctx, &portfolios.ListPortfoliosRequest{})
```

## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build ./...
```
