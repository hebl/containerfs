package utils

import (
	"runtime"
	"syscall"
)

// MemStatus ...
type MemStatus struct {
	All   uint64 `json:"all"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
	Usage float64
}

// MemStat ...
func MemStat() MemStatus {
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}

	//only for linux/mac
	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err == nil {
		mem.All = sysInfo.Totalram
		mem.Free = sysInfo.Freeram
		mem.Used = mem.All - mem.Free
		mem.Usage = float64(mem.Used) / float64(mem.All)
	}
	return mem
}
