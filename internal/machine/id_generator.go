package machine

import (
	"github.com/walterwong1001/snowflake/pkg/snowflake"
	"log"
	"sync"
	"time"
)

var once sync.Once
var generator *snowflake.Snowflake

func InitSnowFlake(key, serviceName string) uint16 {
	var machineId uint16
	once.Do(func() {
		register := GetMachineRegister(key)
		m, err := register.Register(&Machine{
			ServiceName: serviceName,
			CreateTime:  time.Now().UnixMilli(),
		})
		machineId = m.ID
		generator, err = snowflake.NewSnowflake(machineId)
		if err != nil {
			log.Println(err)
		}
	})
	return machineId
}

func GetSnowflake() *snowflake.Snowflake {
	return generator
}

func NextID() uint64 {
	return generator.NextID()
}
