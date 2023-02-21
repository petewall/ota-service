package internal

import . "github.com/petewall/firmware-service/lib"

//counterfeiter:generate . FirmwareService
type FirmwareService interface {
	GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error)
	GetFirmwareByType(firmwareType string) (FirmwareList, error)
	GetFirmwareData(firmwareType, firmwareVersion string) ([]byte, error)
}
