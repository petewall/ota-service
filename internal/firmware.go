package internal

import (
	"sort"

	"github.com/Masterminds/semver"
)

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int    `json:"size"`
	Data    []byte `json:"data"`
}

func (f *Firmware) IsOlderThan(otherFirmware *Firmware) bool {
	version, _ := semver.NewVersion(f.Version)
	otherVersion, _ := semver.NewVersion(otherFirmware.Version)
	return version.LessThan(otherVersion)
}

type FirmwareList []*Firmware

func (f FirmwareList) GetLatest(includingPrerelease bool) *Firmware {
	list := f
	if !includingPrerelease {
		list = f.FilterOutPrerelease()
	}
	if len(list) == 0 {
		return nil
	}
	list.Sort()
	return list[len(list)-1]
}

func (f FirmwareList) Sort() {
	sort.Slice(f, func(a, b int) bool {
		return f[a].IsOlderThan(f[b])
	})
}

func (f FirmwareList) FilterOutPrerelease() FirmwareList {
	var filtered FirmwareList
	for _, firmware := range f {
		version, _ := semver.NewVersion(firmware.Version)
		if version.Prerelease() == "" {
			filtered = append(filtered, firmware)
		}
	}
	return filtered
}
