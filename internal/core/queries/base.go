package queries

import (
	"context"
)

type BaseQueryHandler struct {
}

type BaseQueryRequest struct {
	Input string
}

type BaseQueryResponse struct {
	Input string
}

func (h BaseQueryHandler) Handle(_ context.Context, _ *BaseQueryRequest) (*BaseQueryResponse, error) {
	return &BaseQueryResponse{}, nil
}
