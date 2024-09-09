package handlers

import "github.com/walterwong1001/admin/internal/machine"

func NextId() uint64 {
	return machine.NextID()
}
