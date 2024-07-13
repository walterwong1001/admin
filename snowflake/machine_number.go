package snowflake

var registry = make(map[string]MachineRegister)

type MachineRegister interface {
	NewMachine(m *Machine) (*Machine, error)
	//register(registry map[string]MachineRegister)
}

type Machine struct {
	ID          uint16
	ServiceName string
	CreateTime  int64
}

func Register(key string, register MachineRegister) {
	registry[key] = register
}

func getMachineRegister(key string) MachineRegister {
	return registry[key]
}

func (Machine) TableName() string {
	return "snowflake_machine"
}
