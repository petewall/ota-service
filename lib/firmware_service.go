package lib

//go:generate counterfeiter -generate

type Firmware struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Size    int    `json:"size"`
	Data    []byte `json:"data"`
}

//counterfeiter:generate . FirmwareService
type FirmwareService interface {
	GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error)
	GetLatestFirmware(firmwareType string) (*Firmware, error)
}

type FirmwareServiceImpl struct {
	Host string
	Port int
}

func (f *FirmwareServiceImpl) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	return nil, nil
}

func (f *FirmwareServiceImpl) GetLatestFirmware(firmwareType string) (*Firmware, error) {
	return nil, nil
}
