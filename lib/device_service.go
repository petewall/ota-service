package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//go:generate counterfeiter -generate

//counterfeiter:generate . DeviceService
type DeviceService interface {
	GetDevice(mac string) (*Device, error)
	UpdateDevice(device *Device) error
}

type DeviceServiceImpl struct {
	Host       string
	Port       int
	HTTPClient HTTPClient
}

func (d *DeviceServiceImpl) GetDevice(mac string) (*Device, error) {
	resp, err := d.HTTPClient.Get(fmt.Sprintf("http://%s:%d/%s", d.Host, d.Port, mac))
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read device response: %w", err)
	}

	var device *Device
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, fmt.Errorf("failed to parse device response: %w", err)
	}

	return device, nil
}

func (d *DeviceServiceImpl) UpdateDevice(device *Device) error {
	encoded, err := json.Marshal(device)
	if err != nil {
		return fmt.Errorf("failed to prepare device update request body: %w", err)
	}

	url := fmt.Sprintf("http://%s:%d/%s", d.Host, d.Port, device.MAC)
	resp, err := d.HTTPClient.Post(url, "application/json", bytes.NewReader(encoded))
	if err != nil {
		return fmt.Errorf("failed to send device update request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("device update request failed (%d), and failed to get response body: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("device update request failed (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}
