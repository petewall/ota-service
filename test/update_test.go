package test_test

import . "github.com/onsi/ginkgo/v2"

var _ = Describe("Update", func() {
	XIt("returns the updated firmware", func() {
		By("getting the device information from the firmware service", func() {})
		By("getting the firmware from the firmware service", func() {})
		By("returning the firmware data")
	})

	Context("the device is unknown to the device service", func() {
		XIt("updates the device service", func() {

		})

		When("the device service update request fails", func() {
			XIt("returns 500 error")
		})
	})

	Context("there is no firmware assignment", func() {
		XIt("returns no updates", func() {

		})
	})

	Context("there is no firmware version assigned", func() {
		XIt("returns the latest firmware", func() {

		})

		When("the latest firmware is not newer than what is on the device", func() {
			XIt("returns no updates", func() {

			})
		})

		When("getting the latest firmware from the firmware service fails", func() {
			XIt("returns a 500 error", func() {

			})
		})
	})

	When("the firmware is the same as what is on the device", func() {
		XIt("returns no updates", func() {

		})
	})

	When("getting the firmware from the firmware service fails", func() {
		XIt("returns a 500 error", func() {

		})
	})

	When("the MAC address is not set", func() {
		XIt("returns a 400 error", func() {

		})
	})

	When("the current firmware is not set", func() {
		XIt("returns a 400 error", func() {

		})
	})

	When("the device service request fails", func() {
		XIt("returns a 500 error", func() {

		})
	})
})
