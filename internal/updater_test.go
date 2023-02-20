package internal_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/internal"
	. "github.com/petewall/ota-service/v2/internal/internalfakes"
)

var _ = Describe("Updater", func() {
	var (
		deviceService   *FakeDeviceService
		firmwareService *FakeFirmwareService
		updater         Updater
	)
	BeforeEach(func() {
		deviceService = &FakeDeviceService{}
		firmwareService = &FakeFirmwareService{}
		updater = &UpdaterImpl{
			DeviceService:   deviceService,
			FirmwareService: firmwareService,
		}
	})

	Describe("Update", func() {
		When("mac is empty", func() {
			It("returns an error", func() {
				_, err := updater.Update("", nil)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("mac not set"))

				Expect(deviceService.GetDeviceCallCount()).To(Equal(0))
				Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
				Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
				Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
			})
		})

		When("firmware is nil", func() {
			It("returns an error", func() {
				_, err := updater.Update("aa:bb:cc:dd:ee:ff", nil)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware not set"))

				Expect(deviceService.GetDeviceCallCount()).To(Equal(0))
				Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
				Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
				Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
			})
		})

		When("device service fails to get device details", func() {
			BeforeEach(func() {
				deviceService.GetDeviceReturns(nil, errors.New("get device failed"))
			})
			It("returns an error", func() {
				_, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
					Type:    "bootstrap",
					Version: "1.0.0",
				})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to get device: get device failed"))

				Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
				Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
				Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
				Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
			})
		})

		Context("device service has not seen the device before", func() {
			BeforeEach(func() {
				deviceService.GetDeviceReturns(nil, nil)
			})

			When("an update request comes in", func() {
				It("returns no firmware", func() {
					By("updating the device service", func() {
						firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
							Type:    "bootstrap",
							Version: "1.0.0",
						})
						Expect(err).ToNot(HaveOccurred())
						Expect(firmware).To(BeNil())

						Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
						Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
						Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1))
						device := deviceService.UpdateDeviceArgsForCall(0)
						Expect(device.CurrentFirmware).To(Equal("bootstrap"))
						Expect(device.CurrentVersion).To(Equal("1.0.0"))
						Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
						Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
					})
				})
			})

			When("device service fails to update", func() {
				BeforeEach(func() {
					deviceService.UpdateDeviceReturns(errors.New("update device failed"))
				})

				It("returns an error", func() {
					_, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("unable to update device: update device failed"))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1))
					device := deviceService.UpdateDeviceArgsForCall(0)
					Expect(device.CurrentFirmware).To(Equal("bootstrap"))
					Expect(device.CurrentVersion).To(Equal("1.0.0"))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})
		})

		Context("existing device has no assigned firmware", func() {
			BeforeEach(func() {
				deviceService.GetDeviceReturns(&Device{
					MAC:             "aa:bb:cc:dd:ee:ff",
					CurrentFirmware: "bootstrap",
					CurrentVersion:  "1.0.0",
				}, nil)
			})
			When("an update request comes in", func() {
				It("returns no firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware).To(BeNil())

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})

			Context("the existing device has a different firmware", func() {
				When("an update request comes in", func() {
					It("updates the device", func() {
						firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
							Type:    "bootstrap",
							Version: "2.0.0",
						})
						Expect(err).ToNot(HaveOccurred())
						Expect(firmware).To(BeNil())

						Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
						Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
						Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1))
						device := deviceService.UpdateDeviceArgsForCall(0)
						Expect(device.CurrentFirmware).To(Equal("bootstrap"))
						Expect(device.CurrentVersion).To(Equal("2.0.0"))
						Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
						Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
					})
				})
			})
		})

		Context("existing device has no assigned firmware version", func() {
			BeforeEach(func() {
				deviceService.GetDeviceReturns(&Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					CurrentFirmware:  "bootstrap",
					CurrentVersion:   "1.0.0",
					AssignedFirmware: "temp-sensor",
				}, nil)
				firmwareService.GetLatestFirmwareReturns(&Firmware{
					Type:    "temp-sensor",
					Version: "1.2.3",
					Size:    len("temp-sensor data"),
					Data:    []byte("temp-sensor data"),
				}, nil)
			})

			When("an update request comes in from a device with different firmware", func() {
				It("returns the latest assigned firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware.Type).To(Equal("temp-sensor"))
					Expect(firmware.Version).To(Equal("1.2.3"))
					Expect(firmware.Size).To(Equal(16))
					Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
					firmwareType, acceptsPrerelease := firmwareService.GetLatestFirmwareArgsForCall(0)
					Expect(firmwareType).To(Equal("temp-sensor"))
					Expect(acceptsPrerelease).To(BeFalse())
				})

				When("the device accepts prerelease firmware", func() {
					BeforeEach(func() {
						deviceService.GetDeviceReturns(&Device{
							MAC:               "aa:bb:cc:dd:ee:ff",
							CurrentFirmware:   "bootstrap",
							CurrentVersion:    "1.0.0",
							AssignedFirmware:  "temp-sensor",
							AcceptsPrerelease: true,
						}, nil)
					})
					It("passes that along to the firmware service request", func() {
						firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
							Type:    "bootstrap",
							Version: "1.0.0",
						})
						Expect(err).ToNot(HaveOccurred())
						Expect(firmware.Type).To(Equal("temp-sensor"))
						Expect(firmware.Version).To(Equal("1.2.3"))
						Expect(firmware.Size).To(Equal(16))
						Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

						Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
						Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
						Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
						Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
						Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
						firmwareType, acceptsPrerelease := firmwareService.GetLatestFirmwareArgsForCall(0)
						Expect(firmwareType).To(Equal("temp-sensor"))
						Expect(acceptsPrerelease).To(BeTrue())
					})
				})
			})

			When("an update request comes in from a device with older firmware", func() {
				It("returns the latest firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "1.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware.Type).To(Equal("temp-sensor"))
					Expect(firmware.Version).To(Equal("1.2.3"))
					Expect(firmware.Size).To(Equal(16))
					Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
					Expect(firmwareService.GetLatestFirmwareArgsForCall(0)).To(Equal("temp-sensor"))
				})
			})

			When("an update request comes in from a device with latest firmware", func() {
				It("returns no firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "1.2.3",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware).To(BeNil())

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
					Expect(firmwareService.GetLatestFirmwareArgsForCall(0)).To(Equal("temp-sensor"))
				})
			})
			When("an update request comes in from a device with newer firmware", func() {
				It("returns no firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "2.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware).To(BeNil())

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
					Expect(firmwareService.GetLatestFirmwareArgsForCall(0)).To(Equal("temp-sensor"))
				})
			})

			When("the firmware service fails to return the latest assigned firmware", func() {
				BeforeEach(func() {
					firmwareService.GetLatestFirmwareReturns(nil, errors.New("get latest firmware failed"))
				})
				It("returns an error", func() {
					_, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("unable to get latest firmware: get latest firmware failed"))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(1))
					Expect(firmwareService.GetLatestFirmwareArgsForCall(0)).To(Equal("temp-sensor"))
				})
			})
		})

		Context("existing device has a pinned firmware version", func() {
			BeforeEach(func() {
				deviceService.GetDeviceReturns(&Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					CurrentFirmware:  "bootstrap",
					CurrentVersion:   "1.0.0",
					AssignedFirmware: "temp-sensor",
					AssignedVersion:  "1.2.3",
				}, nil)
				firmwareService.GetFirmwareReturns(&Firmware{
					Type:    "temp-sensor",
					Version: "1.2.3",
					Size:    len("temp-sensor data"),
					Data:    []byte("temp-sensor data"),
				}, nil)
			})

			When("an update request comes in from a device with different firmware", func() {
				It("returns the assigned firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware.Type).To(Equal("temp-sensor"))
					Expect(firmware.Version).To(Equal("1.2.3"))
					Expect(firmware.Size).To(Equal(16))
					Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(1))
					firmwareType, firmwareVersion := firmwareService.GetFirmwareArgsForCall(0)
					Expect(firmwareType).To(Equal("temp-sensor"))
					Expect(firmwareVersion).To(Equal("1.2.3"))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})

			When("an update request comes in from a device with older firmware", func() {
				It("returns the assigned firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "1.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware.Type).To(Equal("temp-sensor"))
					Expect(firmware.Version).To(Equal("1.2.3"))
					Expect(firmware.Size).To(Equal(16))
					Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(1))
					firmwareType, firmwareVersion := firmwareService.GetFirmwareArgsForCall(0)
					Expect(firmwareType).To(Equal("temp-sensor"))
					Expect(firmwareVersion).To(Equal("1.2.3"))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})
			When("an update request comes in from a device with assigned firmware", func() {
				It("returns no firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "1.2.3",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware).To(BeNil())

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(0))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})

			When("an update request comes in from a device with newer firmware", func() {
				It("returns the assigned firmware", func() {
					firmware, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "temp-sensor",
						Version: "2.0.0",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(firmware.Type).To(Equal("temp-sensor"))
					Expect(firmware.Version).To(Equal("1.2.3"))
					Expect(firmware.Size).To(Equal(16))
					Expect(firmware.Data).To(Equal([]byte("temp-sensor data")))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(1)) // Not checking the args, because we tested that above
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(1))
					firmwareType, firmwareVersion := firmwareService.GetFirmwareArgsForCall(0)
					Expect(firmwareType).To(Equal("temp-sensor"))
					Expect(firmwareVersion).To(Equal("1.2.3"))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})

			When("the firmware service fails to return the assigned firmware", func() {
				BeforeEach(func() {
					firmwareService.GetFirmwareReturns(nil, errors.New("get firmware failed"))
				})
				It("returns an error", func() {
					_, err := updater.Update("aa:bb:cc:dd:ee:ff", &Firmware{
						Type:    "bootstrap",
						Version: "1.0.0",
					})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("unable to get firmware: get firmware failed"))

					Expect(deviceService.GetDeviceCallCount()).To(Equal(1))
					Expect(deviceService.GetDeviceArgsForCall(0)).To(Equal("aa:bb:cc:dd:ee:ff"))
					Expect(deviceService.UpdateDeviceCallCount()).To(Equal(0))
					Expect(firmwareService.GetFirmwareCallCount()).To(Equal(1))
					firmwareType, firmwareVersion := firmwareService.GetFirmwareArgsForCall(0)
					Expect(firmwareType).To(Equal("temp-sensor"))
					Expect(firmwareVersion).To(Equal("1.2.3"))
					Expect(firmwareService.GetLatestFirmwareCallCount()).To(Equal(0))
				})
			})
		})
	})
})
