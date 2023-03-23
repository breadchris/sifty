package api

import "github.com/breadchris/sifty/gen/api"

func Serve(apiServer *APIServer) {
	api.NewAPIServer(apiServer, nil)
}
