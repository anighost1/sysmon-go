package memory

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

func GetMemoryData() (*MemoryStats, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &MemoryStats{
		Total:       vm.Total,
		Used:        vm.Used,
		Free:        vm.Free,
		UsedPercent: vm.UsedPercent,
	}, nil
}
