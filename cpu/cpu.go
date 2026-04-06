package cpu

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

type CoreUsage struct {
	CoreID int     `json:"core_id"`
	Usage  float64 `json:"usage"`
}

type CpuStats struct {
	PhysicalCores int         `json:"physical_cores"`
	LogicalCores  int         `json:"logical_cores"`
	TotalUsage    float64     `json:"total_usage"`
	PerCoreUsage  []CoreUsage `json:"per_core_usage"`
}

func GetCpuData() (*CpuStats, error) {
	physical, err := cpu.Counts(false)
	if err != nil {
		return nil, err
	}

	logical, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	totalArr, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	perCoreArr, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}

	perCoreUsage := make([]CoreUsage, len(perCoreArr))
	for i, v := range perCoreArr {
		perCoreUsage[i] = CoreUsage{
			CoreID: i,
			Usage:  v,
		}
	}

	return &CpuStats{
		PhysicalCores: physical,
		LogicalCores:  logical,
		TotalUsage:    totalArr[0],
		PerCoreUsage:  perCoreUsage,
	}, nil
}
