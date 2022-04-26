package lib

import (
	"fmt"
)

//go:generate counterfeiter -generate

//counterfeiter:generate . Updater
type Updater interface {
	Update(mac string, currentFirmware *Firmware) (*Firmware, error)
}

type UpdaterImpl struct {
	DeviceService   DeviceService
	FirmwareService FirmwareService
}

func (u *UpdaterImpl) Update(mac string, currentFirmware *Firmware) (*Firmware, error) {
	if mac == "" {
		return nil, fmt.Errorf("mac not set")
	}

	if currentFirmware == nil {
		return nil, fmt.Errorf("firmware not set")
	}

	device, err := u.DeviceService.GetDevice(mac)
	if err != nil {
		return nil, fmt.Errorf("unable to get device: %w", err)
	}

	if device == nil {
		device = &Device{}
	}

	if device.IsDifferent(currentFirmware) {
		device.CurrentFirmware = currentFirmware.Type
		device.CurrentVersion = currentFirmware.Version
		err = u.DeviceService.UpdateDevice(device)
		if err != nil {
			return nil, fmt.Errorf("unable to update device: %w", err)
		}
	}

	if device.AssignedFirmware != "" {
		if device.AssignedVersion != "" {
			// Pinned version
			assigned := &Firmware{
				Type:    device.AssignedFirmware,
				Version: device.AssignedVersion,
			}
			if device.IsDifferent(assigned) {
				firmware, err := u.FirmwareService.GetFirmware(device.AssignedFirmware, device.AssignedVersion)
				if err != nil {
					return nil, fmt.Errorf("unable to get firmware: %w", err)
				}

				return firmware, nil
			} else {
				return nil, nil
			}
		}

		// Floating version
		firmware, err := u.FirmwareService.GetLatestFirmware(device.AssignedFirmware, device.AcceptsPrerelease)
		if err != nil {
			return nil, fmt.Errorf("unable to get latest firmware: %w", err)
		}

		if device.IsOlderThan(firmware) {
			return firmware, nil
		}
	}

	return nil, nil
}
