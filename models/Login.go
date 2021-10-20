package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Credentials struct para armazenar o account_id e secret no corpo do request
type Credentials struct {
	ID     uuid.UUID `gorm:"type:uuid" json:"id"`
	Secret string    `json:"secret"  validate:"required"`
}

type Claims struct {
	ID uuid.UUID `gorm:"type:uuid" json:"id" validate:"required"`
	jwt.StandardClaims
}
