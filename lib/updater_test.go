package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/lib"
	. "github.com/petewall/ota-service/v2/lib/libfakes"
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
			})
		})

		When("firmware is nil", func() {
			It("returns an error", func() {
				_, err := updater.Update("aa:bb:cc:dd:ee:ff", nil)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("firmware not set"))
			})
		})

		When("device service fails to get device details", func() {
			XIt("returns an error", func() {})
		})

		Context("device service has not seen the device before", func() {
			When("an update request comes it", func() {
				XIt("returns no firmware", func() {
					By("updating the device service", func() {})
				})
			})
		})

		Context("existing device has no assigned firmware", func() {
			When("an update request comes it", func() {
				XIt("returns no firmware", func() {})
			})
		})
		Context("existing device has no assigned firmware version", func() {
			When("an update request comes in from a device with different firmware", func() {
				XIt("returns the latest assigned firmware", func() {})
			})
			When("an update request comes in from a device with older firmware", func() {
				XIt("returns the latest firmware", func() {})
			})
			When("an update request comes in from a device with latest firmware", func() {
				XIt("returns no firmware", func() {})
			})
			When("an update request comes in from a device with newer firmware", func() {
				XIt("returns no firmware", func() {})
			})
		})
		Context("existing device has a pinned firmware version", func() {
			When("an update request comes in from a device with different firmware", func() {
				XIt("returns the assigned firmware", func() {})
			})
			When("an update request comes in from a device with older firmware", func() {
				XIt("returns the assigned firmware", func() {})
			})
			When("an update request comes in from a device with assigned firmware", func() {
				XIt("returns no firmware", func() {})
			})
			When("an update request comes in from a device with newer firmware", func() {
				XIt("returns the assigned firmware", func() {})
			})
		})
	})
})
