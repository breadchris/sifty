//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type SelfserviceSettingsFlows struct {
	ID              uuid.UUID `sql:"primary_key"`
	RequestURL      string
	IssuedAt        time.Time
	ExpiresAt       time.Time
	IdentityID      uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ActiveMethod    *string
	State           string
	Type            string
	UI              *string
	Nid             *uuid.UUID
	InternalContext string
}
