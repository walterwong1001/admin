package snowflake

import (
	"github.com/sony/sonyflake"
	"sync"
	"time"
)

var once sync.Once
var snowflake *sonyflake.Sonyflake

func InitSnowFlake(key, serviceName string) uint16 {
	var machineId uint16
	once.Do(func() {
		var st sonyflake.Settings
		register := GetMachineRegister(key)

		st.MachineID = func() (uint16, error) {
			machine, err := register.Register(&Machine{
				ServiceName: serviceName,
				CreateTime:  time.Now().UnixMilli(),
			})
			machineId = machine.ID
			return machine.ID, err
		}
		snowflake = sonyflake.NewSonyflake(st)
	})
	return machineId
}

func GetSnowflake() *sonyflake.Sonyflake {
	return snowflake
}
