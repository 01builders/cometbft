package client

import "context"

type BlockAPIServiceClient interface{}

type disabledBlockAPIServiceClient struct{}

func newDisabledBlockAPIServiceClient() BlockAPIServiceClient {
	return &disabledBlockAPIServiceClient{}
}

// GetBlockByHeight implements BlockAPIServiceClient GetBlockByHeight - disabled client.
func (*disabledBlockAPIServiceClient) GetBlockByHeight(ctx context.Context, height int64) (*Block, error) {
	panic("block api service client is disabled")
}
