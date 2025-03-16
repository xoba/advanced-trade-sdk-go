# Bitcoin Market Sell Example

This example demonstrates how to sell Bitcoin at market price using the Coinbase Advanced Trade SDK.

## Setup

1. Set up your Coinbase Advanced Trade API credentials at https://portal.cdp.coinbase.com/

2. Export your credentials as environment variables:

   ```bash
   # Export your API key name
   export KEY_NAME="your-key-name"
   
   # Export your private key
   export KEY_VALUE="-----BEGIN EC PRIVATE KEY-----
   MIHcAgEBBEIBH0Cj+...your entire private key here...
   ...more lines of your key...
   -----END EC PRIVATE KEY-----"
   ```

   Tip: For multiline private keys, you can use this approach to preserve the format:
   ```bash
   export KEY_VALUE=$(cat << 'EOF'
   -----BEGIN EC PRIVATE KEY-----
   MIHcAgEBBEIBH0Cj+...your entire private key here...
   ...more lines of your key...
   -----END EC PRIVATE KEY-----
   EOF
   )
   ```

3. Build the example:
   ```bash
   go build .
   ```

## Usage

You can sell Bitcoin in two ways:

1. Specify the amount of BTC to sell (maximum 0.01 BTC):
   ```bash
   ./example1 -btc 0.01
   ```

2. Specify the USD value of BTC to sell (maximum $1000):
   ```bash
   ./example1 -usd 1000
   ```

3. You can also change the trading pair (default is BTC-USD):
   ```bash
   ./example1 -btc 0.005 -product ETH-USD
   ```

## What the Example Does

1. Loads your Coinbase Advanced Trade API credentials from the environment variable
2. Creates a preview of the order to show estimated proceeds and fees
3. Asks for confirmation before proceeding
4. If confirmed, places a market order to sell the specified amount of cryptocurrency
5. Displays the order result

## Important Notes

- Make sure you have sufficient funds in your account for the specified transaction.
- Always review the order preview before confirming.
- The "MarketMarketIoc" order type is used, which is an Immediate-or-Cancel market order.
- The API keys used must have permission to trade.