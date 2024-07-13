package snowflake

var registry = make(map[string]MachineRegister)

type MachineRegister interface {
	Register(m *Machine) (*Machine, error)
	Unregister(id uint16)
}

type Machine struct {
	ID          uint16
	ServiceName string
	CreateTime  int64
}

func Register(key string, register MachineRegister) {
	registry[key] = register
}

func GetMachineRegister(key string) MachineRegister {
	return registry[key]
}
