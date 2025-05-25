package entity

import "time"

type VaccumID int64
type VacuumStatus string

const (
	VacuumStatusDock     VacuumStatus = "dock"
	VacuumStatusCleaning VacuumStatus = "cleaning"
	VacuumStatusPaused   VacuumStatus = "paused"
	VacuumStatusError    VacuumStatus = "error"
	VacuumStatusIdle     VacuumStatus = "idle"
)

type Vacuum struct {
	ID        string       `json:"id"`
	Status    VacuumStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type VacuumRequest struct {
	EntityID string `json:"entity_id" validate:"required"`
}

type VacuumResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type VacuumCommand struct {
	Command string `json:"command" validate:"required,oneof=start pause return"`
}

type Vacuum_command []*Vacuum
