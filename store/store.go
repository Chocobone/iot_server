package store

import (
	"errors"

	"github.com/Chocobone/iot_server/entity"
)

var (
	Vacuum_command = &VacuumStore{Vacuum_command: map[entity.VacuumID]*entity.Vacuum{}}

	ErrNotFound = errors.New("not found")
)

type VacuumStore struct {
	Vacuum_command map[entity.VacuumID]*entity.Vacuum
}

func (vs *VacuumStore) Get(id entity.VacuumID) (*entity.Vacuum, error) {

}
