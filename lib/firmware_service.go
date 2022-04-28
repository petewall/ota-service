package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//go:generate counterfeiter -generate

//counterfeiter:generate . FirmwareService
type FirmwareService interface {
	GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error)
	GetLatestFirmware(firmwareType string, includePrerelease bool) (*Firmware, error)
}

type FirmwareServiceImpl struct {
	Host       string
	Port       int
	HTTPClient HTTPClient
}

func (f *FirmwareServiceImpl) GetFirmware(firmwareType, firmwareVersion string) (*Firmware, error) {
	resp, err := f.HTTPClient.Get(fmt.Sprintf("http://%s:%d/%s/%s", f.Host, f.Port, firmwareType, firmwareVersion))
	if err != nil {
		return nil, fmt.Errorf("failed to get firmware %s %s: %w", firmwareType, firmwareVersion, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read firmware %s %s response: %w", firmwareType, firmwareVersion, err)
	}

	var firmware *Firmware
	err = json.Unmarshal(body, &firmware)
	if err != nil {
		return nil, fmt.Errorf("failed to parse firmware %s %s response: %w", firmwareType, firmwareVersion, err)
	}

	return firmware, nil
}

func (f *FirmwareServiceImpl) GetLatestFirmware(firmwareType string, includePrerelease bool) (*Firmware, error) {
	resp, err := f.HTTPClient.Get(fmt.Sprintf("http://%s:%d/%s", f.Host, f.Port, firmwareType))
	if err != nil {
		return nil, fmt.Errorf("failed to get list of firmware %s: %w", firmwareType, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read list of firmware %s response: %w", firmwareType, err)
	}

	var firmwareList FirmwareList
	err = json.Unmarshal(body, &firmwareList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse list of firmware %s response: %w", firmwareType, err)
	}

	return firmwareList.GetLatest(includePrerelease), nil
}
