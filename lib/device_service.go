package lib

//go:generate counterfeiter -generate

//counterfeiter:generate . DeviceService
type DeviceService interface {
	GetDevice(mac string) (*Device, error)
	UpdateDevice(device *Device) error
}

type DeviceServiceImpl struct {
	Host string
	Port int
}

func (d *DeviceServiceImpl) GetDevice(mac string) (*Device, error) {
	return nil, nil
}

func (d *DeviceServiceImpl) UpdateDevice(device *Device) error {
	return nil
}
