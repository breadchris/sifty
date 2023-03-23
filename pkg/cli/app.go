package cli

import (
	"github.com/breadchris/sifty/pkg/pipeline/normalize"
	"github.com/breadchris/sifty/pkg/pipeline/text"
	"github.com/urfave/cli/v2"
)

func NewApp(normalizer *normalize.AudioNormalizer, summarizer *text.Summarizer) *cli.App {
	return &cli.App{
		Name:  "sifty",
		Usage: "Search for information",
		Commands: []*cli.Command{
			NewSyncCommand(),
			NewCaptureCommand(normalizer),
			NewNormalizeCommand(normalizer),
			NewTextCommand(summarizer),
		},
	}
}
