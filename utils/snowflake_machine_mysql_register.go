package utils

import (
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/repositories"
	"github.com/weitien/admin/snowflake"
	"time"
)

type SnowflakeMachineMySQLRegister struct {
	Repository *repositories.SnowflakeMachineRepository
}

func init() {
	snowflake.Register("mysql", &SnowflakeMachineMySQLRegister{
		Repository: &repositories.SnowflakeMachineRepository{},
	})
}

func (s *SnowflakeMachineMySQLRegister) NewMachine(m *snowflake.Machine) (*snowflake.Machine, error) {
	machine := &models.SnowflakeMachine{
		ServiceName: m.ServiceName,
		CreateTime:  time.Now().UnixMilli(),
	}
	_, err := s.Repository.NewMachine(machine)
	if err == nil {
		m.ID = machine.ID
	}
	return m, err
}
