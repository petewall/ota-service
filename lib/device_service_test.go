package lib_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/lib"
	. "github.com/petewall/ota-service/v2/lib/libfakes"
)

var _ = Describe("Device Service", func() {
	var (
		httpClient    *FakeHTTPClient
		deviceService DeviceService
	)
	BeforeEach(func() {
		httpClient = &FakeHTTPClient{}
		deviceService = &DeviceServiceImpl{
			Host:       "example.petewall.net",
			Port:       9876,
			HTTPClient: httpClient,
		}
	})

	Describe("GetDevice", func() {
		BeforeEach(func() {
			device := &Device{
				MAC:              "aa:bb:cc:dd:ee:ff",
				CurrentFirmware:  "bootstrap",
				CurrentVersion:   "1.2.3",
				AssignedFirmware: "switch",
				AssignedVersion:  "2.0.0",
			}

			encoded, err := json.Marshal(device)
			Expect(err).ToNot(HaveOccurred())

			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(encoded)),
			}
			httpClient.GetReturns(response, nil)
		})

		It("returns the device", func() {
			device, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(device.CurrentFirmware).To(Equal("bootstrap"))
			Expect(device.CurrentVersion).To(Equal("1.2.3"))
			Expect(device.AssignedFirmware).To(Equal("switch"))
			Expect(device.AssignedVersion).To(Equal("2.0.0"))
		})

		When("getting the device fails", func() {
			BeforeEach(func() {
				httpClient.GetReturns(nil, errors.New("get device failed"))
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to get device aa:bb:cc:dd:ee:ff: get device failed"))
			})
		})

		When("reading the response body fails", func() {
			BeforeEach(func() {
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(&FailingReader{Message: "read failed"}),
				}
				httpClient.GetReturns(response, nil)
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to read device aa:bb:cc:dd:ee:ff response: read failed"))
			})
		})

		When("parsing] the response body fails", func() {
			BeforeEach(func() {
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader("this is --- not json")),
				}
				httpClient.GetReturns(response, nil)
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to parse device aa:bb:cc:dd:ee:ff response: invalid character 'h' in literal true (expecting 'r')"))
			})
		})
	})

	Describe("UpdateDevice", func() {
		BeforeEach(func() {
			httpClient.PostReturns(&http.Response{
				StatusCode: http.StatusOK,
			}, nil)
		})

		It("posts the device", func() {
			err := deviceService.UpdateDevice(&Device{
				MAC:              "aa:bb:cc:dd:ee:ff",
				CurrentFirmware:  "bootstrap",
				CurrentVersion:   "1.2.3",
				AssignedFirmware: "switch",
				AssignedVersion:  "2.0.0",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(httpClient.PostCallCount()).To(Equal(1))
			url, contentType, body := httpClient.PostArgsForCall(0)
			Expect(url).To(Equal("http://example.petewall.net:9876/aa:bb:cc:dd:ee:ff"))
			Expect(contentType).To(Equal("application/json"))

			deviceContent, err := ioutil.ReadAll(body)
			Expect(err).ToNot(HaveOccurred())

			var device *Device
			err = json.Unmarshal(deviceContent, &device)
			Expect(err).ToNot(HaveOccurred())
			Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(device.CurrentFirmware).To(Equal("bootstrap"))
			Expect(device.CurrentVersion).To(Equal("1.2.3"))
			Expect(device.AssignedFirmware).To(Equal("switch"))
			Expect(device.AssignedVersion).To(Equal("2.0.0"))
		})

		When("the request fails", func() {
			BeforeEach(func() {
				httpClient.PostReturns(nil, errors.New("update device aa:bb:cc:dd:ee:ff failed"))
			})

			It("returns an error", func() {
				err := deviceService.UpdateDevice(&Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					CurrentFirmware:  "bootstrap",
					CurrentVersion:   "1.2.3",
					AssignedFirmware: "switch",
					AssignedVersion:  "2.0.0",
				})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to send device aa:bb:cc:dd:ee:ff update request: update device aa:bb:cc:dd:ee:ff failed"))
			})
		})

		When("the response status is not OK", func() {
			BeforeEach(func() {
				httpClient.PostReturns(&http.Response{
					StatusCode: http.StatusTeapot,
					Body:       io.NopCloser(strings.NewReader("i'm a teapot")),
				}, nil)
			})

			It("returns an error with the response body", func() {
				err := deviceService.UpdateDevice(&Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					CurrentFirmware:  "bootstrap",
					CurrentVersion:   "1.2.3",
					AssignedFirmware: "switch",
					AssignedVersion:  "2.0.0",
				})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("device aa:bb:cc:dd:ee:ff update request failed (418): i'm a teapot"))
			})

			When("you cannot read the response body", func() {
				BeforeEach(func() {
					httpClient.PostReturns(&http.Response{
						StatusCode: http.StatusTeapot,
						Body:       io.NopCloser(&FailingReader{Message: "oops, all errors"}),
					}, nil)
				})
				It("returns an error without the response body", func() {
					err := deviceService.UpdateDevice(&Device{
						MAC:              "aa:bb:cc:dd:ee:ff",
						CurrentFirmware:  "bootstrap",
						CurrentVersion:   "1.2.3",
						AssignedFirmware: "switch",
						AssignedVersion:  "2.0.0",
					})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("device aa:bb:cc:dd:ee:ff update request failed (418), and failed to get response body: oops, all errors"))
				})
			})
		})
	})
})
