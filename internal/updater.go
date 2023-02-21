package internal

import (
	"fmt"

	d "github.com/petewall/device-service/lib"
)

//counterfeiter:generate . Updater
type Updater interface {
	Update(mac, currentType, currentVersion string) ([]byte, error)
}

type UpdaterImpl struct {
	DeviceService   DeviceService
	FirmwareService FirmwareService
}

func (u *UpdaterImpl) Update(mac, currentType, currentVersion string) ([]byte, error) {
	if mac == "" {
		return nil, fmt.Errorf("mac not set")
	}

	if currentType == "" {
		return nil, fmt.Errorf("firmware not set")
	}

	device, err := u.DeviceService.GetDevice(mac)
	if err != nil {
		return nil, fmt.Errorf("unable to get device: %w", err)
	}

	if device == nil {
		device = &d.Device{}
	}

	if device.IsDifferent(currentType, currentVersion) {
		device.Firmware = currentType
		device.Version = currentVersion
		err = u.DeviceService.UpdateDevice(mac, currentType, currentVersion)
		if err != nil {
			return nil, fmt.Errorf("unable to update device: %w", err)
		}
	}

	if device.AssignedFirmware != "" {
		if device.AssignedVersion != "" {
			// Pinned version
			if device.IsDifferent(device.AssignedFirmware, device.AssignedVersion) {
				firmware, err := u.FirmwareService.GetFirmware(device.AssignedFirmware, device.AssignedVersion)
				if err != nil {
					return nil, fmt.Errorf("unable to get firmware %s %s: %w", device.AssignedFirmware, device.AssignedVersion, err)
				}

				return u.FirmwareService.GetFirmwareData(firmware.Type, firmware.Version)
			} else {
				return nil, nil
			}
		}

		// Floating version
		firmwareList, err := u.FirmwareService.GetFirmwareByType(device.AssignedFirmware)
		if err != nil {
			return nil, fmt.Errorf("unable to get firmware for type %s: %w", device.AssignedFirmware, err)
		}

		firmware := firmwareList.GetLatest(false)
		if device.IsOlderThan(firmware.Version) {
			return u.FirmwareService.GetFirmwareData(firmware.Type, firmware.Version)
		}
	}

	return nil, nil
}
