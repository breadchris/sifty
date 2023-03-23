package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model

	Type     string `json:"type"`
	Metadata datatypes.JSON
	Data     string `json:"data"`
}
