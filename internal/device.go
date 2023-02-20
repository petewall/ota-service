package internal

type Device struct {
	MAC               string `json:"mac" redis:"mac"`
	CurrentFirmware   string `json:"currentFirmware" redis:"currentFirmware"`
	CurrentVersion    string `json:"currentVersion" redis:"currentVersion"`
	AssignedFirmware  string `json:"assignedFirmware" redis:"assignedFirmware"`
	AssignedVersion   string `json:"assignedVersion" redis:"assignedVersion"`
	AcceptsPrerelease bool   `json:"acceptsPrerelease" redis:"acceptsPrerelease"`
}

func (d *Device) IsDifferent(firmware *Firmware) bool {
	return d.CurrentFirmware != firmware.Type || d.CurrentVersion != firmware.Version
}

func (d *Device) IsOlderThan(firmware *Firmware) bool {
	currentFirmware := &Firmware{
		Type:    d.CurrentFirmware,
		Version: d.CurrentVersion,
	}
	return currentFirmware.IsOlderThan(firmware)
}
