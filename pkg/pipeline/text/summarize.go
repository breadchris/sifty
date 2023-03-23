package text

import (
	"context"
	"github.com/breadchris/sifty/gen/python"
	"os"
)

type Summarizer struct {
	client python.PythonClient
}

func (s *Summarizer) Summarize(fileName string) (summary string, err error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	resp, err := s.client.Summarize(context.Background(), &python.SummarizeRequest{
		Content: string(content),
	})
	if err != nil {
		return
	}
	summary = resp.Summary
	return
}

func NewSummarizer(client python.PythonClient) (*Summarizer, error) {
	return &Summarizer{
		client: client,
	}, nil
}
