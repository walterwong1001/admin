package repositories

import (
	"github.com/weitien/admin/models"
)

type SnowflakeMachineRepository struct{}

func (s *SnowflakeMachineRepository) Add(m *models.SnowflakeMachine) (*models.SnowflakeMachine, error) {
	err := db.Create(m).Error
	return m, err
}

func (s *SnowflakeMachineRepository) Remove(id uint16) {
	db.Delete(&models.SnowflakeMachine{}, id)
}
