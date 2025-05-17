package entity

import "time"

type VacuumStatus string

const (
	VacuumStatusDock VacuumStatus = "dock"
	VacuumStatusCleaning VacuumStatus = "cleaning"
	VacuumStatusPaused VacuumStatus = "paused"
	VacuumStatusError VacuumStatus = "error"
	VacuumStatusIdle VacuumStatus = "idle"
)

type Vaccum struct {
	ID string `json:"id"`