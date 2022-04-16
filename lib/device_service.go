package lib

//go:generate counterfeiter -generate

type Device struct {
	MAC               string `json:"mac" redis:"mac"`
	CurrentFirmware   string `json:"currentFirmware" redis:"currentFirmware"`
	CurrentVersion    string `json:"currentVersion" redis:"currentVersion"`
	AssignedFirmware  string `json:"assignedFirmware" redis:"assignedFirmware"`
	AssignedVersion   string `json:"assignedVersion" redis:"assignedVersion"`
	AcceptsPrerelease bool   `json:"acceptsPrerelease" redis:"acceptsPrerelease"`
}

//counterfeiter:generate . DeviceService
type DeviceService interface {
	GetDevice(mac string) (*Device, error)
}

type DeviceServiceImpl struct {
	Host string
	Port int
}

func (d *DeviceServiceImpl) GetDevice(mac string) (*Device, error) {
	return nil, nil
}
