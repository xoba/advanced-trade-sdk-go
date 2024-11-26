/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package public

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
)

type PublicService interface {
	GetPublicMarketTrades(ctx context.Context, request *GetPublicMarketTradesRequest) (*GetPublicMarketTradesResponse, error)
	GetPublicProduct(ctx context.Context, request *GetPublicProductRequest) (*GetPublicProductResponse, error)
	GetPublicProductBook(ctx context.Context, request *GetPublicProductBookRequest) (*GetPublicProductBookResponse, error)
	GetPublicProductCandles(ctx context.Context, request *GetPublicProductCandlesRequest) (*GetPublicProductCandlesResponse, error)
	GetServerTime(ctx context.Context, request *GetServerTimeRequest) (*GetServerTimeResponse, error)
	ListPublicProducts(ctx context.Context, request *ListPublicProductsRequest) (*ListPublicProductsResponse, error)
}

func NewPublicService(c client.RestClient) PublicService {
	return &publicServiceImpl{client: c}
}

type publicServiceImpl struct {
	client client.RestClient
}
