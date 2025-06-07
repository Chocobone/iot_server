package entity

import "time"

type VacuumID int64
type VacuumToken string
type VacuumStatus string

const (
	VacuumStatusDock     VacuumStatus = "dock"
	VacuumStatusCleaning VacuumStatus = "cleaning"
	VacuumStatusPaused   VacuumStatus = "paused"
	VacuumStatusError    VacuumStatus = "error"
	VacuumStatusIdle     VacuumStatus = "idle"
)

type Vacuum struct {
	ID        VacuumID     `json:"id"`
	Token     VacuumToken  `json:"vacuum_token"`
	Status    VacuumStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
