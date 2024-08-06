package machine

var registry = make(map[string]SnowflakeMachineRegister)

type SnowflakeMachineRegister interface {
	Register(m *Machine) (*Machine, error)
	Unregister(id uint16)
}

type Machine struct {
	ID          uint16
	ServiceName string
	CreateTime  int64
}

func Register(key string, register SnowflakeMachineRegister) {
	registry[key] = register
}

func GetMachineRegister(key string) SnowflakeMachineRegister {
	return registry[key]
}
