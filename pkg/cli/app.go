package cli

import (
	"github.com/breadchris/sifty/pkg/api"
	"github.com/breadchris/sifty/pkg/pipeline/normalize"
	"github.com/breadchris/sifty/pkg/pipeline/text"
	"github.com/urfave/cli/v2"
)

func NewApp(
	httpServer api.HTTPServer,
	normalizer normalize.Normalizer,
	summarizer text.Summarizer,
) *cli.App {
	return &cli.App{
		Name:  "lunabrain",
		Usage: "Save and search for information.",
		Commands: []*cli.Command{
			NewServeCommand(httpServer),
			NewSyncCommand(),
			NewCaptureCommand(normalizer),
			NewNormalizeCommand(normalizer),
			NewTextCommand(summarizer),
		},
	}
}
