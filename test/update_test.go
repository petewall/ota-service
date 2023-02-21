package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	d "github.com/petewall/device-service/lib"
	f "github.com/petewall/firmware-service/lib"
	"net/http"
)

var _ = Describe("Update", func() {
	When("a firmware update is available for this device", func() {
		BeforeEach(func() {
			deviceService.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/aa:bb:cc:dd:ee:ff"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, d.Device{
						MAC:              "aa:bb:cc:dd:ee:ff",
						Firmware:         "bootstrap",
						Version:          "1.0.0",
						AssignedFirmware: "lightswitch",
						AssignedVersion:  "2.0.0",
					}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/aa:bb:cc:dd:ee:ff"),
					ghttp.VerifyJSONRepresenting(&d.UpdateDevicePayload{
						Firmware: "bootstrap",
						Version:  "1.2.3",
					}),
					ghttp.RespondWith(http.StatusOK, nil),
				),
			)

			firmwareData := []byte("this is the firmware data")
			firmwareService.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/lightswitch/2.0.0"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, f.Firmware{
						Type:    "lightswitch",
						Version: "2.0.0",
						Size:    int64(len(firmwareData)),
					}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/lightswitch/2.0.0/data"),
					ghttp.RespondWith(http.StatusOK, firmwareData),
				),
			)
		})

		It("returns the updated firmware", func() {
			firmwareData, err := client.Update("aa:bb:cc:dd:ee:ff", "bootstrap", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareData).To(Equal([]byte("this is the firmware data")))
		})
	})

	Context("there is no firmware assignment", func() {
		BeforeEach(func() {
			deviceService.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/aa:bb:cc:dd:ee:ff"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, d.Device{
						MAC:      "aa:bb:cc:dd:ee:ff",
						Firmware: "bootstrap",
						Version:  "1.0.0",
					}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/aa:bb:cc:dd:ee:ff"),
					ghttp.VerifyJSONRepresenting(&d.UpdateDevicePayload{
						Firmware: "bootstrap",
						Version:  "1.2.3",
					}),
					ghttp.RespondWith(http.StatusOK, nil),
				),
			)
		})

		It("returns no updates", func() {
			firmwareData, err := client.Update("aa:bb:cc:dd:ee:ff", "bootstrap", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareData).To(BeNil())
		})
	})

	Context("there is no firmware version assigned", func() {
		BeforeEach(func() {
			deviceService.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/aa:bb:cc:dd:ee:ff"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, d.Device{
						MAC:              "aa:bb:cc:dd:ee:ff",
						Firmware:         "bootstrap",
						Version:          "1.0.0",
						AssignedFirmware: "lightswitch",
					}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/aa:bb:cc:dd:ee:ff"),
					ghttp.VerifyJSONRepresenting(&d.UpdateDevicePayload{
						Firmware: "bootstrap",
						Version:  "1.2.3",
					}),
					ghttp.RespondWith(http.StatusOK, nil),
				),
			)
			firmwareData := []byte("this is the firmware data")
			firmwareService.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/lightswitch"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, f.FirmwareList{
						&f.Firmware{
							Type:    "lightswitch",
							Version: "2.0.0",
							Size:    int64(len(firmwareData)),
						},
						&f.Firmware{
							Type:    "lightswitch",
							Version: "1.0.0",
							Size:    int64(len(firmwareData)),
						},
					}),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/lightswitch/2.0.0/data"),
					ghttp.RespondWith(http.StatusOK, firmwareData),
				),
			)
		})

		It("returns the latest firmware", func() {
			firmwareData, err := client.Update("aa:bb:cc:dd:ee:ff", "bootstrap", "1.2.3")
			Expect(err).ToNot(HaveOccurred())
			Expect(firmwareData).To(Equal([]byte("this is the firmware data")))
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
		It("returns a 400 error", func() {
			_, err := client.Update("", "", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("update request failed: 400 Bad Request: X-ESP8266-STA-MAC is not set"))
		})
	})

	When("the current firmware is not set", func() {
		It("returns a 400 error", func() {
			_, err := client.Update("aa:bb:cc:dd:ee:ff", "", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("update request failed: 400 Bad Request: current firmware type was not sent"))
		})
	})

	When("the device service request fails", func() {
		BeforeEach(func() {
			deviceService.AppendHandlers(
				ghttp.VerifyRequest(http.MethodGet, "/aa:bb:cc:dd:ee:ff"),
				ghttp.RespondWith(http.StatusInternalServerError, nil),
			)
		})
		It("returns a 500 error", func() {
			_, err := client.Update("aa:bb:cc:dd:ee:ff", "bootstrap", "1.2.3")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("update request failed: 500 Internal Server"))
		})
	})
})
