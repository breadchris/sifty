package api

import (
	"context"
	"github.com/twitchtv/twirp"
	"log"
)

func NewLoggingServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			method, _ := twirp.MethodName(ctx)
			log.Println("Method: " + method)
			return ctx, nil
		},
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			log.Println("Error: " + string(twerr.Msg()))
			return ctx
		},
		ResponseSent: func(ctx context.Context) {
			log.Println("Response Sent (error or success)")
		},
	}
}
