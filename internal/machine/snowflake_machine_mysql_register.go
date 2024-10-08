package machine

import (
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
	"time"
)

type SnowflakeMachineMySQLRegister struct {
	Repository *repositories.SnowflakeMachineRepository
}

func init() {
	Register("mysql", &SnowflakeMachineMySQLRegister{
		Repository: &repositories.SnowflakeMachineRepository{},
	})
}

func (s *SnowflakeMachineMySQLRegister) Register(machine *Machine) (*Machine, error) {
	m := &models.SnowflakeMachine{
		ServiceName: machine.ServiceName,
		CreateTime:  time.Now().UnixMilli(),
	}
	_, err := s.Repository.Add(m)
	if err == nil {
		machine.ID = m.ID
	}
	return machine, err
}

func (s *SnowflakeMachineMySQLRegister) Unregister(id uint16) {
	s.Repository.Remove(id)
}
