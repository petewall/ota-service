package lib

import (
	"sort"

	"github.com/blang/semver/v4"
)

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int    `json:"size"`
	Data    []byte `json:"data"`
}

func (f *Firmware) IsOlderThan(otherFirmware *Firmware) bool {
	version, _ := semver.Make(f.Version)
	otherVersion, _ := semver.Make(otherFirmware.Version)
	return version.LT(otherVersion)
}

type FirmwareList []*Firmware

func (f FirmwareList) GetLatest(includingPrerelease bool) *Firmware {
	if len(f) == 0 {
		return nil
	}
	f.Sort()
	return f[0]
}

func (f FirmwareList) Sort() {
	sort.Slice(f, func(a, b int) bool {
		return f[a].IsOlderThan(f[b])
	})
}
