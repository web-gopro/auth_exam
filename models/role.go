package models

import (
	"time"

	"github.com/google/uuid"
)

// Role represents the roles table structure
type Roles struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Status    string     `json:"status" gorm:"type:varchar(50);not null" validate:"required,oneof=active deleted"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null;unique" validate:"required,min=3,max=255"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy *uuid.UUID `json:"created_by" gorm:"type:uuid"`
}

// CreateRoleRequest represents the request structure for creating a role
type CreateRoleRequest struct {
	Name   string `json:"name" validate:"required,min=3,max=255"`
	Status string `json:"status" validate:"required,oneof=active deleted"`
}

// UpdateRoleRequest represents the request structure for updating a role
type UpdateRoleRequest struct {
	Name   string `json:"name" validate:"omitempty,min=3,max=255"`
	Status string `json:"status" validate:"omitempty,oneof=active deleted"`
	ID     string `json:"id"`
}

// RoleResponse represents the response structure for role operations
type RoleResponse struct {
	ID        uuid.UUID  `json:"id"`
	Status    string     `json:"status"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
}
