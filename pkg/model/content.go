package model

import (
	"gorm.io/datatypes"
)

type Content struct {
	Base

	Type              int32 `json:"type"`
	Metadata          datatypes.JSON
	Data              string              `json:"data"`
	NormalizedContent []NormalizedContent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
