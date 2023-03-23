package db

import (
	"encoding/json"
	"github.com/breadchris/sifty/pkg/model"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	ProviderSet = wire.NewSet(
		NewConfig,
		NewDB,
	)
)

type Store interface {
	SaveContent(contentType, data string, metadata json.RawMessage) (uint, error)
}

type store struct {
	db *gorm.DB
}

func NewDB() (*store, error) {
	db, err := connect()
	if err != nil {
		return nil, errors.WithMessagef(err, "could not connect to database: %v", err)
	}
	return &store{db: db}, nil
}

func (s *store) SaveContent(contentType, data string, metadata json.RawMessage) (uint, error) {
	content := model.Content{
		Type:     contentType,
		Data:     data,
		Metadata: datatypes.JSON(metadata),
	}
	res := s.db.Create(&content)

	if res.Error != nil {
		return 0, errors.WithMessagef(res.Error, "could not save content: %v", res.Error)
	}
	return content.ID, nil
}
