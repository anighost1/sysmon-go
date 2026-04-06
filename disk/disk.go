package disk

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskStats struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

func GetDiskData() (*DiskStats, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &DiskStats{
		Path:        usage.Path,
		Total:       usage.Total,
		Used:        usage.Used,
		Free:        usage.Free,
		UsedPercent: usage.UsedPercent,
	}, nil
}
