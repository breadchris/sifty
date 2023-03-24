package api

import (
	"context"
	"encoding/json"
	genapi "github.com/breadchris/sifty/gen/api"
	"github.com/breadchris/sifty/pkg/pipeline/normalize"
	"github.com/breadchris/sifty/pkg/store"
	"github.com/breadchris/sifty/pkg/store/db"
	"github.com/pkg/errors"
	"path"
)

type APIServer struct {
	fileStore  *store.Files
	db         db.Store
	normalizer *normalize.AudioNormalizer
}

func (s APIServer) Save(ctx context.Context, content *genapi.Content) (*genapi.ContentID, error) {
	data, err := json.Marshal(content.Metadata)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to marshal metadata for content %s", content.Type.String())
	}

	contentID, err := s.db.SaveContent(content.Type, string(content.Data), data)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to save content %s", content.Type.String())
	}

	if content.Type == genapi.ContentType_AUDIO {
		err = s.fileStore.Bucket.WriteAll(ctx, contentID.String(), content.Data, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to save audio content to bucket")
		}

		transcript, err := s.normalizer.Normalize(path.Join(s.fileStore.Location, contentID.String()))
		if err != nil {
			return nil, errors.Wrapf(err, "unable to normalize audio content")
		}

		_, err = s.db.SaveNormalizedContent(contentID, transcript)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to save normalized content")
		}
	}

	return &genapi.ContentID{
		Id: contentID.String(),
	}, nil
}

func (s APIServer) Search(ctx context.Context, query *genapi.Query) (*genapi.Results, error) {
	content, err := s.db.GetStoredContent()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get stored content")
	}

	var storedContent []*genapi.StoredContent
	for _, c := range content {
		normalData := ""
		if len(c.NormalizedContent) > 0 {
			normalData = c.NormalizedContent[0].Data
		}

		storedContent = append(storedContent, &genapi.StoredContent{
			Content: &genapi.Content{
				Data:      []byte(c.Data),
				Type:      genapi.ContentType(c.Type),
				Metadata:  nil,
				CreatedAt: c.CreatedAt.String(),
			},
			Normalized: &genapi.NormalizedContent{
				Data: normalData,
			},
		})
	}
	return &genapi.Results{
		StoredContent: storedContent,
	}, nil
}

func NewAPIServer(db db.Store, fileStore *store.Files, normalizer *normalize.AudioNormalizer) *APIServer {
	return &APIServer{
		db:         db,
		fileStore:  fileStore,
		normalizer: normalizer,
	}
}
