package api

import (
	"context"
	"github.com/breadchris/sifty/gen/api"
)

type APIServer struct {
}

func (s APIServer) Save(ctx context.Context, content *api.Content) (*api.ContentID, error) {

}

func (s APIServer) Search(ctx context.Context, query *api.Query) (*api.Results, error) {
}

func NewAPIServer() *APIServer {
	return &APIServer{}
}
