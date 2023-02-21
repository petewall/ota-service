package internal

import (
	. "github.com/petewall/device-service/lib"
)

//counterfeiter:generate . DeviceService
type DeviceService interface {
	GetDevice(mac string) (*Device, error)
	UpdateDevice(mac, firmwareType, firmwareVersion string) error
}
