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

type SelfserviceVerificationFlows struct {
	ID           uuid.UUID `sql:"primary_key"`
	RequestURL   string
	IssuedAt     time.Time
	ExpiresAt    time.Time
	CsrfToken    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Type         string
	State        string
	ActiveMethod *string
	UI           *string
	Nid          *uuid.UUID
}
