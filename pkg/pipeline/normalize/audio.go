package normalize

import (
	"context"
	"github.com/breadchris/sifty/gen/python"
)

type AudioNormalizer struct {
	client python.PythonClient
}

type Normalizer interface {
	Normalize(fileName string) (transcript string, err error)
}

func (s *AudioNormalizer) Normalize(fileName string) (transcript string, err error) {
	resp, err := s.client.Transcribe(context.Background(), &python.TranscribeRequest{
		File: fileName,
	})
	if err != nil {
		return
	}
	transcript = resp.Transcription
	return
}

func NewAudioNormalizer(client python.PythonClient) (*AudioNormalizer, error) {
	return &AudioNormalizer{
		client: client,
	}, nil
}
