package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/internal"
)

var _ = Describe("Firmware", func() {
	Describe("IsOlderThan", func() {
		When("the versions are the same", func() {
			It("returns false", func() {
				firmwareA := &Firmware{
					Type:    "bootstrap",
					Version: "1.0.0",
				}
				firmwareB := &Firmware{
					Type:    "bootstrap",
					Version: "1.0.0",
				}
				Expect(firmwareA.IsOlderThan(firmwareB)).To(BeFalse())
			})
		})

		When("the first version is older than the second", func() {
			It("returns true", func() {
				firmwareA := &Firmware{
					Type:    "bootstrap",
					Version: "1.0.0",
				}
				firmwareB := &Firmware{
					Type:    "bootstrap",
					Version: "2.0.0",
				}
				Expect(firmwareA.IsOlderThan(firmwareB)).To(BeTrue())
			})
		})

		When("the first version is newer than the second", func() {
			It("returns false", func() {
				firmwareA := &Firmware{
					Type:    "bootstrap",
					Version: "2.0.0",
				}
				firmwareB := &Firmware{
					Type:    "bootstrap",
					Version: "1.0.0",
				}
				Expect(firmwareA.IsOlderThan(firmwareB)).To(BeFalse())
			})
		})
	})
})

var _ = Describe("FirmwareList", func() {
	Describe("GetLatest", func() {
		var list FirmwareList

		BeforeEach(func() {
			list = FirmwareList{
				&Firmware{Type: "a", Version: "3.0.0-rc.1"},
				&Firmware{Type: "a", Version: "1.0.0"},
				&Firmware{Type: "a", Version: "1.0.0-rc.1"},
				&Firmware{Type: "a", Version: "2.0.0"},
			}
		})

		It("returns the latest version", func() {
			firmware := list.GetLatest(false)
			Expect(firmware).ToNot(BeNil())
			Expect(firmware.Version).To(Equal("2.0.0"))
		})

		When("accepting prerelease versions", func() {
			It("returns the latest version, including prerelease", func() {
				firmware := list.GetLatest(true)
				Expect(firmware).ToNot(BeNil())
				Expect(firmware.Version).To(Equal("3.0.0-rc.1"))
			})
		})

		When("the list is empty", func() {
			BeforeEach(func() {
				list = FirmwareList{}
			})
			It("returns nil", func() {
				firmware := list.GetLatest(false)
				Expect(firmware).To(BeNil())
			})
		})

		When("there are no released versions", func() {
			BeforeEach(func() {
				list = FirmwareList{
					&Firmware{Type: "a", Version: "3.0.0-rc.1"},
					&Firmware{Type: "a", Version: "1.0.0-rc.1"},
				}
			})
			It("returns nil", func() {
				firmware := list.GetLatest(false)
				Expect(firmware).To(BeNil())
			})
		})
	})

	Describe("Sort", func() {
		It("sorts the list", func() {
			firmwareList := FirmwareList{
				&Firmware{Type: "a", Version: "1.0.0"},
				&Firmware{Type: "a", Version: "3.0.0"},
				&Firmware{Type: "a", Version: "1.0.0-rc.1"},
				&Firmware{Type: "a", Version: "2.0.0"},
			}
			firmwareList.Sort()
			Expect(firmwareList[0].Version).To(Equal("1.0.0-rc.1"))
			Expect(firmwareList[1].Version).To(Equal("1.0.0"))
			Expect(firmwareList[2].Version).To(Equal("2.0.0"))
			Expect(firmwareList[3].Version).To(Equal("3.0.0"))
		})
	})

	Describe("FilterOutPrerelease", func() {
		It("filters out prerelease versions", func() {
			firmwareList := FirmwareList{
				&Firmware{Type: "a", Version: "1.0.0"},
				&Firmware{Type: "a", Version: "3.0.0"},
				&Firmware{Type: "a", Version: "1.0.0-rc.1"},
				&Firmware{Type: "a", Version: "2.0.0"},
			}
			filtered := firmwareList.FilterOutPrerelease()
			Expect(filtered[0].Version).To(Equal("1.0.0"))
			Expect(filtered[1].Version).To(Equal("3.0.0"))
			Expect(filtered[2].Version).To(Equal("2.0.0"))
		})

		Context("empty list", func() {
			It("returns an empty list", func() {
				firmwareList := FirmwareList{}
				filtered := firmwareList.FilterOutPrerelease()
				Expect(filtered).To(BeEmpty())
			})
		})
	})
})
