package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
	return nil, nil
}
