package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/credentials"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/advanced-trade-sdk-go/orders"
	"github.com/google/uuid"
)

func main() {
	// Define command line flags
	var btcAmount float64
	var usdAmount float64
	var productId string
	var useBtc, useUsd bool

	flag.Float64Var(&btcAmount, "btc", 0, "Amount of BTC to sell (e.g., 0.01)")
	flag.Float64Var(&usdAmount, "usd", 0, "USD value of BTC to sell (e.g., 1000)")
	flag.StringVar(&productId, "product", "BTC-USD", "Product ID to trade (default: BTC-USD)")
	flag.Parse()

	// Determine which flags were actually set by user
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "btc" {
			useBtc = true
		}
		if f.Name == "usd" {
			useUsd = true
		}
	})

	// Validate flags
	if !useBtc && !useUsd {
		fmt.Println("Error: You must specify either -btc or -usd amount")
		fmt.Println("Usage examples:")
		fmt.Println("  ./example1 -btc 0.01      # Sell 0.01 BTC (maximum allowed)")
		fmt.Println("  ./example1 -usd 1000      # Sell $1000 worth of BTC (maximum allowed)")
		os.Exit(1)
	}

	if useBtc && useUsd {
		fmt.Println("Error: Specify either -btc or -usd, not both")
		os.Exit(1)
	}

	// Safety limits
	const MAX_BTC_AMOUNT = 0.01
	const MAX_USD_AMOUNT = 1000.0

	// Check BTC safety limit
	if useBtc {
		if btcAmount <= 0 {
			fmt.Println("Error: BTC amount must be greater than 0")
			os.Exit(1)
		}
		if btcAmount > MAX_BTC_AMOUNT {
			fmt.Printf("Error: BTC amount exceeds safety limit of %.5f BTC\n", MAX_BTC_AMOUNT)
			os.Exit(1)
		}
	}

	// Check USD safety limit
	if useUsd {
		if usdAmount <= 0 {
			fmt.Println("Error: USD amount must be greater than 0")
			os.Exit(1)
		}
		if usdAmount > MAX_USD_AMOUNT {
			fmt.Printf("Error: USD amount exceeds safety limit of $%.2f\n", MAX_USD_AMOUNT)
			os.Exit(1)
		}
	}

	// Load credentials from environment variables
	keyName := os.Getenv("KEY_NAME")
	if keyName == "" {
		log.Fatalf("KEY_NAME environment variable not set")
	}

	privateKey := os.Getenv("KEY_VALUE")
	if privateKey == "" {
		log.Fatalf("KEY_VALUE environment variable not set")
	}

	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	// Create credentials manually
	creds := &credentials.Credentials{
		AccessKey:     keyName,
		PrivatePemKey: privateKey,
	}

	// Print confirmation of credentials loaded
	fmt.Printf("Credentials loaded from environment variables - Key name: %s, Private key: %d bytes\n",
		creds.AccessKey, len(creds.PrivatePemKey))

	// Create HTTP client
	httpClient, err := client.DefaultHttpClient()
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create REST client
	restClient := client.NewRestClient(creds, httpClient)

	// Create orders service
	ordersService := orders.NewOrdersService(restClient)

	// Build order configuration based on specified amount type
	orderConfig := model.OrderConfiguration{
		MarketMarketIoc: &model.MarketIoc{},
	}

	var orderTypeDesc string
	if useBtc {
		// For BTC amount, directly use BaseSize (already correct)
		// Format BTC amount to 8 decimal places (Bitcoin's standard precision)
		btcSizeStr := fmt.Sprintf("%.8f", btcAmount)
		orderConfig.MarketMarketIoc.BaseSize = btcSizeStr
		orderTypeDesc = fmt.Sprintf("%s BTC", btcSizeStr)
	} else {
		// For USD amount with SELL, we need to convert to base currency first
		// We'll preview the order to get the equivalent BTC amount for the USD value
		
		// Create a temporary preview request with QuoteSize (for estimation only)
		usdSizeStr := fmt.Sprintf("%.2f", usdAmount)
		tempPreviewRequest := &orders.CreateOrderPreviewRequest{
			ProductId: productId,
			Side:      "SELL",
			OrderConfiguration: model.OrderConfiguration{
				MarketMarketIoc: &model.MarketIoc{
					QuoteSize: usdSizeStr,
				},
			},
		}

		fmt.Printf("Getting BTC equivalent for $%s...\n", usdSizeStr)
		previewResp, err := ordersService.CreateOrderPreview(context.Background(), tempPreviewRequest)
		if err != nil {
			log.Fatalf("Failed to get BTC equivalent: %v", err)
		}

		// Extract the base size (BTC amount) from the preview
		btcEquivalent := previewResp.BaseSize
		if btcEquivalent == "" {
			log.Fatalf("Failed to get BTC equivalent from preview")
		}

		// Round the BTC amount to 8 decimal places (Bitcoin's standard precision)
		btcFloat, err := strconv.ParseFloat(btcEquivalent, 64)
		if err != nil {
			log.Fatalf("Failed to parse BTC equivalent: %v", err)
		}
		
		// Format to 8 decimal places
		btcRounded := fmt.Sprintf("%.8f", btcFloat)
		
		fmt.Printf("$%s is approximately %s BTC at current market price\n", usdSizeStr, btcRounded)
		
		// Now use the rounded BaseSize for the actual order
		orderConfig.MarketMarketIoc.BaseSize = btcRounded
		orderTypeDesc = fmt.Sprintf("%s BTC (equivalent to $%s)", btcRounded, usdSizeStr)
	}

	// Optional: Preview the order first to see estimated cost/fees
	previewRequest := &orders.CreateOrderPreviewRequest{
		ProductId:          productId,
		Side:               "SELL",
		OrderConfiguration: orderConfig,
	}

	fmt.Printf("Previewing order to sell %s...\n", orderTypeDesc)
	previewResponse, err := ordersService.CreateOrderPreview(context.Background(), previewRequest)
	if err != nil {
		log.Fatalf("Order preview failed: %v", err)
	}

	// Print preview details
	previewJson, _ := json.MarshalIndent(previewResponse, "", "  ")
	fmt.Printf("Order Preview:\n%s\n\n", previewJson)

	// Ask for confirmation before proceeding
	fmt.Print("Proceed with order? (y/n): ")
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Println("Order canceled")
		os.Exit(0)
	}

	// Create and submit the actual order
	orderRequest := &orders.CreateOrderRequest{
		ProductId:          productId,
		Side:               "SELL",
		ClientOrderId:      uuid.New().String(),
		OrderConfiguration: orderConfig,
	}

	fmt.Println("Placing order...")
	orderResponse, err := ordersService.CreateOrder(context.Background(), orderRequest)
	if err != nil {
		log.Fatalf("Order creation failed: %v", err)
	}

	// Check order response and print details
	if orderResponse.Success {
		fmt.Println("Order successfully placed!")
		fmt.Printf("Order ID: %s\n", orderResponse.SuccessResponse.OrderId)

		// Print the full response
		responseJson, _ := json.MarshalIndent(orderResponse, "", "  ")
		fmt.Printf("Response:\n%s\n", responseJson)
	} else {
		fmt.Printf("Order failed: %s\n", orderResponse.FailureReason)
		orderJson, _ := json.MarshalIndent(orderResponse, "", "  ")
		fmt.Printf("Error details:\n%s\n", orderJson)
	}
}
