package snowflake

import (
	"github.com/sony/sonyflake"
	"sync"
	"time"
)

var once sync.Once
var snowflake *sonyflake.Sonyflake

func InitSnowFlake(key, serviceName string) {
	once.Do(func() {
		var st sonyflake.Settings
		register := getMachineRegister(key)

		st.MachineID = func() (uint16, error) {
			machine, err := register.NewMachine(&Machine{
				ServiceName: serviceName,
				CreateTime:  time.Now().UnixMilli(),
			})
			return machine.ID, err
		}
		snowflake = sonyflake.NewSonyflake(st)
	})
}

func GetSnowflake() *sonyflake.Sonyflake {
	return snowflake
}
