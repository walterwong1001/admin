package repositories

import (
	"github.com/weitien/admin/models"
)

type SnowflakeMachineRepository struct{}

func (s *SnowflakeMachineRepository) NewMachine(m *models.SnowflakeMachine) (*models.SnowflakeMachine, error) {
	err := DB.Create(m).Error
	return m, err
}
