package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/internal"
)

var _ = Describe("Device", func() {
	Describe("IsDifferent", func() {
		var device *Device

		BeforeEach(func() {
			device = &Device{
				CurrentFirmware: "bootstrap",
				CurrentVersion:  "1.2.3",
			}
		})

		When("the values are the same", func() {
			It("returns false", func() {
				firmware := &Firmware{
					Type:    "bootstrap",
					Version: "1.2.3",
				}
				Expect(device.IsDifferent(firmware)).To(BeFalse())
			})
		})

		When("the type is different", func() {
			It("returns true", func() {
				firmware := &Firmware{
					Type:    "lightswitch",
					Version: "1.2.3",
				}
				Expect(device.IsDifferent(firmware)).To(BeTrue())
			})
		})

		When("the version is different", func() {
			It("returns true", func() {
				firmware := &Firmware{
					Type:    "bootstrap",
					Version: "2.3.4",
				}
				Expect(device.IsDifferent(firmware)).To(BeTrue())
			})
		})
	})
})
