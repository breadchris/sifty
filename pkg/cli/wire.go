//go:build wireinject
// +build wireinject

package cli

import (
	"github.com/breadchris/sifty/pkg/api"
	"github.com/breadchris/sifty/pkg/config"
	"github.com/breadchris/sifty/pkg/pipeline/normalize"
	"github.com/breadchris/sifty/pkg/pipeline/text"
	"github.com/breadchris/sifty/pkg/python"
	"github.com/breadchris/sifty/pkg/store"
	"github.com/breadchris/sifty/pkg/store/db"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func Wire() (*cli.App, error) {
	panic(wire.Build(
		NewApp,
		config.ProviderSet,
		api.ProviderSet,
		python.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		normalize.NewAudioNormalizer,
		text.NewSummarizer,
	))
}
